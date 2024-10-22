package services

import (
	"fmt"
	"github.com/juliofilizzola/bot_discord/application/domain"
	"github.com/juliofilizzola/bot_discord/application/domain/model"
	"github.com/juliofilizzola/bot_discord/application/domain/repository"
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/juliofilizzola/bot_discord/application/port/input"
)

func NewWebhookDomainService(discord *discordgo.Session, repoPr *repository.PrRepository, repoUser *repository.UserRepo) input.WebhookDomainService {
	return &webhookDomainService{
		server:   discord,
		repo:     repoPr,
		userRepo: repoUser,
	}
}

type webhookDomainService struct {
	server   *discordgo.Session
	repo     *repository.PrRepository
	userRepo *repository.UserRepo
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

func (web webhookDomainService) Save(dataGit *domain.Github) {
	user, err := web.userRepo.GetUserByGithubUsername(strconv.Itoa(dataGit.PullRequest.User.Id))

	if err != nil {
		err := web.userRepo.CreateUser(&model.User{
			ID:             "",
			Name:           "",
			UserId:         "",
			GithubUsername: "",
			AvatarUrl:      "",
			PRS:            nil,
			Base:           model.Base{},
		})
		if err != nil {
			return
		}
	}

	data := model.PR{
		ID:              strconv.Itoa(dataGit.PullRequest.Id),
		Base:            model.Base{},
		Url:             dataGit.PullRequest.Url,
		Number:          strconv.Itoa(dataGit.PullRequest.Number),
		State:           dataGit.PullRequest.State,
		HtmlUrl:         dataGit.PullRequest.HtmlUrl,
		Title:           dataGit.PullRequest.Title,
		Description:     dataGit.PullRequest.Body,
		CreatedAtPr:     dataGit.PullRequest.CreatedAt,
		ClosedAt:        dataGit.PullRequest.ClosedAt,
		Color:           "",
		OwnerPR:         user,
		OwnerID:         strconv.Itoa(dataGit.PullRequest.User.Id),
		Reviewers:       nil,
		Locked:          false,
		CommitsUrl:      dataGit.PullRequest.CommitsUrl,
		BranchName:      dataGit.PullRequest.Head.Ref,
		IntroBranchName: dataGit.PullRequest.Base.Ref,
	}

	err = web.repo.Save(&data)
	if err != nil {
		return
	}
}
