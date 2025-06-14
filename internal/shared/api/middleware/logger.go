package middleware

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logrus.Errorf("Request %s %s failed %v", r.Method, r.RequestURI, err)
			}
		}()

		reqContent, _ := io.ReadAll(r.Body)

		ctxLog := logrus.WithField("msgData", string(reqContent))
		ctxLog.Infof("Request %s %s", r.Method, r.RequestURI)

		rr := newResponseRecorder(w)

		startTime := time.Now()

		next.ServeHTTP(rr, r)

		duration := time.Since(startTime).Milliseconds()
		resContent, _ := io.ReadAll(rr.body)

		ctxLog = logrus.WithField("msgData", string(resContent))
		ctxLog.Infof("Response %s %s. Duration=%v", r.Method, r.RequestURI, duration)
	})
}

type responseRecorder struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func newResponseRecorder(w http.ResponseWriter) *responseRecorder {
	return &responseRecorder{
		w,
		new(bytes.Buffer),
	}
}
func (r responseRecorder) Write(buf []byte) (int, error) {
	r.body.Write(buf)
	return r.ResponseWriter.Write(buf)
}
