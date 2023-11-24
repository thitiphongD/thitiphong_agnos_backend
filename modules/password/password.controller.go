package password

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thitiphongD/thitiphong_agnos_backend/helpers"
	"github.com/thitiphongD/thitiphong_agnos_backend/requests"
)

func StrongPasswordStepsController(c *gin.Context) {
	requestBody := &requests.RequestPassword{}
	if err := c.BindJSON(requestBody); err != nil {
		if err.Error() == "EOF" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "body must be json",
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	numOfSteps := helpers.PasswordStrengthSteps(requestBody.InitPassword)

	c.JSON(http.StatusOK, gin.H{
		"num_of_steps": numOfSteps,
	})
}
