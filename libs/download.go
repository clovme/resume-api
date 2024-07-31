package libs

import (
	"archive/zip"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/chromedp/chromedp"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type ProgressWriter struct {
	Total      int64
	Downloaded int64
}

func (pw *ProgressWriter) Write(p []byte) (int, error) {
	n := len(p)
	pw.Downloaded += int64(n)
	downloadedMB := float64(pw.Downloaded) / (1024 * 1024)
	totalMB := float64(pw.Total) / (1024 * 1024)
	fmt.Printf("\r下载进度: %.2f/%.2f MB (%.2f%%)", downloadedMB, totalMB, downloadedMB*100/totalMB)
	return n, nil
}

func fileExists(filepath string) (bool, error) {
	info, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false, nil
	}
	return !info.IsDir(), err
}

func folderExists(folderpath string) bool {
	info, err := os.Stat(folderpath)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func getFileSize(filepath string) (int64, error) {
	info, err := os.Stat(filepath)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

func saveFile(url string, dest string) (int64, error) {
	// 创建HTTP请求
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// 检查HTTP响应状态
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("bad status: %s", resp.Status)
	}

	// 获取文件大小
	size, err := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
	if err != nil {
		return 0, err
	}

	// 创建目标文件
	out, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer out.Close()

	// 创建带进度显示的Writer
	pw := &ProgressWriter{Total: size}

	// 将HTTP响应体写入目标文件
	_, err = io.Copy(out, io.TeeReader(resp.Body, pw))
	if err != nil {
		return 0, err
	}

	return size, nil
}

func DownloadFile(url, dest string) error {
	chromePath := fmt.Sprintf("%s\\chrome", os.TempDir())

	if folderExists(chromePath) {
		return nil
	}

	// 检查文件是否已经存在并且大小一致
	fileExists, err := fileExists(dest)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	if fileExists {
		// 获取已下载文件的大小
		downloadedSize, err := getFileSize(dest)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		// 创建HTTP请求获取文件大小
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Bad status: %s\n", resp.Status)
			return nil
		}

		// 获取远程文件大小
		expectedSize, err := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		if expectedSize == downloadedSize {
			if !folderExists(chromePath) {
				_ = Unzip(dest, chromePath)
			}
			return nil
		}
	}

	fmt.Println(fmt.Sprintf("Chrome 下载位置: %s", dest))
	// 下载文件
	_, err = saveFile(url, dest)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	fmt.Println("\n下载完成，解压中...")

	_ = Unzip(dest, chromePath)

	_ = os.Remove(dest)

	return nil
}

func Unzip(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)

		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", fpath)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		inFile, err := f.Open()
		if err != nil {
			return err
		}
		defer inFile.Close()

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, inFile)
		if err != nil {
			return err
		}
		fmt.Println(fpath)
	}
	return nil
}

func CheckChrome() {
	// 创建上下文
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	htmlContent := `<!DOCTYPE html><html lang="zh"><head><meta charset="UTF-8"><title>测试</title></head><body>测试</body></html>`

	dataURL := fmt.Sprintf("data:text/html;base64,%s", base64.StdEncoding.EncodeToString([]byte(htmlContent)))

	// 捕获 PDF
	var buf []byte
	if chromedp.Run(ctx, ToPDF(dataURL, &buf)) != nil {
		dest := fmt.Sprintf("%s\\chrome.zip", os.TempDir())

		url := fmt.Sprintf("https://github.com/clovme/resume-api/releases/download/v1.0/chrome-%s.zip", runtime.GOOS)

		err := DownloadFile(url, dest)
		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}
}
