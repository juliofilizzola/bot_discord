package model

type PR struct {
	ID              string `json:"id" valid:"uuid" gorm:"type:uuid;primary_key default:uuid_generate_v4()"`
	Base            `valid:"required"`
	Url             string     `json:"url" gorm:"type:varchar(255)"`
	Number          string     `json:"number" gorm:"type:varchar(255);uniqueIndex:number"`
	State           string     `json:"state" gorm:"type:varchar(255)"`
	HtmlUrl         string     `json:"html_url" gorm:"type:varchar(255)"`
	Title           string     `json:"title" gorm:"type:varchar(255)"`
	Description     string     `json:"description" gorm:"type:varchar(255)"`
	CreatedAtPr     string     `json:"created_at_pr" gorm:"type:varchar(255)"`
	ClosedAt        string     `json:"closed_at" gorm:"type:varchar(255)"`
	Color           string     `json:"color" gorm:"type:varchar(255)"`
	OwnerPR         *Account   `gorm:"foreignKey:OwnerId"`
	OwnerID         string     `json:"owner_id"`
	Reviewers       []*Account `gorm:"many2many:reviewers_pr;"`
	Locked          bool       `json:"locked"`
	CommitsUrl      string     `json:"commits_url" gorm:"type:varchar(255)"`
	BranchName      string
	IntroBranchName string
}
