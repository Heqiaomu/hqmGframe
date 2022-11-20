package mresp

// 用户返回的类型

type UserLogin struct {
	UserID    string `json:"userID"`
	Username  string `json:"username"`
	Token     string `json:"token"`
	LoginTime string `json:"loginTime"`
	ClientIP  string `json:"clientIP"`
}
