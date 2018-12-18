package main

import (
	"flag"
	"fmt"
	"github.com/c-bata/go-prompt"
	cmd "github.com/kns-it/wtf/internal/app/builtInCommands"
	"github.com/kns-it/wtf/pkg/commands"
	"os"
	"strings"
)

func main() {
	ctx := commands.NewCommandContext()
	stdoutWriter := os.Stdout

	cmd.InitCommands(ctx, stdoutWriter)
	suggests := ctx.GetCommandSuggests()

	shallExit := false

	for !shallExit {
		t := prompt.Input(getCursor(ctx), func(document prompt.Document) []prompt.Suggest {

			if strings.TrimSpace(document.TextBeforeCursor()) == "" {
				return suggests
			}

			args := strings.Split(strings.TrimSpace(document.TextBeforeCursor()), " ")

			if len(args) >= 1 {
				if currentCmd, ok := ctx.GetCommand(args[0]); ok {
					flags := currentCmd.Flags()
					subSuggests := make([]prompt.Suggest, 0)
					flags.VisitAll(func(flag *flag.Flag) {
						subSuggests = append(subSuggests, prompt.Suggest{
							Text:        fmt.Sprintf("-%s", flag.Name),
							Description: flag.Usage,
						})
					})
					return prompt.FilterHasPrefix(subSuggests, document.GetWordBeforeCursor(), true)
				}
			}
			return prompt.FilterHasPrefix(suggests, document.GetWordBeforeCursor(), true)
		})

		if strings.TrimSpace(t) == "" {
			continue
		}

		args := strings.Split(t, " ")
		if cmdToCall, ok := ctx.GetCommand(args[0]); ok {
			if err := cmdToCall.Flags().Parse(args[1:]); err == nil {
				cmdToCall.Call()
			}
		}
	}
}

func getCursor(ctx *commands.CommandContext) string {
	host := "No host"
	port := "No port or service"
	command := "No command"

	if ctx.GetHost() != "" {
		host = ctx.GetHost()
	}

	if ctx.GetPort() != 0 {
		port = string(ctx.GetPort())
	}

	if ctx.GetCurrentCommand() != nil {
		command = ctx.GetCurrentCommand().Text()
	}

	return fmt.Sprintf("[ %s ] [ %s ] [ %s ] >>> ", host, port, command)
}
