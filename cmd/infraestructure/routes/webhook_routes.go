package routes

import (
	"webhook_multi/cmd/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, webhookController *controllers.WebhookController) {
	router.POST("/webhook", webhookController.HandleWebhook)
}
