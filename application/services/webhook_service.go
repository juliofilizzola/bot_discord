package services

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/juliofilizzola/bot_discord/application/port/input"
)

func NewWebhookDomainService(discord *discordgo.Session) input.WebhookDomainService {
	return &webhookDomainService{
		server: discord,
	}
}

type webhookDomainService struct {
	server *discordgo.Session
}

func (web webhookDomainService) Send(dataGit *discordgo.WebhookParams, webhookId, webhookToken, action string) string {
	if action == "opened" || action == "closed" {
		webhook, err := web.server.WebhookExecute(webhookId, webhookToken, true, dataGit)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(webhook)

		return "deu bom"
	}
	return "deu ruim!"
}
