package builtInCommands

import (
	"flag"
	"fmt"
	"github.com/kns-it/wtf/pkg/commands"
)

type SetHostCommand struct {
	context *commands.CommandContext
	flags   *flag.FlagSet
}

func NewSetHostCommand(ctx *commands.CommandContext) commands.Command {
	return &SetHostCommand{
		context: ctx,
	}
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

	shc.context.SetHost(args[0])
}
