package logger

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Gin returns a middleware for logging gin http requests
func Gin(service string) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()

		fields := Map{
			"duration": end.Sub(start),
			"ip":       c.ClientIP(),
			"method":   c.Request.Method,
			"path":     c.Request.URL.Path,
			"status":   c.Writer.Status(),
		}

		Info(fmt.Sprintf("[%s][gin]", service), fields)
	}
}

func GinContext(c *gin.Context) Map {
	return Map{
		"request":  c.Request,
		"response": c.Writer,
		"errors":   c.Errors,
		"params":   c.Params,
	}
}
