package models

// ApplicationInfo 报考信息
type ApplicationInfo struct {
	BaseModelWithRIDUID
	Name  string `json:"name"`  // 学校名称
	Major string `json:"major"` // 所学专业
	CName string `json:"cname"` // 自定义名称
}

type CourseGrade struct {
	BaseModelWithRIDUID
	Name  string `json:"name"`  // 课程名称
	Score string `json:"score"` // 课程分数
}
