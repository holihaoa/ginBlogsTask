package user

import "github.com/gin-gonic/gin"

type UserBaseRouter struct{}

func (u *UserBaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	userBase := Router.Group("user")
	{
		userBase.POST("login", baseApi.Login)
		userBase.POST("registry", baseApi.Registry)
	}
}
