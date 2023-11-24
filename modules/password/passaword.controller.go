package password

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	database "github.com/thitiphongD/thitiphong_agnos_backend/db"
	"github.com/thitiphongD/thitiphong_agnos_backend/helpers"
	"github.com/thitiphongD/thitiphong_agnos_backend/requests"
)

type Logger struct {
	ID         uint   `gorm:"primarykey"`
	Request    string
	Response   string
	ClientIP   string
	Method     string
	Path       string
	Proto      string
	StatusCode int   
	Latency    string
	UserAgent  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func InitPassword(r *gin.Engine) {
	db := database.InitDB()
	db.AutoMigrate(&Logger{})

	r.POST("/api/strong_password_steps", func(c *gin.Context) {
		start := time.Now()
		requestBody := &requests.RequestPassword{}

		if err := c.ShouldBindJSON(requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		latency := time.Since(start)
		numOfSteps := helpers.PasswordStrengthSteps(requestBody.InitPassword)

		responseBody := gin.H{
			"num_of_steps": numOfSteps,
		}

		requestBodyJSON, err := json.Marshal(requestBody)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to marshal request body",
            })
            return
        }
	
		logger := Logger{
			Request:    string(requestBodyJSON),
			Response:   toJsonString(responseBody),
			ClientIP:   c.ClientIP(),
			Method:     c.Request.Method,
			Path:       c.Request.URL.Path,
			Proto:      c.Request.Proto,
			StatusCode: c.Writer.Status(),
			Latency:    latency.String(),
			UserAgent:  c.Request.UserAgent(),
		}

		result := db.Create(&logger)
		if result.Error != nil {
			panic("failed to create logger")
		}

		c.JSON(http.StatusOK, responseBody)
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
