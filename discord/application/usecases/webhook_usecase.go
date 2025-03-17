package usecases

import (
	"fmt"
	"webhook_multi/discord/domain/entities"
	"webhook_multi/discord/infraestructure/adapters"
)

type WebhookUsecase struct {
    discordAdapter *adapters.DiscordAdapter
}

func NewWebhookUsecase(discordAdapter *adapters.DiscordAdapter) *WebhookUsecase {
    return &WebhookUsecase{discordAdapter: discordAdapter}
}

func (uc *WebhookUsecase) ProcessGitHubEvent(event entities.GitHubEvent) (string, string) {
    channel, message := uc.getChannelAndMessage(event)
    if channel != "" {
        uc.discordAdapter.SendMessage(channel, message)
    }
    return channel, message
}

func (uc *WebhookUsecase) getChannelAndMessage(event entities.GitHubEvent) (string, string) {
    if event.Action == "opened" {
        return "Desarrollo", formatPRMessage(event)
    }

    if event.Action == "ready_for_review" {
        return "Desarrollo", formatReviewMessage(event)
    }

    if event.Action == "reopened" {
        return "Desarrollo", formatReopenedMessage(event)
    }

    if event.Action == "synchronize" { // Cuando se hace un push a un PR
        return "Desarrollo", formatPushMessage(event)
    }

    if event.WorkflowRun != nil { // Para eventos de GitHub Actions
        return "Pruebas", formatActionsMessage(event)
    }

    return "", ""
}



func formatPRMessage(event entities.GitHubEvent) string {
    return fmt.Sprintf("📌 **[Pull Request] Nueva actividad**\n\n📝 **Título:** %s\n🔗 [Ver PR](%s)",
        event.PullRequest.Title, event.PullRequest.URL)
}

func formatPushMessage(event entities.GitHubEvent) string {
    return fmt.Sprintf("🚀 **[Push] Se han subido nuevos cambios**\n\n🔗 [Ver cambios](%s)", event.PullRequest.URL)
}

func formatReviewMessage(event entities.GitHubEvent) string {
    return fmt.Sprintf("🧐 **[Review] Un Pull Request está listo para revisión**\n\n📝 **Título:** %s\n🔗 [Revisar PR](%s)",
        event.PullRequest.Title, event.PullRequest.URL)
}

func formatReopenedMessage(event entities.GitHubEvent) string {
    return fmt.Sprintf("🔄 **[Reabierto] Un Pull Request ha sido reabierto**\n\n📝 **Título:** %s\n🔗 [Abrir PR](%s)",
        event.PullRequest.Title, event.PullRequest.URL)
}

func formatActionsMessage(event entities.GitHubEvent) string {
    conclusion := event.WorkflowRun.Conclusion
    if conclusion == "" {
        conclusion = "Sin conclusión"
    }

    return fmt.Sprintf("⚙️ **[GitHub Actions] Workflow ejecutado**\n\n✅ **Estado:** %s\n📌 **Conclusión:** %s\n🔗 [Ver detalles](%s)",
        event.WorkflowRun.Status, conclusion, event.WorkflowRun.URL)
}

