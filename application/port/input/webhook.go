package input

import (
	"github.com/bwmarrin/discordgo"
	"github.com/juliofilizzola/bot_discord/application/domain"
)

type WebhookDomainService interface {
	Send(dataGit *discordgo.WebhookParams, webhookId, webhookToken, action string) string
	Save(dataGit domain.Github)
}
