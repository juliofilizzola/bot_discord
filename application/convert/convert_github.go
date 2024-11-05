package convert

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/juliofilizzola/bot_discord/application/domain"
	"github.com/juliofilizzola/bot_discord/config/env"
	"strconv"
	"time"
)

func ConvertGithubToDiscord(data *domain.Github) discordgo.WebhookParams {
	var reviews []string
	var assignees []string

	if len(data.PullRequest.RequestedReviewers) > 0 {
		for _, value := range data.PullRequest.RequestedReviewers {
			reviews = append(reviews, value.Login)
		}
	}

	if len(data.PullRequest.Assignees) > 0 {
		for _, value := range data.PullRequest.Assignees {
			assignees = append(assignees, value.Login)
		}
	}

	embed := &discordgo.MessageEmbed{
		URL:         data.PullRequest.HtmlUrl,
		Type:        discordgo.EmbedTypeLink,
		Title:       data.PullRequest.Title,
		Description: data.PullRequest.Body,
		Timestamp:   time.Now().Format(time.RFC3339),
		Footer: &discordgo.MessageEmbedFooter{
			Text:         data.Organization.Login,
			IconURL:      data.Organization.AvatarUrl,
			ProxyIconURL: "",
		},
		Image: &discordgo.MessageEmbedImage{
			URL:      data.Organization.AvatarUrl,
			ProxyURL: "",
			Width:    280,
			Height:   20,
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL:      data.Organization.AvatarUrl,
			ProxyURL: "",
			Width:    280,
			Height:   20,
		},
		Author: &discordgo.MessageEmbedAuthor{
			URL:          data.PullRequest.User.HtmlUrl,
			Name:         data.PullRequest.User.Login,
			IconURL:      data.PullRequest.User.AvatarUrl,
			ProxyIconURL: "",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Branch:",
				Value:  data.PullRequest.Head.Ref,
				Inline: false,
			},
			{
				Name:   "Merge into:",
				Value:  fmt.Sprintf("%s from %s", data.PullRequest.Base.Ref, data.PullRequest.Head.Ref),
				Inline: false,
			},
			{
				Name:   "Status:",
				Value:  data.PullRequest.State,
				Inline: false,
			},
			{
				Name: "Assinado:",
				Value: func() string {
					if len(data.PullRequest.Assignee.Login) == 0 {
						return "Não teve assinatura"
					}
					return data.PullRequest.Assignee.Login
				}(),
				Inline: false,
			},
			{
				Name:   "Codigo adicionado:",
				Value:  strconv.Itoa(data.PullRequest.Additions),
				Inline: true,
			},
			{
				Name:   "Codigo deletado",
				Value:  strconv.Itoa(data.PullRequest.Deletions),
				Inline: true,
			},
			{
				Name:   "Commits:",
				Value:  fmt.Sprintf("[commits](%s/commits)", data.PullRequest.HtmlUrl),
				Inline: false,
			},
			{
				Name:   "Reviews",
				Value:  fmt.Sprintf("%v", reviews),
				Inline: false,
			},
		},
	}

	return discordgo.WebhookParams{
		Content:    "Nova PR no Repositorio: " + data.Repository.Name,
		Username:   env.Username,
		AvatarURL:  env.AvatarURL,
		TTS:        false,
		Files:      nil,
		Components: nil,
		Embeds:     []*discordgo.MessageEmbed{embed},
		AllowedMentions: &discordgo.MessageAllowedMentions{
			Parse: []discordgo.AllowedMentionType{
				discordgo.AllowedMentionTypeEveryone,
			},
			Roles:       nil,
			Users:       nil,
			RepliedUser: false,
		},
		Flags: 0,
	}
}
