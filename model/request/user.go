package request

type Register struct {
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Code     int    `json:"code" yaml:"code"`   // 验证码
	Email    string `json:"email" yaml:"email"` // 设置邮箱
}

type Login struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Phone    int64  `json:"phone" form:"phone"`
	Code     int64  `json:"code" form:"code"`
}
