package comment

import (
	"ginBlogsTask/server/global"
	"ginBlogsTask/server/model"
	"ginBlogsTask/server/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CommentApi struct{}

func (g *CommentApi) Create(c *gin.Context) {
	var com model.Comment
	c.ShouldBindJSON(&com)
	value, _ := c.Get("userid")
	com.UserID = uint(value.(float64))
	comment := &model.Comment{Content: com.Content, PostID: com.PostID, UserID: com.UserID}
	err := commentService.Create(comment)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}
func (g *CommentApi) Search(c *gin.Context) {
	var commentReq model.Comment
	c.ShouldBindJSON(&commentReq)
	comment := &model.Comment{PostID: commentReq.PostID}
	rep, err := commentService.Search(comment)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithDetailed(rep, "查询失败", c)
		return
	}
	response.OkWithDetailed(rep, "查询成功", c)
}
