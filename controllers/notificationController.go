package controllers

import (
	"io"

	"github.com/alysonsz/gopportunities.git/services"
	"github.com/gin-gonic/gin"
)

type NotificationController interface {
	StreamNotifications(ctx *gin.Context)
}

type notificationController struct {
	notificationService *services.NotificationService
}

func NewNotificationController(notificationService *services.NotificationService) NotificationController {
	return &notificationController{notificationService}
}

// @Summary Subscribe to real-time job notifications
// @Description Open a Server-Sent Events (SSE) stream to receive real-time updates when new opportunities are created.
// @Tags notifications
// @Produce text/event-stream
// @Success 200 {string} string "SSE stream"
// @Router /api/v1/notifications [get]
func (ctrl *notificationController) StreamNotifications(ctx *gin.Context) {
	clientChan := make(chan string)
	ctrl.notificationService.AddClient(clientChan)

	ctx.Stream(func(w io.Writer) bool {
		if msg, ok := <-clientChan; ok {
			ctx.SSEvent("message", msg)
			return true
		}
		return false
	})
}
