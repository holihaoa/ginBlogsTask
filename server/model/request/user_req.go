package request

type UserReq struct {
	Username string `json:"userName" example:"用户名"`
	Password string `json:"passWord" example:"密码"`
	Email    string `json:"email" example:"电子邮箱"`
}
