package commands

import (
	"flag"
	"fmt"
	"github.com/kns-it/wtf/internal/app"
)

type SetHostCommand struct {
	Context *app.CommandContext
	flags *flag.FlagSet
}

func (shc *SetHostCommand) Text() string {
	return "host"
}

func (shc *SetHostCommand) Description() string {
	return "Set host context"
}

func (shc *SetHostCommand) Flags() *flag.FlagSet {
	shc.flags = flag.NewFlagSet("Host", flag.ContinueOnError)
	return shc.flags
}

func (shc *SetHostCommand) Call() {
	args := shc.flags.Args()

	if len(args) != 1 {
		fmt.Println("Argument count does not match")
	}

	shc.Context.SetHost(args[0])
}