package middlewares

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		ctxLog := logrus.WithFields(logrus.Fields{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
		})

		var reqBody []byte
		if c.Request.Body != nil {
			reqBody, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
		}

		ctxLog = ctxLog.WithFields(logrus.Fields{
			"request": string(reqBody),
		})
		ctxLog.Info("Request received")

		rr := newResponseRecorder(c.Writer)
		c.Writer = rr

		c.Next()

		duration := time.Since(startTime).Milliseconds()
		status := c.Writer.Status()

		ctxLog = ctxLog.WithFields(logrus.Fields{
			"status":   status,
			"response": rr.body.String(),
			"duration": duration,
		})

		if status >= http.StatusBadRequest {
			ctxLog.Error("Request completed with error")
		} else {
			ctxLog.Info("Request completed successfully")
		}
	}
}

type responseRecorder struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func newResponseRecorder(w gin.ResponseWriter) *responseRecorder {
	return &responseRecorder{
		w,
		bytes.NewBufferString(""),
	}
}
func (r responseRecorder) Write(buf []byte) (int, error) {
	r.body.Write(buf)
	return r.ResponseWriter.Write(buf)
}
