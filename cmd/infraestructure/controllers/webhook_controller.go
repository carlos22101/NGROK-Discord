package controllers

import (
	"log"
	"net/http"
	"webhook_multi/cmd/application/usecases"
	"webhook_multi/cmd/domain/entities"

	"github.com/gin-gonic/gin"
)

type WebhookController struct {
	webhookUsecase *usecases.WebhookUsecase
}

func NewWebhookController(usecase *usecases.WebhookUsecase) *WebhookController {
	return &WebhookController{webhookUsecase: usecase}
}

func (wc *WebhookController) HandleWebhook(c *gin.Context) {
    var event entities.GitHubEvent

    // Leer el JSON de GitHub
    if err := c.ShouldBindJSON(&event); err != nil {
        log.Println("❌ Error al parsear JSON:", err) // <-- Agrega este log
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    log.Println("✅ Webhook recibido:", event) // <-- Agrega este log

    channel, message := wc.webhookUsecase.ProcessGitHubEvent(event)
    if channel != "" {
        c.JSON(http.StatusOK, gin.H{"channel": channel, "message": message})
    } else {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo procesar el evento"})
    }
}

