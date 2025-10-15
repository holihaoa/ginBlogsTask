package comment

import "github.com/gin-gonic/gin"

/*评论功能
实现评论的创建功能，已认证的用户可以对文章发表评论。
实现评论的读取功能，支持获取某篇文章的所有评论列表。*/

type CommentRouter struct{}

func (*CommentRouter) InitCommentRouter(Router *gin.RouterGroup) {
	commentRouter := Router.Group("comment")
	{
		commentRouter.POST("create", commentApi.Create)
		commentRouter.POST("search", commentApi.Search)
	}
}
