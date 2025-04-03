package enums

import (
	"log"
	"reflect"
	"resume/libs/capitalize"
)

type browserExePath struct {
	Edge   string
	Chrome string
}

func (b *browserExePath) Set(fieldName string, value string) {
	fieldName = capitalize.Capitalize(fieldName)

	v := reflect.ValueOf(b)

	// 获取结构体的值
	v = v.Elem()

	// 获取字段
	field := v.FieldByName(fieldName)

	// 检查字段是否存在
	if !field.IsValid() {
		log.Printf("字段 %s 不存在\n", fieldName)
		return
	}

	field.SetString(value)
}

func (b *browserExePath) Get(fieldName string) string {
	fieldName = capitalize.Capitalize(fieldName)

	v := reflect.ValueOf(b)

	// 获取结构体的值
	v = v.Elem()

	// 获取字段
	field := v.FieldByName(fieldName)

	// 检查字段是否存在
	if !field.IsValid() {
		log.Printf("%s 不存在，使用默认浏览器：%s\n", fieldName, b.Edge)
		return b.Edge
	}

	// 获取字段
	return field.String()
}
