package password

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thitiphongD/thitiphong_agnos_backend/helpers"
	"github.com/thitiphongD/thitiphong_agnos_backend/requests"
)

func NewHTTPPassword(r *gin.Engine) {
	r.POST("/api/strong_password_steps", func(c *gin.Context) {
		requestBody := &requests.RequestPassword{}
		if err := c.BindJSON(requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		numOfSteps := helpers.PasswordStrengthSteps(requestBody.InitPassword)

		c.JSON(http.StatusOK, gin.H{
			"num_of_steps": numOfSteps,
		})

	})
}
