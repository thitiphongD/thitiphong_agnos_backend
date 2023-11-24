package middlewares

import (
	"github.com/gin-gonic/gin"
)

func NotFoundMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.JSON(404, gin.H{
            "error_message": "404 page not found",
        })
    }
}