package libs

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"regexp"
	"strings"
)

func HTMLContent(styleByte []byte, title, body string) []byte {
	style := strings.Replace(string(styleByte), "html,body,#app{background-color:#39394d}", "", 1)
	re := regexp.MustCompile(`@font-face{.*?}::`)
	matches := re.FindAllString(style, -1)

	// 打印匹配结果
	for _, match := range matches {
		style = strings.Replace(style, match, "::", 1)
	}

	htmlContent := `<!DOCTYPE html>
<html lang="zh">
<head>
    <meta charset="UTF-8">
    <title>` + title + `的简历</title>
    <style>` + style + `</style>
</head>
<body>` + body + `</body>
</html>`
	//os.WriteFile("index.html", []byte(htmlContent), 0644)
	return []byte(htmlContent)
}

// ToPDF 将指定 HTML 转换为 PDF
func ToPDF(res *[]byte, style []byte, title, body string) chromedp.Tasks {
	dataURL := fmt.Sprintf("data:text/html;base64,%s", base64.StdEncoding.EncodeToString(HTMLContent(style, title, body)))

	return chromedp.Tasks{
		chromedp.Navigate(dataURL),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// 设置打印参数
			buf, _, err := page.PrintToPDF().
				WithLandscape(false).
				WithPrintBackground(true).
				WithScale(1.0).
				WithPaperWidth(210.0 / 25.4).  // 将毫米转换为英寸
				WithPaperHeight(297.0 / 25.4). // 将毫米转换为英寸
				WithMarginTop(0.2).
				WithMarginBottom(0.2).
				WithMarginLeft(0.0).
				WithMarginRight(0.0).
				Do(ctx) // 生成PDF

			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}
