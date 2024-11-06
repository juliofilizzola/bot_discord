package model

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Base struct {
	CreatedAt time.Time  `json:"created_at" valid:"-"`
	UpdatedAt time.Time  `json:"updated_at" valid:"-"`
	DeletedAt *time.Time `json:"deleted-at" valid:"-" default:"null"`
}
