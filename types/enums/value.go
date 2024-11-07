package enums

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

var (
	TempPath      = ""
	ChromeExePath = filepath.Join(os.TempDir(), "chrome", "chrome")
	ChromeZipPath = filepath.Join(os.TempDir(), "chrome.7z")
	ChromeUrl     = "https://gitee.com/lovme/tool-kit/releases/download/tools/chrome-%s.7z"
)

const (
	Vx32 string = "00000000000000000000000000000000"
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
	TempPath = filepath.Join(filepath.Dir(execPath), "data", "temp")
}
