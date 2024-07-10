package libs

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"resume/types"
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

// Transaction 开启事务
func Transaction(s HttpStatus, callback types.Callback) {
	// 开始事务
	tx := s.Begin()
	// 事务操作
	if err := tx.Error; err != nil {
		log.Println("事务开启失败！", err.Error())
		s.Msg(http.StatusForbidden, "创建简历模版创建失败，请重试！")
		return
	}

	callback(tx)

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		log.Println("创建简历模版事务提交失败！", err.Error())
		s.Msg(http.StatusForbidden, "创建简历模版创建失败，请重试！")
		return
	}
}

func DeleteData(s HttpStatus, msg string, result *gorm.DB) bool {
	// 检查删除操作是否成功
	if result.Error != nil {
		s.Msg(http.StatusForbidden, fmt.Sprintf("%s删除失败！", msg))
		return false
	} else if result.RowsAffected == 0 {
		s.Msg(http.StatusForbidden, fmt.Sprintf("%s没有找到符合条件的数据！", msg))
		return false
	}
	return true
}
