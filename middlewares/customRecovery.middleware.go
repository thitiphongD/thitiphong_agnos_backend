package middlewares

import (
	"github.com/gin-gonic/gin"
)

func CustomRecoveryMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
            c.JSON(500, gin.H{
                "error_message": recovered,
            })
        })(c)
    }
}