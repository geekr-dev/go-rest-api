package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/satori/uuid"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 X-Request-Id，如果不存在则生成
		requestId := c.Request.Header.Get("X-Request-Id")
		if requestId == "" {
			uuidV4 := uuid.NewV4()
			requestId = uuidV4.String()
		}
		// 然后将其回写到上下文
		c.Set("X-Request-Id", requestId)
		// 以及 HTTP 请求头
		c.Writer.Header().Set("X-Request-Id", requestId)
		c.Next()
	}
}
