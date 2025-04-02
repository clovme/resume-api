package enums

import (
	"log"
	"os"
	"path/filepath"
)

var (
	/*
		TempPath      = os.TempDir()
		ChromeZipPath = filepath.Join(TempPath, "chrome.7z")
		ChromeUrl     = "https://gitee.com/lovme/tool-kit/releases/download/%s/chrome.7z"
	*/

	DataPath      = ""
	DatabasePath  = "data/resume.db"
	DatabasePath1 = "data/resume.db"
	ConfigPath    = "data/config.ini"
	TempPath      = filepath.Join(os.TempDir(), "~resume-tmp")
	EdgeExePath   = filepath.Join("C:\\", "Program Files (x86)", "Microsoft", "Edge", "Application", "msedge.exe")
	ChromeExePath = filepath.Join("C:\\", "Program Files", "Google", "Chrome", "Application", "chrome.exe")
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
	DatabasePath1 = filepath.Join(DataPath, "resume.db")
}
