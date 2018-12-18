package builtInCommands

import (
	"github.com/kns-it/wtf/pkg/commands"
	"io"
)

func InitCommands(ctx *commands.CommandContext, writer io.Writer) {
	ctx.RegisterCommand(NewSetHostCommand(ctx))
	ctx.RegisterCommand(commands.NewDNSCommand(ctx, writer))
	ctx.RegisterCommand(&ExitCommand{})
}
