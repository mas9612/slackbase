package message

// Message represents message sent between user and slackbot
type Message struct {
	Text    string
	Channel string
}
