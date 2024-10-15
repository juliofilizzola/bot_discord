package convert

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/juliofilizzola/bot_discord/application/domain/repository"
	"github.com/juliofilizzola/bot_discord/application/services"
	db2 "github.com/juliofilizzola/bot_discord/db"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/juliofilizzola/bot_discord/application/domain"
	"github.com/juliofilizzola/bot_discord/config/env"
)

func DomainGithub(githubDomain *domain.Github) discordgo.WebhookParams {
	var reviews []string
	var assignees []string

	if len(githubDomain.PullRequest.RequestedReviewers) > 0 {
		for _, value := range githubDomain.PullRequest.RequestedReviewers {
			reviews = append(reviews, value.Login)
		}
	}

	if len(githubDomain.PullRequest.Assignees) > 0 {
		for _, value := range githubDomain.PullRequest.Assignees {
			assignees = append(assignees, value.Login)
		}
	}

	embed := &discordgo.MessageEmbed{
		URL:         githubDomain.PullRequest.HtmlUrl,
		Type:        discordgo.EmbedTypeLink,
		Title:       githubDomain.PullRequest.Title,
		Description: githubDomain.PullRequest.Body,
		Timestamp:   time.Now().Format(`2006-01-02 15:04:05`),
		Color:       16776960,
		Footer: &discordgo.MessageEmbedFooter{
			Text:         githubDomain.Organization.Login,
			IconURL:      githubDomain.Organization.AvatarUrl,
			ProxyIconURL: "",
		},
		Image: &discordgo.MessageEmbedImage{
			URL:      githubDomain.Organization.AvatarUrl,
			ProxyURL: "",
			Width:    280,
			Height:   20,
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL:      githubDomain.Organization.AvatarUrl,
			ProxyURL: "",
			Width:    280,
			Height:   20,
		},
		Author: &discordgo.MessageEmbedAuthor{
			URL:          githubDomain.PullRequest.User.HtmlUrl,
			Name:         githubDomain.PullRequest.User.Login,
			IconURL:      githubDomain.PullRequest.User.AvatarUrl,
			ProxyIconURL: "",
		},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Branch:",
				Value:  githubDomain.PullRequest.Head.Ref,
				Inline: false,
			},
			{
				Name:   "Merge into:",
				Value:  githubDomain.PullRequest.Base.Ref + " from " + githubDomain.PullRequest.Head.Ref,
				Inline: false,
			},
			{
				Name:   "Status:",
				Value:  githubDomain.PullRequest.State,
				Inline: false,
			},
			{
				Name: "Assinado:",
				Value: func() string {
					if len(githubDomain.PullRequest.Assignee.Login) == 0 {
						return "Não teve assinatura"
					}
					return getUserDiscord(assignees)
				}(),
				Inline: false,
			},
			{
				Name:   "Codigo adicionado:",
				Value:  strconv.Itoa(githubDomain.PullRequest.Additions),
				Inline: true,
			},
			{
				Name:   "Codigo deletado",
				Value:  strconv.Itoa(githubDomain.PullRequest.Deletions),
				Inline: true,
			},
			{
				Name:   "Commits:",
				Value:  "[commits](" + githubDomain.PullRequest.HtmlUrl + "/commits)",
				Inline: false,
			},
			{
				Name:   "Reviews",
				Value:  getUserDiscord(reviews),
				Inline: false,
			},
		},
	}

	return discordgo.WebhookParams{
		Content:    "Nova PR no Repositorio: " + githubDomain.Repository.Name,
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

func getUserDiscord(reviews []string) string {
	var formattedReviews []string
	db, err := db2.ConnectDB()
	if err != nil {
		return strings.Join(reviews, ", ")
	}

	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("Error to close connection with database")
		}
	}(db)

	repo := repository.NewUserRepository(db)
	service := services.NewUserService(repo)

	for _, username := range reviews {
		user, err := service.GetUserByGithubUsername(username)
		if err != nil {
			fmt.Println("Error fetching user:", err)
		}
		if user == nil {
			formattedReviews = append(formattedReviews, username)
			continue
		}
		formattedReviews = append(formattedReviews, fmt.Sprintf("<@%s>", user.UserId))
	}

	return strings.Join(formattedReviews, ", ")
}
