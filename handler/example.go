package handler

import (
	"github.com/nlopes/slack"
)

// PingHandler handles ping request from user
func PingHandler(rtm *slack.RTM, request *slack.MessageEvent) *slack.OutgoingMessage {
	return rtm.NewOutgoingMessage("pong", request.Channel)
}
