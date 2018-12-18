package commands

import (
	"github.com/c-bata/go-prompt"
	"github.com/kns-it/wtf/internal/app"
	cmds "github.com/kns-it/wtf/pkg/commands"
	"io"
)

type Commands struct {
	commands []cmds.Command
}

func InitCommands(ctx *app.CommandContext, writer io.Writer) Commands {
	return Commands{
		commands: []cmds.Command{
			&SetHostCommand{
				Context: ctx,
			},
			&ExitCommand{},
			cmds.NewDNSCommand(ctx, writer),
		},
	}
}

func (commands *Commands) ToSuggests() ([]prompt.Suggest, map[string]cmds.Command) {
	suggests := make([]prompt.Suggest, 0)
	suggestMap := make(map[string]cmds.Command)
	for _, com := range commands.commands {
		suggests = append(suggests, prompt.Suggest{
			Text:        com.Text(),
			Description: com.Description(),
		})
		suggestMap[com.Text()] = com
	}

	return suggests, suggestMap
}
