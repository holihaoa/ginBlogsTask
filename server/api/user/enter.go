package user

import "ginBlogsTask/server/service"

type ApiGroup struct {
	BaseApi
}

var (
	userService = service.ServiceGroupApp.UserServiceGroup.UserService
)
