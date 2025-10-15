package comment

import "ginBlogsTask/server/service"

type ApiGroup struct {
	CommentApi
}

var (
	commentService = service.ServiceGroupApp.CommentServiceGroup.CommentService
)
