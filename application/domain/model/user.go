package model

type User struct {
	ID             string `json:"id" valid:"uuid" gorm:"type:uuid;primary_key default:uuid_generate_v4()"`
	Name           string `json:"name" valid:"required"`
	GithubUsername string `json:"github_username"`
	AvatarUrl      string `json:"avatar_url"`
	PRS            []*PR  `json:"prs"`
	Base           `valid:"required"`
}
