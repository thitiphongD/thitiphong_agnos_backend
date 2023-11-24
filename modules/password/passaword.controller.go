package password

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thitiphongD/thitiphong_agnos_backend/helpers"
	"github.com/thitiphongD/thitiphong_agnos_backend/requests"
)

func InitPassword(r *gin.Engine) {
	r.POST("/api/strong_password_steps", func(c *gin.Context) {
		body := &requests.RequestPassword{}
		if err := c.ShouldBindJSON(body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		numOfSteps := helpers.PasswordStrengthSteps(body.InitPassword)
		c.JSON(200, gin.H{
			"num_of_steps": numOfSteps,
		})
	})
}
