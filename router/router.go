package router

import (
	"net/http"

	"github.com/geekr-dev/go-rest-api/handler/sd"
	"github.com/geekr-dev/go-rest-api/router/middleware"
	"github.com/gin-gonic/gin"
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
