package password

import (
	"github.com/gin-gonic/gin"
)

func NewHTTPPassword(r *gin.Engine) {
	r.POST("/api/strong_password_steps", StrongPasswordStepsController)
}
