package handler

import (
	"github.com/mas9612/slackbase/message"
	"github.com/nlopes/slack"
)

// PingHandler handles ping request from user
func PingHandler(rtm *slack.RTM, request *slack.MessageEvent) message.Message {
	return message.Message{
		Text:    "pong",
		Channel: request.Channel,
	}
}
