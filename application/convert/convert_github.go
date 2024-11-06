package convert

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/juliofilizzola/bot_discord/application/domain"
	"github.com/juliofilizzola/bot_discord/application/domain/model"
	"github.com/juliofilizzola/bot_discord/application/services"
	"github.com/juliofilizzola/bot_discord/config/env"
	"strconv"
	"strings"
	"time"
)

const (
	ColorRed    = 15548997
	ColorGreen  = 5763719
	ColorBlue   = 3447003
	ColorYellow = 16776960
	ColorPurple = 10181046
	ColorOrange = 15105570
	ColorWhite  = 16777215
	ColorBlack  = 0
)

const (
	hotfix  = "hotfix"
	fix     = "fix"
	feat    = "feat"
	feature = "feature"
	hot     = "hot"
)

var userService *services.UserService

func Init(us *services.UserService) {
	userService = us
}

func GithubToDiscord(data *domain.Github) discordgo.WebhookParams {
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
		Color:       getColorByString(data.PullRequest.Title),
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
					return formaterUserDiscord(getUserDiscord(data.PullRequest.Assignee.Login))
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

func GithubToDataBase(data *domain.Github) model.PR {
	var reviews []string

	for _, review := range data.PullRequest.RequestedReviewers {
		reviews = append(reviews, review.Login)
	}

	return model.PR{
		ID:              uuid.New().String(),
		Base:            model.Base{},
		Url:             data.PullRequest.Url,
		Number:          strconv.Itoa(data.PullRequest.Number),
		State:           data.PullRequest.State,
		HtmlUrl:         data.PullRequest.HtmlUrl,
		Title:           data.PullRequest.Title,
		Description:     data.PullRequest.Body,
		CreatedAtPr:     data.PullRequest.CreatedAt,
		ClosedAt:        data.PullRequest.ClosedAt,
		Color:           getColorByString(data.PullRequest.Title),
		OwnerPR:         getUserDiscord(data.PullRequest.User.Login),
		OwnerID:         strconv.Itoa(data.PullRequest.User.Id),
		Reviewers:       getListUserDiscord(reviews),
		Locked:          false,
		CommitsUrl:      data.PullRequest.CommitsUrl,
		BranchName:      data.PullRequest.Head.Ref,
		IntroBranchName: data.PullRequest.Base.Ref,
	}
}

func getUserDiscord(userLogin string) *model.User {
	user, err := userService.GetUserByGithubUsername(userLogin)
	if err != nil {
		return &model.User{
			Name:           userLogin,
			GithubUsername: userLogin,
		}
	}

	return user
}

func getListUserDiscord(usersLogin []string) []*model.User {
	var list []*model.User
	for _, user := range usersLogin {
		list = append(list, getUserDiscord(user))
	}
	return list
}
func formaterUserDiscord(user *model.User) string {
	if user.UserId == "" {
		return user.Name
	}
	return fmt.Sprintf("<@%s>", user.UserId)
}
func getColorByString(state string) int {
	switch {
	case strings.Contains(state, hot) || strings.Contains(state, hotfix):
		return ColorRed
	case strings.Contains(state, fix):
		return ColorOrange
	case strings.Contains(state, feat) || strings.Contains(state, feature):
		return ColorGreen
	default:
		return ColorWhite
	}
}
