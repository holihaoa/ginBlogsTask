package service

import (
	"ginBlogsTask/server/service/comment"
	"ginBlogsTask/server/service/post"
	"ginBlogsTask/server/service/user"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	UserServiceGroup    user.ServiceGroup
	PostServiceGroup    post.ServiceGroup
	CommentServiceGroup comment.ServiceGroup
}
