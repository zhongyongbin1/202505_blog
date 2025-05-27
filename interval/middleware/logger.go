package middleware

import (
	"context"
	"gin-blog-newest/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LoggerMiddleware struct {
	// 这里可以添加一些日志相关的配置

}

func ApiLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 生成请求ID
		requestID := uuid.New().String()

		// 添加上下文信息
		reqLogger := logger.Get().WithFields(map[string]interface{}{
			"request_id": requestID,
			"client_ip":  c.ClientIP(),
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
		})

		// 将logger放入上下文
		ctx := reqLogger.WithContext(c.Request.Context())
		ctx = context.WithValue(ctx, "request_id", requestID)
		c.Request = c.Request.WithContext(ctx)

		// 记录请求开始
		reqLogger.Info().Msg("Request started")

		// 执行请求处理
		c.Next()

		// 记录请求结束
		reqLogger.Info().
			Int("status", c.Writer.Status()).
			Msg("Request completed")
	}
}
