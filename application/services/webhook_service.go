package services

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/juliofilizzola/bot_discord/application/domain"
	"github.com/juliofilizzola/bot_discord/application/domain/model"
	"github.com/juliofilizzola/bot_discord/application/domain/repository"
	"github.com/juliofilizzola/bot_discord/application/port/input"
	"log"
	"strconv"
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
		_, err := web.server.WebhookExecute(webhookId, webhookToken, true, dataGit)
		if err != nil {
			log.Fatal(err)
		}
		return "deu bom"
	}
	return "deu ruim!"
}

func (web webhookDomainService) Save(dataGit model.PR) {
	if err := web.repo.Save(&dataGit); err != nil {
		fmt.Println(err)
	}
}

func (web webhookDomainService) getUserOrCreate(sender domain.User) (*model.User, error) {
	user, err := web.userRepo.GetUserByGithubUsername(sender.Login)
	if err != nil {
		if err := web.createUser(sender); err != nil {
			return nil, err
		}
		return web.userRepo.GetUserByGithubUsername(sender.Login)
	}
	return user, nil
}

func (web webhookDomainService) createUser(sender domain.User) error {
	user := model.User{
		ID:             uuid.New().String(),
		Name:           sender.Login,
		UserId:         strconv.Itoa(sender.Id),
		GithubUsername: sender.Login,
		AvatarUrl:      sender.AvatarUrl,
		PRS:            nil,
		Base:           model.Base{},
	}

	return web.userRepo.CreateUser(&user)
}
