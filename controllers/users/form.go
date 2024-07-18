package users

// FUsers 用户表
type FUsers struct {
	Username        string `json:"username"`        // 用户名
	Password        string `json:"password"`        // 密码
	ConfirmPassword string `json:"confirmPassword"` // 确定密码
}
