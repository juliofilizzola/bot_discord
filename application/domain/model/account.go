package model

type Account struct {
	ID        string `json:"id" valid:"uuid" gorm:"type:uuid;primary_key default:uuid_generate_v4()"`
	Base      `valid:"required"`
	AvatarUrl string
	Url       string
	OwnerName string `json:"owner_name" gorm:"column:owner_name;type:varchar(255);not null" valid:"notnull"`
	GitId     string `json:"git_id" gorm:"column:git_id;type:varchar(255);not null" valid:"notnull"`
}
