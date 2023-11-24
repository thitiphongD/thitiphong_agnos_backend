package password

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/thitiphongD/thitiphong_agnos_backend/helpers"
	"github.com/thitiphongD/thitiphong_agnos_backend/requests"
	"log"
	"net/http"
)

func NewHTTPPassword(r *gin.Engine) {
	r.POST("/api/strong_password_steps", func(c *gin.Context) {
		//start := time.Now()
		requestBody := &requests.RequestPassword{}

		if err := c.BindJSON(requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		//latency := time.Since(start)
		numOfSteps := helpers.PasswordStrengthSteps(requestBody.InitPassword)

		//responseBody :=
		//
		//requestBodyJSON, err := json.Marshal(requestBody)
		//if err != nil {
		//	c.JSON(http.StatusInternalServerError, gin.H{
		//		"error": "Failed to marshal request body",
		//	})
		//	return
		//}
		//
		//logger := Logger{
		//	Request:    string(requestBodyJSON),
		//	Response:   toJsonString(responseBody),
		//	ClientIP:   c.ClientIP(),
		//	Method:     c.Request.Method,
		//	Path:       c.Request.URL.Path,
		//	Proto:      c.Request.Proto,
		//	StatusCode: c.Writer.Status(),
		//	Latency:    latency.String(),
		//	UserAgent:  c.Request.UserAgent(),
		//}
		//
		//result := db.Create(&logger)
		//if result.Error != nil {
		//	panic("failed to create logger")
		//}

		c.JSON(http.StatusBadGateway, gin.H{
			"num_of_steps": numOfSteps,
		})
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
