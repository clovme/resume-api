package libs

import (
	"embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"resume/types/enums"
	"runtime"
	"strconv"
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

func isFileExists(filePath string) bool {
	// 获取文件状态
	info, err := os.Stat(filePath)

	// 文件不存在且不是文件夹
	if os.IsNotExist(err) || (err == nil && info.IsDir()) {
		return false
	}
	return true
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

func DownloadChromeFile(static embed.FS) error {
	chromePath := filepath.Dir(enums.ChromeExePath)

	if isFileExists(enums.ChromeExePath) {
		return nil
	}

	// 检查文件是否已经存在并且大小一致
	if isFileExists(enums.ChromeZipPath) {
		// 获取已下载文件的大小
		downloadedSize, err := getFileSize(enums.ChromeZipPath)
		if err != nil {
			log.Println("获取已下载文件的大小失败，错误信息:", err)
			return err
		}

		// 创建HTTP请求获取文件大小
		resp, err := http.Head(enums.ChromeUrl)
		if err != nil {
			log.Println("创建HTTP请求获取文件大小失败，错误信息:", err)
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Println("获取下载请求，响应码:", resp.Status)
			return nil
		}

		// 获取远程文件大小
		expectedSize, err := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
		if err != nil {
			log.Println("获取远程文件大小失败，错误信息:", err)
			return err
		}

		if expectedSize == downloadedSize {
			extract7z(enums.ChromeZipPath, chromePath, static)
			return nil
		}
	}

	fmt.Println(fmt.Sprintf("正在下载缺少的Chrome浏览器，位置: %s", enums.ChromeZipPath))
	// 下载文件
	_, err := saveFile(enums.ChromeUrl, enums.ChromeZipPath)
	if err != nil {
		log.Println("文件下载失败，错误信息:", err)
		return err
	}

	fmt.Println("\n下载完成，解压中...")

	// 调用封装好的解压方法
	extract7z(enums.ChromeZipPath, chromePath, static)

	_ = os.Remove(enums.ChromeZipPath)

	return nil
}

// 解压 7z 文件到指定目录
func extract7z(archivePath, destDir string, static embed.FS) {
	localPath := filepath.Dir(enums.ChromeZipPath)
	files := fsSaveFiles(static, fmt.Sprintf("public/%s", runtime.GOOS), localPath)

	// 执行 7z 解压命令
	log.Println(fmt.Sprintf("解压命令：%s x %s -o%s -bb3", filepath.Join(localPath, "7z.exe"), archivePath, destDir))
	cmd := exec.Command(filepath.Join(localPath, "7z.exe"), "x", archivePath, "-o"+destDir, "-bb3")

	// 设置命令的标准输出和标准错误输出为控制台
	cmd.Stdout = os.Stdout // 将标准输出重定向到控制台
	cmd.Stderr = os.Stderr // 将标准错误输出重定向到控制台

	// 执行命令并返回结果
	err := cmd.Run()
	if err != nil {
		log.Println("解压失败:", err)
	}

	for _, file := range files {
		_ = os.Remove(file)
	}

	_ = files
}

// fsSaveFiles 遍历嵌入的 public/windows 文件夹并将文件保存到本地目录
func fsSaveFiles(fsys embed.FS, dir, localPath string) []string {
	// 读取嵌入文件夹中的所有文件
	files, err := fsys.ReadDir(dir)
	if err != nil {
		log.Println("读取嵌入文件夹失败:", err)
	}

	var destFilePaths []string

	// 遍历文件夹中的每个文件
	for _, file := range files {
		// 获取文件的嵌入路径和本地保存路径
		embedFilePath := fmt.Sprintf("%s/%s", dir, file.Name()) // 使用 Unix 风格路径
		destFilePath := filepath.Join(localPath, file.Name())

		destFilePaths = append(destFilePaths, destFilePath)
		// 读取嵌入的文件内容
		data, err := fsys.ReadFile(embedFilePath)
		if err != nil {
			log.Println("读取文件失败:", err)
		}

		// 保存文件到本地
		err = os.WriteFile(destFilePath, data, 0644)
		if err != nil {
			log.Println("保存文件失败:", err)
		}
	}
	return destFilePaths
}
