package router

import (
	"ginBlogsTask/server/router/comment"
	"ginBlogsTask/server/router/post"
	"ginBlogsTask/server/router/user"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	User    user.RouterGroup
	Post    post.RouterGroup
	Comment comment.RouterGroup
}
