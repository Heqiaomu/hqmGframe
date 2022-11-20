package response

import "time"

type LoginResp struct {
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	Token     string    `json:"token"`
	LoginTime time.Time `json:"login_time"`
	ClientIP  string    `json:"client_ip"`
}
