package configs

import (
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/slack-go/slack"
)

type slackConfig struct {
	OAuthToken string `envconfig:"SLACK_OAUTH_TOKEN" default:""`
	ChannelID  string `envconfig:"SLACK_CHANNEL_ID" default:""`
}

// NewSlackConfig for slack notify
func NewSlackConfig(title string, url string) {
	var slackConfig slackConfig

	err := envconfig.Process("", &slackConfig)
	if err != nil {
		log.Fatal("Slack config not found")
	}

	api := slack.New(slackConfig.OAuthToken)
	attachment := slack.Attachment{
		TitleLink: url,
		Title:     title,
	}

	channelID, timestamp, err := api.PostMessage(
		slackConfig.ChannelID,
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
	)
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	log.Printf("Message successfully sent to Channel %s at %s\n", channelID, timestamp)
}
