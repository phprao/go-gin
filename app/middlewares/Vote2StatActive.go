package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Vote2StatActive() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// 执行其他中间件
		c.Next()
		// 计算耗时
		cost := time.Since(start)
		log.Println("[ Vote2StatActive ] - ", cost)
	}
}