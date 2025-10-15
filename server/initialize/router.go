package initialize

import (
	"ginBlogsTask/server/global"
	"ginBlogsTask/server/router"
	"ginBlogsTask/server/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrPermission
	}

	return f, nil
}

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}

	userRouter := router.RouterGroupApp.User
	postRouter := router.RouterGroupApp.Post
	commentRouter := router.RouterGroupApp.Comment
	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面3行注释
	// Router.StaticFile("/favicon.ico", "./dist/favicon.ico")
	// Router.Static("/assets", "./dist/assets")   // dist里面的静态资源
	// Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

	// 方便统一添加路由组前缀 多服务器上线使用

	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	PrivateGroup.Use(utils.JWTAuth())
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		userRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}

	{
		postRouter.InitPostRouter(PrivateGroup)       // 注册功能api路由
		commentRouter.InitCommentRouter(PrivateGroup) // jwt相关路由
	}

	global.GVA_LOG.Info("router register success")
	return Router
}
