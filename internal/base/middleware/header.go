package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func HeadersByRequestURI() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.RequestURI, "/static/") {
			c.Header("cache-control", "public, max-age=31536000")
		}
	}
}
