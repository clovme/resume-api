package enums

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	TempPath      = ""
	ChromeExePath = "C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe"
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

	// 获取文件所在的目录
	TempPath = filepath.Join(filepath.Dir(execPath), "data", "temp")
}
