package auth

type LoginReq struct{
	Email string `json:"email"`
	Password string `json:"password"`
}
type RegisterReq struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Nickname string `json:"nickname"`
	Email string `json:"email"`
	Password string `json:"password"`
	Age string `json:"age"`
	Gender string `json:"gender"`
	Avatar string `json:"avatar"`
}
var session_id string