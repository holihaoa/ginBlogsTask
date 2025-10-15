package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				zap.L().Error("Panic recovered", zap.Any("error", err))
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "服务器异常"})
				c.Abort()
			}
		}()
		c.Next()
	}
}
