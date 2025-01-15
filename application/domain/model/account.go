package model

type Account struct {
	Base      `valid:"required"`
	AvatarUrl string
	Url       string
	OwnerName string `json:"owner_name" gorm:"column:owner_name;type:varchar(255);not null" valid:"notnull"`
	GitId     string `json:"git_id" gorm:"column:git_id;type:varchar(255);not null" valid:"notnull"`
}
