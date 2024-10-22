package domain

type Session struct {
	Token  string `json:"token" redis:"token"`
	UserID string `json:"user_id" redis:"user_id"`
}
