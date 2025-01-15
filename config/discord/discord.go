package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/juliofilizzola/bot_discord/config/env"
	"log"
)

func StartDiscord() (*discordgo.Session, error) {
	discord, err := discordgo.New("Bot " + env.TokenDiscord)
	if err != nil {
		return discord, err
	}
	discord.AddHandler(InteractiveMessage)
	discord.AddHandlerOnce(ReadMessage)
	err = discord.Open()
	if err != nil {
		return discord, err
	}

	log.Println("Bot está rodando. Pressione CTRL+C para sair.")
	return discord, nil
}
