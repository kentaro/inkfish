package inkfish

type Queue struct {
	message chan *Message
	quit    chan bool
}

func NewQueue() (queue *Queue) {
	message := make(chan *Message)
	quit := make(chan bool, 1)
	queue = &Queue{message: message, quit: quit}
	return
}

func (self *Queue) Run() {
	<- self.quit
}

func (self *Queue) Push(message *Message) {
	self.message <- message
}

func (self *Queue) Pop() (message *Message) {
	message = <-self.message
	return
}

func (self *Queue) Quit() {
	self.quit <- true
}
