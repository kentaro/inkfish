package inkfish

type Message struct {
	Command string
	Body    string
}

func NewMessage(command, body string) (message *Message) {
	return &Message{Command: command, Body: body}
}
