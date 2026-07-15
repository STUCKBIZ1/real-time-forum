package auth

import "time"

type User struct {
	ID         int    `json:"id"`
	Nickname   string `json:"nickname"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Age        string `json:"age"`
	Gender     string `json:"gender"`
	Avatar     string `json:"avatar"`
}
type SessionData struct {
	UserID    int
	SessionID string
	ExpiresAt time.Time
}
