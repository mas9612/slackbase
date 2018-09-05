package slackbase

import (
	"github.com/mas9612/slackbase/message"
	"github.com/nlopes/slack"
)

// Client holds Slack client and some data which is used to send message to Slack bot
type Client struct {
	Client   *slack.Client
	commands map[string]MessageHandler
}

// MessageHandler represents handler function type
type MessageHandler func(rtm *slack.RTM, request *slack.MessageEvent) message.Message

// NewClient returns the pointer to Client struct
func NewClient(apiToken string) *Client {
	client := slack.New(apiToken)
	return &Client{
		Client:   client,
		commands: map[string]MessageHandler{},
	}
}

// Run start new Slack bot
func (c *Client) Run() {
	rtm := c.Client.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case event := <-rtm.IncomingEvents:
			switch msg := event.Data.(type) {
			case *slack.MessageEvent:
				c.handleMessage(rtm, msg)
			}
		}
	}
}

// AddCommand adds new command and handler function to Slack bot
func (c *Client) AddCommand(command string, handler MessageHandler) {
	c.commands[command] = handler
}

func (c *Client) handleMessage(rtm *slack.RTM, request *slack.MessageEvent) {
	msgHandler, ok := c.commands[request.Text]
	if !ok {
		return
	}
	msg := msgHandler(rtm, request)
	sendMsg := rtm.NewOutgoingMessage(msg.Text, msg.Channel)
	rtm.SendMessage(sendMsg)
}
