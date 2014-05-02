package inkfish

import (
	"fmt"
)

type Core struct {}

func (self *Core) Run(opts Option) {
	fmt.Println(opts)
}
