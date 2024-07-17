package controllers

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"resume/libs"
	"strings"
)

// 将指定 HTML 转换为 PDF
func printToPDF(dataURL string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(dataURL),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// 设置打印参数
			buf, _, err := page.PrintToPDF().
				WithLandscape(false).
				WithPrintBackground(true).
				WithScale(1.0).
				WithPaperWidth(8.5).
				WithPaperHeight(11.0).
				WithMarginTop(0.0).
				WithMarginBottom(0.0).
				WithMarginLeft(0.0).
				WithMarginRight(0.0).Do(ctx) // 生成PDF

			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}

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

	// 创建 chromedp 上下文
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	readDir, err := s.Embed.ReadDir("public/assets")
	if err != nil {
		log.Println("Html2PDF 读取 Embed 文件夹失败：", err.Error())
		s.Msg(http.StatusInternalServerError, "系统错误，请重试！")
		return
	}

	var fileIndex string

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

	// 使用 data: URL scheme 将 HTML 字符串内容传递给 chromedp
	dataURL := "data:text/html;base64," + base64.StdEncoding.EncodeToString([]byte(`<!doctype html>
<html lang="zh">
  <head>
    <meta charset="UTF-8" />
    <link rel="icon" type="image/svg+xml" href="/assets/vite.svg" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
	<style type="text/css">`+string(indexStyle)+`</style>
    <title>`+s.Resume.Name+`的简历</title>
  </head>
  <body style="background-color: #FFF">`+request.HTMLContent+`</body>
</html>`))

	// 捕获 PDF
	var buf []byte
	err = chromedp.Run(ctx, printToPDF(dataURL, &buf))

	if err != nil {
		log.Println("Html2PDF 生成 PDF 文档失败：", err.Error())
		s.Msg(http.StatusInternalServerError, "PDF 文档生成失败，请重试！")
		return
	}

	// 返回生成的 PDF
	c.Data(http.StatusOK, "application/pdf", buf)
}
