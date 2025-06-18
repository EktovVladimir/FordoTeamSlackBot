package management_handler

import (
	"context"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetupRoutes(ctx context.Context, router *gin.RouterGroup) {
	router.Use(func(c *gin.Context) {
		c.Request = c.Request.WithContext(h.contextWithApi(ctx))
		c.Next()
	})

	h.userHandler.SetupRoutes(router)
	h.settingHandler.SetupRoutes(router)
	h.deploymentsHandler.SetupRoutes(router)
	h.codeReviewsHandler.SetupRoutes(router)
}
