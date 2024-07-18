package libs

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"resume/models"
	"resume/types/enums"
)

// Transaction 开启事务
// @param s HttpStatus
// @param callback 回调函数
// @param callback tx 事务对象
func Transaction(s HttpStatus, callback func(tx *gorm.DB)) {
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

// Update 更新数据
// @param msg 提示消息
// @param models 数据模型
// @param s HttpStatus
// @param Model 回调函数
// @param Model index 数据模型item索引
// @param Model model 数据模型item数据
func Update[T any](msg string, models []T, s HttpStatus, Model func(model *T) *gorm.DB) {
	status := true
	for _, model := range models {
		if result := Model(&model).Select("*").Where("rid != ? AND uid != ?", enums.Vx32, enums.Vx32).Omit("id", "rid", "uid", "sort", "CreatedAt").Updates(model); result.Error != nil {
			log.Println(fmt.Sprintf("%s数据查询异常:", msg), result.Error)
			status = false
		}
	}
	if status {
		s.Msg(http.StatusOK, "数据保存成功！")
		return
	}
	s.Msg(http.StatusNotFound, "数据保存失败，请重试！")
}

// Create 创建数据
// @param s HttpStatus
// @param msg 提示消息
// @param model 数据
// @param emptyModel 查询空模型
// @param query 查询条件
// @param args 查询参数
func Create[T any](s HttpStatus, msg string, model *T, emptyModel T, query string, args ...interface{}) {
	if result := s.Where(query, args...).First(&emptyModel); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 数据不存在，插入新记录
			if err := s.Create(&model).Error; err != nil {
				log.Println(fmt.Sprintf("%s创建失败:", msg), err)
				s.Msg(http.StatusBadRequest, "数据保存失败，请重试！")
				return
			} else {
				log.Println(fmt.Sprintf("%s创建成功:", msg), model)
				s.Json(http.StatusOK, "数据保存成功！", model)
				return
			}
		} else {
			log.Println(fmt.Sprintf("%s查询失败:", msg), result.Error)
		}
	} else if result.RowsAffected > 0 {
		s.Msg(http.StatusBadRequest, "数据已存在！")
	}
}

// SwitchTagsStatus 改变标签状态
// @param db 数据库连接对象
// @param resume 简历ID和uid
// @param tagsName 选中的标签
func SwitchTagsStatus(db *gorm.DB, resume models.Resumes, tagsName []string, tagsType enums.Tags) {
	where := db.Model(&models.Tags{}).Where("uid = ? AND rid = ? AND type = ?", resume.UID, resume.ID, tagsType)
	where.Update("is_checked", false)
	where.Where("name in ?", tagsName).Update("is_checked", true)
}

func Sort[T any](msg string, c *gin.Context, model T) {
	s := Context(c)

	var ids []string
	if err := c.ShouldBind(&ids); err != nil {
		log.Println(fmt.Sprintf("%s数据反序列化失败！", msg))
		s.Msg(http.StatusBadRequest, "排序失败，请重试！")
		return
	}

	success := 0
	Transaction(s, func(tx *gorm.DB) {
		for index, id := range ids {
			result := tx.Model(&model).Where("id = ? AND uid = ? AND rid = ?", id, s.User.ID, s.Resume.ID).Update("sort", index)
			// 检查更新是否成功
			if result.RowsAffected > 0 {
				success++
			}
		}
	})

	s.Msg(http.StatusOK, fmt.Sprintf("排序完成%d/%d条!", success, len(ids)))
}

// DeleteData 删除数据返回状态
// @param s HttpStatus
// @param msg 消息状态
// @param result ORM操作结果
func DeleteData[T any](s HttpStatus, tx *gorm.DB, msg string, model T, query string, args ...interface{}) bool {
	// 检查删除操作是否成功
	if result := tx.Where(query, args...).Delete(&model); result.Error != nil {
		s.Msg(http.StatusForbidden, fmt.Sprintf("%s删除失败！", msg))
		return false
	}
	return true
}

func QueryDataBaseArray[T any](c *gin.Context, models []T) {
	s := Context(c)

	if err := s.Order("sort asc").Find(&models, "rid = ? AND uid = ?", s.Resume.ID, s.User.ID).Error; err != nil {
		s.Msg(http.StatusNotFound, "没有查询到数据！")
		return
	}
	s.Json(http.StatusOK, "数据查询成功", models)
}
