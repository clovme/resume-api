package enums

import (
	"log"
	"os"
	"path/filepath"
)

type Tags uint

const (
	Skills  Tags = iota
	Honors  Tags = iota
	Hobbies Tags = iota
)

const (
	Vx32         = "00000000000000000000000000000000"
	DatabasePath = "data/resume.db"
)

var (
	DataPath       = ""
	ConfigPath     = "data/config.ini"
	TempPath       = filepath.Join(os.TempDir(), "~resume-tmp")
	BrowserExePath = browserExePath{
		Edge:   filepath.Join("C:\\", "Program Files (x86)", "Microsoft", "Edge", "Application", "msedge.exe"),
		Chrome: filepath.Join("C:\\", "Program Files", "Google", "Chrome", "Application", "chrome.exe"),
	}
)

func init() {
	// 获取当前执行文件的绝对路径
	execPath, err := os.Executable()
	if err != nil {
		log.Println("获取文件路径出错:", err)
		return
	}

	// 获取文件所在的目录
	DataPath = filepath.Join(filepath.Dir(execPath), "data")
	ConfigPath = filepath.Join(DataPath, "config.ini")
}
