package router

import (
	"gee-Init/controller"
	"gee-Init/middleware"
	"gee-Init/util/logger"

	swaggerFiles "github.com/swaggo/files"

	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	//r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	//r.Use(middleware.Cors())
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middleware.Cors_gee, middleware.RefreshToken)
	//r.Use(middleware.CurrentUser()) // 暂时不用

	// 配置swagger文档
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	// 路由
	api_v1_group := r.Group("/api/v1")
	routerV1(api_v1_group)

	// 404处理
	r.NoRoute(func(c *gin.Context) {
		controller.ResponseError(c, controller.ErrorNotFound)
	})
	return r
}
