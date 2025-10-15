package comment

import "ginBlogsTask/server/api"

type RouterGroup struct {
	CommentRouter
}

var (
	commentApi = api.ApiGroupApp.CommentApiGroup.CommentApi
)
