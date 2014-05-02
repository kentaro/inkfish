package inkfish

type Core struct{}

func (self *Core) Run(opts Option) {
	queue := NewQueue()
	queue.Run()
}
