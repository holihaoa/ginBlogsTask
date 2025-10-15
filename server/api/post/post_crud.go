package post

import (
	"ginBlogsTask/server/global"
	"ginBlogsTask/server/model"
	"ginBlogsTask/server/model/request"
	"ginBlogsTask/server/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PostApi struct{}

func (p *PostApi) Create(c *gin.Context) {
	var postReq request.Post
	zap.L().Info("PostApi Create")
	c.ShouldBindJSON(&postReq)
	value, _ := c.Get("userid")
	post := &model.Post{Content: postReq.Content, Title: postReq.Title, UserID: uint(value.(float64))}
	err := postService.Create(post)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (p *PostApi) SearchAll(c *gin.Context) {
	rep, err := postService.SearchAll()
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithDetailed(rep, "查询成功", c)
}

func (p *PostApi) Search(c *gin.Context) {
	var postReq request.Post
	c.ShouldBindJSON(&postReq)
	post := &model.Post{Model: gorm.Model{ID: postReq.PostId}}
	rep, err := postService.Search(post)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithDetailed(rep, "查询失败", c)
		return
	}
	response.OkWithDetailed(rep, "查询成功", c)
}

func (p *PostApi) Update(c *gin.Context) {
	var postReq request.Post
	c.ShouldBindJSON(&postReq)
	value, _ := c.Get("userid")
	post := &model.Post{Content: postReq.Content, Title: postReq.Title, UserID: uint(value.(float64)), Model: gorm.Model{ID: postReq.PostId}}
	err := postService.Update(post)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (p *PostApi) Delete(c *gin.Context) {
	var postReq request.Post
	c.ShouldBindJSON(&postReq)
	value, _ := c.Get("userid")
	post := &model.Post{UserID: uint(value.(float64)), Model: gorm.Model{ID: postReq.PostId}}
	err := postService.Delete(post)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
