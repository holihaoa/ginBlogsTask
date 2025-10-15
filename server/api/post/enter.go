package post

import "ginBlogsTask/server/service"

type ApiGroup struct {
	PostApi
}

var (
	postService = service.ServiceGroupApp.PostServiceGroup.PostService
)
