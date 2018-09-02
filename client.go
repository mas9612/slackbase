package slackbase

import (
	"github.com/mas9612/slackbase/handler"
	"github.com/nlopes/slack"
)

// Run start new Slack bot
func Run(apiToken string) {
	client := slack.New(apiToken)
	rtm := client.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case event := <-rtm.IncomingEvents:
			switch msg := event.Data.(type) {
			case *slack.MessageEvent:
				handleMessage(rtm, msg)
			}
		}
	}
}

func handleMessage(rtm *slack.RTM, request *slack.MessageEvent) {
	var msg *slack.OutgoingMessage

	if request.Text == "ping" {
		msg = handler.PingHandler(rtm, request)
	}

	rtm.SendMessage(msg)
}
