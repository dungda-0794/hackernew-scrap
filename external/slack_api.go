package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hackernew-scrap/infrastructure"
	"net/http"
	"net/url"

	"github.com/kelseyhightower/envconfig"
)

type slackAPIConfig struct {
	URL       string `required:"true" envconfig:"SLACK_API_URL"`
	ChannelID string `required:"true" envconfig:"SLACK_CHANNEL_ID"`
}

// PostMessage for send notif to slack
func PostMessage(title string, link string) {
	var slackAPIConfig slackAPIConfig

	err := envconfig.Process("SLACK", &slackAPIConfig)
	if err != nil || slackAPIConfig.URL == "" {
		infrastructure.ErrLog.Fatal("Slack config not found")
	}
	httpposturl := slackAPIConfig.URL
	u, _ := url.ParseRequestURI(httpposturl)
	blocks := []interface{}{
		map[string]interface{}{
			"type": "header",
			"text": map[string]string{
				"type": "plain_text",
				"text": fmt.Sprintf(":newspaper: %s", title),
			},
		},
		map[string]interface{}{
			"type": "section",
			"text": map[string]string{
				"type": "mrkdwn",
				"text": link,
			},
		},
	}
	postBody, err := json.Marshal(map[string]interface{}{
		"service":   "slack",
		"channel":   slackAPIConfig.ChannelID,
		"receivers": "here",
		"message": ":alphabet-white-h:" +
			":alphabet-white-a:" +
			":alphabet-white-c:" +
			":alphabet-white-k:" +
			":alphabet-white-e:" +
			":alphabet-white-r:" +
			":alphabet-white-n:" +
			":alphabet-white-e:" +
			":alphabet-white-w:" +
			":alphabet-white-s:",
		"blocks": blocks,
	})
	if err != nil {
		infrastructure.ErrLog.Fatal("fail to generate body", err)
	}

	responseBody := bytes.NewBuffer(postBody)
	res, err := http.Post(u.String(), "application/json", responseBody)
	if err != nil {
		infrastructure.ErrLog.Fatal("fail to post message", err)
	}
	err = res.Body.Close()

	if err != nil {
		infrastructure.ErrLog.Fatal("fail to post message", err)
	}
}
