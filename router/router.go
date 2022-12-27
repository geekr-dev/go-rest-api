package router

import (
	"net/http"

	_ "github.com/geekr-dev/go-rest-api/docs"
	"github.com/gin-contrib/pprof"

	"github.com/geekr-dev/go-rest-api/handler/sd"
	"github.com/geekr-dev/go-rest-api/handler/user"
	"github.com/geekr-dev/go-rest-api/router/middleware"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Load(g *gin.Engine, m ...gin.HandlerFunc) *gin.Engine {
	// 中间件
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(m...)

	// 404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "错误的 API 路由")
	})

	// swagger api docs
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 对路由进行性能分析
	pprof.Register(g)

	// auth
	g.POST("/login", user.Login)

	// user
	u := g.Group("/user")
	u.Use(middleware.Authenticate())
	{
		u.POST("", user.Create)       // 创建用户
		u.DELETE("/:id", user.Delete) // 删除用户
		u.PUT("/:id", user.Update)    // 更新用户
		u.GET("", user.List)          // 用户列表
		u.GET("/:username", user.Get) // 获取用户信息
	}

	// 健康检查
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
