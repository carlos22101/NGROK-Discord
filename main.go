// main.go
package main

import (
	"os"
	"webhook_multi/cmd/application/usecases"
	"webhook_multi/cmd/infraestructure/adapters"
	"webhook_multi/cmd/infraestructure/controllers"
	"webhook_multi/cmd/infraestructure/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
  
    if err := godotenv.Load(); err != nil {
        log.Println("No se pudo cargar el archivo .env")
    }

   
    webhookDev := os.Getenv("WEBHOOK_DEV")
    webhookTest := os.Getenv("WEBHOOK_TEST")

    if webhookDev == "" || webhookTest == "" {
        log.Fatal("Las variables de entorno WEBHOOK_DEV y/o WEBHOOK_TEST no están definidas")
    }

    r := gin.Default()

    discordAdapter := adapters.NewDiscordAdapter(webhookDev, webhookTest)
    webhookUsecase := usecases.NewWebhookUsecase(discordAdapter)
    webhookController := controllers.NewWebhookController(webhookUsecase)

    routes.RegisterRoutes(r, webhookController)

    r.Run(":8080")
}
