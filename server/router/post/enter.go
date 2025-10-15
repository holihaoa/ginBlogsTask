package post

import "ginBlogsTask/server/api"

type RouterGroup struct {
	PostBaseRouter
}

var (
	postApi = api.ApiGroupApp.PostApiGroup.PostApi
)
