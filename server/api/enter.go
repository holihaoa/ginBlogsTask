package api

import (
	"ginBlogsTask/server/api/comment"
	"ginBlogsTask/server/api/post"
	"ginBlogsTask/server/api/user"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	UserApiGroup    user.ApiGroup
	PostApiGroup    post.ApiGroup
	CommentApiGroup comment.ApiGroup
}
