package repository

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/juliofilizzola/bot_discord/application/domain/model"
	"github.com/lib/pq"
)

type PRRepository interface {
	Save(pr *model.PR) error
	FindByID(id string) (*model.PR, error)
	FindAll(limit, offset int) ([]*model.PR, error)
	Delete(id string) error
}

type PrRepository struct {
	db *gorm.DB
}

func NewPRRepository(db *gorm.DB) *PrRepository {
	return &PrRepository{db}
}

func (r *PrRepository) Save(pr *model.PR) error {
	err := r.db.Create(pr).Error
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" { // Unique violation error code
				return fmt.Errorf("duplicate key value violates uniq: %v", pqErr.Detail)
			}
		}
		return err
	}
	return nil
}

func (r *PrRepository) FindByID(id string) (*model.PR, error) {
	var pr model.PR
	err := r.db.First(&pr, "id = ?", id).Error
	return &pr, err
}

func (r *PrRepository) FindAll(limit, offset int) ([]*model.PR, error) {
	var prs []*model.PR
	err := r.db.Limit(limit).Offset(offset).Find(&prs).Error
	return prs, err
}

func (r *PrRepository) Delete(id string) error {
	return r.db.Delete(&model.PR{}, "id = ?", id).Error
}
