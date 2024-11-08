package enums

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

var (
	DataPath      = ""
	TempPath      = os.TempDir()
	ChromeZipPath = filepath.Join(TempPath, "chrome.7z")
	ChromeUrl     = "https://gitee.com/lovme/tool-kit/releases/download/%s/chrome.7z"
	ChromeExePath = filepath.Join(filepath.Dir(TempPath), "Google", "Chrome", "Application", "chrome.exe")
)

func init() {
	// 获取当前执行文件的绝对路径
	execPath, err := os.Executable()
	if err != nil {
		fmt.Println("获取文件路径出错:", err)
		return
	}

	ChromeUrl = fmt.Sprintf(ChromeUrl, runtime.GOOS)
	// 获取文件所在的目录
	DataPath = filepath.Join(filepath.Dir(execPath), "data")
}
