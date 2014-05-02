package inkfish

type Message struct {
	Command string
	Channel string
	Body    string
}

func NewMessage(command, channel, body string) (message *Message) {
	message = &Message{
		Command: command,
		Channel: channel,
		Body: body,
	}
	return
}
