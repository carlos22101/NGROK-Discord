package adapters

import (
    "bytes"
    "encoding/json"
    "net/http"
    "log"
)

type DiscordAdapter struct {
    DevWebhookURL   string
    TestWebhookURL  string
}

func NewDiscordAdapter(devWebhookURL, testWebhookURL string) *DiscordAdapter {
    return &DiscordAdapter{
        DevWebhookURL:   devWebhookURL,
        TestWebhookURL:  testWebhookURL,
    }
}

func (d *DiscordAdapter) SendMessage(channel string, message string) error {
    var webhookURL string
    if channel == "Desarrollo" {
        webhookURL = d.DevWebhookURL
    } else if channel == "Pruebas" {
        webhookURL = d.TestWebhookURL
    } else {
        log.Println("⚠️ Canal no reconocido:", channel)
        return nil
    }

    payload := map[string]string{"content": message}
    body, _ := json.Marshal(payload)

    resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(body))
    if err != nil {
        log.Printf("❌ Error enviando mensaje a Discord: %v", err)
        return err
    }
    defer resp.Body.Close()

    log.Printf("✅ Mensaje enviado a Discord en %s: %s", channel, message)
    return nil
}
