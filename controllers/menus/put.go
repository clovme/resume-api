package menus

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"resume/libs"
	"resume/models"
)

// SwitchStatus 切换菜单状态
func SwitchStatus(c *gin.Context) {
	s := libs.Context(c)

	var menu models.Menus
	if err := c.ShouldBind(&menu); err != nil {
		log.Println("传入的菜单数据反序列化失败！")
		s.Msg(http.StatusBadRequest, "排序失败，请重试！")
		return
	}
	result := s.Model(&menu).Where("id = ? AND uid =  ? AND rid = ?", menu.ID, s.User.ID, s.Resume.ID).Update("is_checked", menu.IsChecked)
	// 检查更新是否成功
	if result.RowsAffected > 0 {
		s.Msg(http.StatusOK, "状态更新成功！")
		return
	}
	s.Msg(http.StatusBadRequest, "状态更新失败！")
}

// EditName 修改模块名称
func EditName(c *gin.Context) {
	s := libs.Context(c)

	var menu models.Menus
	if err := c.ShouldBind(&menu); err != nil {
		log.Println("传入的菜单数据反序列化失败！")
		s.Msg(http.StatusBadRequest, "排序失败，请重试！")
		return
	}

	if err := s.First(&menu, "uid = ? AND rid = ? AND title = ?", s.User.ID, s.Resume.ID, menu.Title).Error; err == nil {
		s.Msg(http.StatusBadRequest, "模块名相同，不能修改！")
		return
	}

	if result := s.Model(&menu).Where("uid = ? AND rid = ?", s.User.ID, s.Resume.ID).Update("title", menu.Title); result.Error != nil {
		log.Println("模块名称更新失败！", result.Error)
		s.Msg(http.StatusBadRequest, "模块名称更新失败！")
		return
	}
	s.Msg(http.StatusOK, "模块名称更新成功！")
}

// Sort 排序
func Sort(c *gin.Context) {
	libs.Sort[models.Menus]("模块名称", c, models.Menus{})
}
