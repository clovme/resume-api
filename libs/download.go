package libs

import (
	"embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"resume/types/enums"
	"strings"
)

// 检查文件是否存在
func isFileExists(filePath string) bool {
	// 获取文件状态
	info, err := os.Stat(filePath)

	// 文件不存在且不是文件夹
	if os.IsNotExist(err) || (err == nil && info.IsDir()) {
		return false
	}
	return true
}

// 检查文件大小
func getFileSize(filepath string) (int64, error) {
	info, err := os.Stat(filepath)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

// DownloadChromeFile 下载Chrome文件
func DownloadChromeFile(static embed.FS) error {
	chromePath := filepath.Dir(enums.ChromeExePath)

	if isFileExists(enums.ChromeExePath) {
		return nil
	}

	// 检查文件是否已经存在并且大小一致
	if isFileExists(enums.ChromeZipPath) {
		// 获取本地文件的大小
		downloadedSize, err := getFileSize(enums.ChromeZipPath)
		if err != nil {
			log.Println("获取已下载文件的大小失败，错误信息:", err)
			return err
		}

		if size := getHttpFileSize(enums.ChromeUrl); size == downloadedSize {
			extract7z(enums.ChromeZipPath, chromePath)
			return nil
		}
	}
	// 下载文件
	if _, err := downloadHttpFile(enums.ChromeUrl, enums.ChromeZipPath, "下载Chrome浏览器"); err != nil {
		fmt.Println("文件下载失败:", enums.ChromeUrl)
		return err
	}
	fmt.Println("\n下载完成，准备解压...")
	// 执行解压
	extract7z(enums.ChromeZipPath, chromePath)
	_ = os.Remove(enums.ChromeZipPath)
	return nil
}

// 解压 7z 文件到指定目录
func extract7z(archivePath, destDir string) {
	var files []string
	for _, ext := range []string{"7z.exe", "7z.dll"} {
		ext7zPath := filepath.Join(enums.TempPath, ext)
		if isFileExists(ext7zPath) {
			continue
		}

		newUrl := strings.Replace(enums.ChromeUrl, "chrome.7z", ext, 1)
		if _, err := downloadHttpFile(newUrl, ext7zPath, "下载解压工具"); err != nil {
			fmt.Printf("%s文件下载失败:%s\n", ext, newUrl)
			return
		}
		files = append(files, ext7zPath)
		fmt.Println("")
	}

	// 执行 7z 解压命令
	cmd := exec.Command(filepath.Join(enums.TempPath, "7z.exe"), "x", archivePath, "-o"+destDir, "-bb3")

	// 设置命令的标准输出和标准错误输出为控制台
	cmd.Stdout = os.Stdout // 将标准输出重定向到控制台
	cmd.Stderr = os.Stderr // 将标准错误输出重定向到控制台

	// 执行命令并返回结果
	if err := cmd.Run(); err != nil {
		fmt.Println("解压失败:", err)
	}

	for _, file := range files {
		_ = os.Remove(file)
	}
}
