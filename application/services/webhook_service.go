package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/juliofilizzola/bot_discord/application/domain"
	"github.com/juliofilizzola/bot_discord/application/domain/model"
	"github.com/juliofilizzola/bot_discord/application/domain/repository"
	"log"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/juliofilizzola/bot_discord/application/port/input"
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

func (web webhookDomainService) Save(dataGit domain.Github) {
	user, err := web.getUserOrCreate(dataGit.Sender)

	if err != nil {
		return
	}

	var reviews []*model.User

	reviewers := dataGit.PullRequest.RequestedReviewers
	userReviewers := make(map[string]*model.User, len(reviewers))

	for _, reviewer := range reviewers {
		userReviewer, err := web.userRepo.GetUserByGithubUsername(reviewer.Login)
		if err == nil {
			//userReviewers[reviewer.Login] = userReviewer
			fmt.Println("userReviewer", userReviewer)
		}
	}

	for _, userReviewer := range userReviewers {
		reviews = append(reviews, userReviewer)
	}

	pr := createPRFromGithubData(&dataGit, user, reviews)

	if err := web.repo.Save(&pr); err != nil {
		fmt.Println(err)
	}
}

func createPRFromGithubData(dataGit *domain.Github, user *model.User, reviewers []*model.User) model.PR {
	return model.PR{
		ID:          uuid.New().String(),
		Base:        model.Base{},
		Url:         dataGit.PullRequest.Url,
		Number:      strconv.Itoa(dataGit.PullRequest.Number),
		State:       dataGit.PullRequest.State,
		HtmlUrl:     dataGit.PullRequest.HtmlUrl,
		Title:       dataGit.PullRequest.Title,
		Description: dataGit.PullRequest.Body,
		CreatedAtPr: dataGit.PullRequest.CreatedAt,
		ClosedAt:    dataGit.PullRequest.ClosedAt,
		Color:       GetColorByString(dataGit.PullRequest.State),
		//OwnerPR:         user,
		OwnerID: strconv.Itoa(dataGit.PullRequest.User.Id),
		//Reviewers:       reviewers,
		Locked:          false,
		CommitsUrl:      dataGit.PullRequest.CommitsUrl,
		BranchName:      dataGit.PullRequest.Head.Ref,
		IntroBranchName: dataGit.PullRequest.Base.Ref,
	}
}

func GetColorByString(input string) int {
	switch {
	case strings.Contains(input, "hot") || strings.Contains(input, "hotfix"):
		return ColorRed
	case strings.Contains(input, "fix"):
		return ColorOrange
	case strings.Contains(input, "feat") || strings.Contains(input, "feature"):
		return ColorGreen
	default:
		return ColorWhite
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
