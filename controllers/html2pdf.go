package controllers

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"resume/libs"
	"resume/types/enums"
	"strings"
)

var request struct {
	HTMLContent string `json:"htmlContent" binding:"required"`
}

func Html2PDF(c *gin.Context) {
	s := libs.Context(c)

	bt := s.GetHeader("Browser-Type")

	browserExePath := enums.BrowserExePath.Get(bt)

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Html2PDF 参数序列化失败：", err.Error())
		s.Msg(http.StatusInternalServerError, "参数错误，请重试！")
		return
	}

	readDir, err := s.Embed.ReadDir("public/assets")
	if err != nil {
		log.Println("Html2PDF 读取 Embed 文件夹失败：", err.Error())
		s.Msg(http.StatusInternalServerError, "系统错误，请重试！")
		return
	}

	var indexStyles []string
	for _, file := range readDir {
		if (strings.HasPrefix(file.Name(), "index-") || strings.HasPrefix(file.Name(), "vendor-")) && strings.HasSuffix(file.Name(), ".css") {
			indexStyle, _ := s.Embed.ReadFile(fmt.Sprintf("public/assets/%s", file.Name()))
			indexStyles = append(indexStyles, string(indexStyle))
		}
	}

	// 配置浏览器启动选项
	opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.ExecPath(browserExePath))

	// 启动浏览器分配器
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// 创建浏览器上下文
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// 捕获 PDF
	var buf []byte
	if err := chromedp.Run(ctx, libs.ToPDF(&buf, []byte(strings.Join(indexStyles, "")), s.Resume.Name, request.HTMLContent)); err != nil {
		ctx.Done()
		allocCtx.Done()
		log.Println("Html2PDF 生成 PDF 文档失败：", err.Error())
		log.Println("Html2PDF 生成 PDF 文档失败，可能缺少Chrome/Edge浏览器，请检查配置文件浏览器配置路径。")
		s.Msg(http.StatusInternalServerError, "PDF 文档生成失败，请重试！")
		return
	}

	_filePath := filepath.Join(enums.TempPath, s.User.ID)

	err = os.WriteFile(_filePath, buf, 0o644)
	if err != nil {
		log.Println("Html2PDF PDF 报存失败：", err.Error())
		s.Msg(http.StatusInternalServerError, "PDF 文档生成失败，请重试！")
		return
	}

	// 返回生成的 PDF
	//s.Json(http.StatusOK, s.Resume.Name+"的简历.pdf", strings.Replace(_filePath, enums.DataPath, "", 1))
	s.Json(http.StatusOK, s.Resume.Name+"的简历.pdf", fmt.Sprintf("/temp/%s", s.User.ID))
}

func DeletePDF(c *gin.Context) {
	s := libs.Context(c)

	_filePath := filepath.Join(enums.TempPath, s.User.ID)
	if err := os.Remove(_filePath); err != nil {
		log.Println(fmt.Sprintf("%s 删除失败！", _filePath), err.Error())
		s.Msg(http.StatusInternalServerError, "PDF 文档生成失败，请重试！")
		return
	}
	s.Msg(http.StatusOK, "PDF 文档生成成功！")
}
