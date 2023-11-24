package password

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thitiphongD/thitiphong_agnos_backend/helpers"

	// "github.com/thitiphongD/thitiphong_agnos_backend/middlewares"
	"github.com/thitiphongD/thitiphong_agnos_backend/requests"
)

func InitPassword(r *gin.Engine) {
	// r.Use(middlewares.LoggerMiddleware())
	r.POST("/api/strong_password_steps", func(c *gin.Context) {
		
		body := &requests.RequestPassword{}
		if err := c.ShouldBindJSON(body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		log.Printf("Request Body: %+v\n", body)

		log.Printf(
			"ClientIP: %s, TimeStamp: %s, Method: %s, Path: %s, Proto: %s, StatusCode: %d, Latency: %s, UserAgent: %s, ErrorMessage: %s\n",
			c.ClientIP(),
			time.Now().Format(time.RFC1123),
			c.Request.Method,
			c.Request.URL.Path,
			c.Request.Proto,
			c.Writer.Status(),
			c.Writer.Header().Get("Content-Length"),
			c.Request.UserAgent(),
			c.Errors.ByType(gin.ErrorTypePrivate).String(),
		)

		numOfSteps := helpers.PasswordStrengthSteps(body.InitPassword)
		responseBody := gin.H{"num_of_steps": numOfSteps}
		log.Println(toJsonString(responseBody))
		c.JSON(http.StatusOK, responseBody)
		
		// c.JSON(200, gin.H{
		// 	"num_of_steps": numOfSteps,
		// })
	})
}

func toJsonString(data gin.H) string {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Println("Error converting to JSON:", err)
		return ""
	}
	return string(jsonBytes)
}
