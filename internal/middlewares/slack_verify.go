package middlewares

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/config"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

const (
	slackSignatureVersion = "v0"
)

func VerifySlackRequest(cfg config.SlackConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		timestamp := c.GetHeader("X-Slack-Request-Timestamp")
		signature := c.GetHeader("X-Slack-Signature")

		if timestamp == "" || signature == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Slack signature headers"})
			return
		}

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
			return
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		baseString := fmt.Sprintf("%s:%v:%s", slackSignatureVersion, timestamp, string(body))

		mac := hmac.New(sha256.New, []byte(cfg.Signing))
		mac.Write([]byte(baseString))
		expectedSignature := fmt.Sprintf("%s=%s", slackSignatureVersion, hex.EncodeToString(mac.Sum(nil)))

		if !hmac.Equal([]byte(expectedSignature), []byte(signature)) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid signature"})
			return
		}

		c.Next()
	}
}
