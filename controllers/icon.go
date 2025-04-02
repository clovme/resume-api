package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func GetIcon(c *gin.Context) {
	// 打开文件，以只读方式打开
	file, err := os.Open("D:\\develop\\vue\\resume-web\\src\\styles\\webfont\\variables.scss")
	if err != nil {
		log.Printf("Failed to open file: %v\n", err)
		return
	}
	defer file.Close()

	// 读取文件内容
	data := make([]byte, 1024) // 假设文件不超过1024字节
	count, err := file.Read(data)
	if err != nil {
		log.Printf("Failed to read file: %v\n", err)
		return
	}

	log.Printf("Read %d bytes: %s\n", count, data[:count])
	c.String(http.StatusOK, string(data))
}
