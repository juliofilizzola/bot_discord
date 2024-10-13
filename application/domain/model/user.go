package model

type User struct {
	Name           string `json:"name" valid:"required"`
	GithubUsername string `json:"github_username"`
	AvatarUrl      string `json:"avatar_url"`
	PRS            []*PR  `json:"prs"`
	Base           `valid:"required"`
}
