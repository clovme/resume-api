package libs

import (
	"crypto/md5"
	"encoding/hex"
	"gopkg.in/ini.v1"
	"log"
	"os"
)

func CreateDir(folderPath string) {
	// 检查文件夹是否存在
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 文件夹不存在，创建文件夹
		err := os.Mkdir(folderPath, os.ModePerm)
		if err != nil {
			log.Fatalf("创建文件夹失败: %v\n", err)
		}
	} else if err != nil {
		// 处理其他可能的错误
		log.Fatalf("检查文件夹失败: %v\n", err)
	}
}

func WriteConfig(host, port, username, password, database string) {
	// 创建一个新的 INI 文件对象
	cfg := ini.Empty()
	cfg.Section("mysql").Key("host").SetValue(host)
	cfg.Section("mysql").Key("port").SetValue(port)
	cfg.Section("mysql").Key("username").SetValue(username)
	cfg.Section("mysql").Key("password").SetValue(password)
	cfg.Section("mysql").Key("database").SetValue(database)

	// 保存到文件
	err := cfg.SaveTo("config.ini")
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}

	log.Println("INI file written successfully")
}

func MD5(str string) string {
	// 计算MD5哈希值
	hash := md5.Sum([]byte(str))

	// 将哈希值转换为十六进制字符串
	return hex.EncodeToString(hash[:])
}
