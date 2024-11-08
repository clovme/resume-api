package libs

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"resume/types/enums"
	"strconv"
)

// httpGetFileSize 获取文件大小
func getHttpFileSize(url string) int64 {
	// 创建HTTP请求获取文件大小
	resp, err := http.Head(enums.ChromeUrl)
	if err != nil {
		log.Println("创建HTTP请求获取文件大小失败，错误信息:", err)
		return -99
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("获取下载请求，响应码:", resp.Status)
		return -99
	}

	// 获取远程文件大小
	expectedSize, err := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
	if err != nil {
		log.Println("获取远程文件大小失败，错误信息:", err)
		return -99
	}

	return expectedSize
}

// downloadHttpFile 下载文件
func downloadHttpFile(url, dest, name string) (int64, error) {
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
	pw := &ProgressWriter{Total: size, Name: fmt.Sprintf("%s：%s", name, dest)}

	// 将HTTP响应体写入目标文件
	if _, err = io.Copy(out, io.TeeReader(resp.Body, pw)); err != nil {
		return 0, err
	}

	return size, nil
}
