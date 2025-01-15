package services

import (
	"github.com/juliofilizzola/bot_discord/application/domain/model"
	"github.com/juliofilizzola/bot_discord/application/domain/repository"
)

type PRService struct {
	repo repository.PRRepository
}

func NewPRService(repo repository.PRRepository) *PRService {
	return &PRService{repo}
}

func (s *PRService) CreatePR(pr *model.PR) error {
	return s.repo.Save(pr)
}

func (s *PRService) GetPRByID(id string) (*model.PR, error) {
	return s.repo.FindByID(id)
}

func (s *PRService) GetAllPRs() ([]*model.PR, error) {
	return s.repo.FindAll(1, 2)
}

func (s *PRService) DeletePR(id string) error {
	return s.repo.Delete(id)
}
