package user

import "ginBlogsTask/server/api"

type RouterGroup struct {
	UserBaseRouter
}

var (
	baseApi = api.ApiGroupApp.UserApiGroup.BaseApi
)
