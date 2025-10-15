package user

import (
	"ginBlogsTask/server/global"
	"ginBlogsTask/server/model"
	"ginBlogsTask/server/model/request"
	"ginBlogsTask/server/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BaseApi struct{}

func (b *BaseApi) Registry(c *gin.Context) {
	var r request.UserReq
	c.ShouldBindJSON(&r)
	user := &model.User{Username: r.Username, Password: r.Password, Email: r.Email}
	userReturn, err := userService.Register(*user)
	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(response.UserResponse{User: userReturn}, "注册失败", c)
		return
	}
	response.OkWithDetailed(response.UserResponse{User: userReturn}, "注册成功", c)
}

func (b *BaseApi) Login(c *gin.Context) {
	var r request.UserReq
	c.ShouldBindJSON(&r)
	user := &model.User{Username: r.Username, Password: r.Password}
	userReturn, err := userService.Login(*user)
	if err != nil {
		global.GVA_LOG.Error("登录失败!", zap.Error(err))
		response.FailWithDetailed(userReturn, "登录失败", c)
		return
	}
	response.OkWithDetailed(userReturn, "登录成功", c)
}
