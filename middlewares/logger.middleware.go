package middlewares

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thitiphongD/thitiphong_agnos_backend/db"
	"github.com/thitiphongD/thitiphong_agnos_backend/models"
)

type customResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *customResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *customResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		customWriter := &customResponseWriter{
			ResponseWriter: c.Writer,
			body:           new(bytes.Buffer),
		}

		c.Writer = customWriter
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println("failed to read body")
		}

		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		c.Next()

		statusCode := c.Writer.Status()
		latency := time.Since(startTime)

		err = db.Database.Create(&models.Logger{
			Request:    string(bodyBytes),
			Response:   customWriter.body.String(),
			StatusCode: statusCode,
			ClientIP:   c.ClientIP(),
			Method:     c.Request.Method,
			Path:       c.Request.URL.Path,
			Proto:      c.Request.Proto,
			UserAgent:  c.Request.UserAgent(),
			Latency:    latency.String(),
		}).Error
		if err != nil {
			fmt.Println("failed to create logger")
		}
	}

}
