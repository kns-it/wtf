package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/kns-it/wtf/internal/app"
	cmd "github.com/kns-it/wtf/internal/app/commands"
	"os"
	"strings"
)

func main() {
	ctx := app.CommandContext{}
	stdoutWriter := bufio.NewWriter(os.Stdout)

	cmds := cmd.InitCommands(&ctx, stdoutWriter)
	suggests, cmdMap := cmds.ToSuggests()

	shallExit := false

	for !shallExit {
		t := prompt.Input(getCursor(&ctx), func(document prompt.Document) []prompt.Suggest {

			if strings.TrimSpace(document.TextBeforeCursor()) == "" {
				return suggests
			}

			args := strings.Split(strings.TrimSpace(document.TextBeforeCursor()), " ")

			if len(args) >= 1 {
				if currentCmd, ok := cmdMap[args[0]]; ok {
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
		if cmdToCall, ok := cmdMap[args[0]]; ok {
			if err := cmdToCall.Flags().Parse(args[1:]); err == nil {
				cmdMap[args[0]].Call()
			}
		}
	}
}

func getCursor(ctx *app.CommandContext) string {
	host := "No host"
	port := "No port or service"

	if ctx.GetHost() != "" {
		host = ctx.GetHost()
	}

	if ctx.GetPort() != 0 {
		port = string(ctx.GetPort())
	}

	return fmt.Sprintf("[ %s ] [ %s ] ", host, port)
}
