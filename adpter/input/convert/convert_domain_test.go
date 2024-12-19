package convert

import (
	"github.com/bwmarrin/discordgo"
	"github.com/juliofilizzola/bot_discord/application/domain"
	"reflect"
	"testing"
)

var githubTestCases = []struct {
	name   string
	input  *domain.Github
	output discordgo.WebhookParams
}{
	{
		name:  "EmptyGithub",
		input: &domain.Github{},
		output: discordgo.WebhookParams{
			Content: "Nova PR no Repositorio: ",
			TTS:     false,
			AllowedMentions: &discordgo.MessageAllowedMentions{
				Parse: []discordgo.AllowedMentionType{
					discordgo.AllowedMentionTypeEveryone,
				},
				RepliedUser: false,
			},
			Flags: 0,
		},
	},
	// other cases
}

func TestDomainGithub(t *testing.T) {
	for _, tc := range githubTestCases {
		t.Run(tc.name, func(t *testing.T) {
			got := DomainGithub(tc.input)
			if !reflect.DeepEqual(got, tc.output) {
				t.Errorf("Expected %v, got %v", tc.output, got)
			}
		})
	}
}
