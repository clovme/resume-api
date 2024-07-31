package controllers

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"resume/libs"
	"strings"
)

func Html2PDF(c *gin.Context) {
	s := libs.Context(c)
	var request struct {
		HTMLContent string `json:"htmlContent" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Html2PDF 参数序列化失败：", err.Error())
		s.Msg(http.StatusInternalServerError, "参数错误，请重试！")
		return
	}

	var fileIndex string
	readDir, err := s.Embed.ReadDir("public/assets")
	if err != nil {
		log.Println("Html2PDF 读取 Embed 文件夹失败：", err.Error())
		s.Msg(http.StatusInternalServerError, "系统错误，请重试！")
		return
	}

	for _, file := range readDir {
		if strings.HasPrefix(file.Name(), "index-") && strings.HasSuffix(file.Name(), ".css") {
			fileIndex = file.Name()
			break
		}
	}

	indexStyle, err := s.Embed.ReadFile(fmt.Sprintf("public/assets/%s", fileIndex))
	if err != nil {
		log.Println("Html2PDF 打开样式文件失败：", err.Error())
		s.Msg(http.StatusInternalServerError, "系统错误，请重试！")
		return
	}

	// 创建 chromedp 上下文
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// 捕获 PDF
	var buf []byte
	if err := chromedp.Run(ctx, libs.Base64ToPDF(&buf, indexStyle, s.Resume.Name, request.HTMLContent)); err != nil {
		log.Println("Html2PDF 生成 PDF 文档失败，可能缺少Google浏览器，请下载Google浏览器，错误信息：", err.Error())
		s.Msg(http.StatusInternalServerError, "Html2PDF 生成 PDF 文档失败，可能缺少Google浏览器，请下载Google浏览器，请重试！")
		return
	}

	filepath := fmt.Sprintf("./data/temp/%s", s.User.ID)

	err = os.WriteFile(filepath, buf, 0o644)
	if err != nil {
		log.Println("Html2PDF PDF 报存失败：", err.Error())
		s.Msg(http.StatusInternalServerError, "PDF 文档生成失败，请重试！")
		return
	}

	// 返回生成的 PDF
	s.Json(http.StatusOK, s.Resume.Name+"的简历.pdf", strings.Replace(filepath, "./data", "", 1))
}

func DeletePDF(c *gin.Context) {
	s := libs.Context(c)

	filepath := fmt.Sprintf("./data/temp/%s", s.User.ID)
	if err := os.Remove(filepath); err != nil {
		log.Println(fmt.Sprintf("%s 删除失败！", filepath), err.Error())
		s.Msg(http.StatusInternalServerError, "PDF 文档生成失败，请重试！")
		return
	}
	s.Msg(http.StatusOK, "PDF 文档生成成功！")
}
