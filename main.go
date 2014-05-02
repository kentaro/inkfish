package main

import (
	flags "github.com/jessevdk/go-flags"
	"github.com/kentaro/inkfish/lib"
	"os"
)

var opts inkfish.Option

func main() {
	parser := flags.NewParser(&opts, flags.Default)
	_, err := parser.Parse()

	if err != nil {
		os.Exit(1)
	}

	inkfish := &inkfish.Core{}
	inkfish.Run(opts)
}
