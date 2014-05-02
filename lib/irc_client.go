package inkfish

import (
	"fmt"
	"github.com/thoj/go-ircevent"
	"log"
)

type IrcClient struct {
	client *irc.Connection
	option *Option
}

func NewIrcClient(option *Option) (ircClient *IrcClient) {
	client := irc.IRC(option.Nickname, option.User)

	if option.Keyword != "" {
		client.Password = option.Keyword
	}

	ircClient = &IrcClient{client: client, option: option}
	return
}

func (self *IrcClient) Run() {
	err := self.client.Connect(fmt.Sprintf("%s:%d", self.option.Host, self.option.Port))
	defer func() {
		log.Println("INFO: Disconnect")
		self.client.Disconnect()
	}()

	if err != nil {
		log.Fatalf("ERROR: %s\n", err.Error())
	}

	go self.handleError()
}

func (self *IrcClient) join(message *Message) {
	self.client.Join("#" + message.Channel)
}

func (self *IrcClient) leave(message *Message) {
	self.client.Part("#" + message.Channel)
}

func (self *IrcClient) notice(message *Message) {
	self.client.Notice("#"+message.Channel, message.Body)
}

func (self *IrcClient) privmsg(message *Message) {
	self.client.Privmsg("#"+message.Channel, message.Body)
}

func (self *IrcClient) handleError() {
	err := <-self.client.Error
	log.Printf("ERROR: %s", err.Error())
}
