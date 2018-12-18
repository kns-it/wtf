package commands

import (
	"github.com/c-bata/go-prompt"
)

type svcOrPortAssignmentError struct {
	msg string
}

func (err svcOrPortAssignmentError) Error() string {
	return err.msg
}

type Service int16
type Port int16

type ServiceOrPort struct {
	port int16
}

func (sop *ServiceOrPort) SetService(svc Service) {
	sop.port = int16(svc)
}

func (sop *ServiceOrPort) SetPort(port Port) {
	sop.port = int16(port)
}

type CommandContext struct {
	commandByName map[string]Command
	host          string
	command       Command
	port          ServiceOrPort
}

func NewCommandContext() *CommandContext {
	return &CommandContext{
		commandByName: make(map[string]Command),
	}
}

func (ctx *CommandContext) RegisterCommand(cmd Command) {
	if cmd != nil {
		ctx.commandByName[cmd.Text()] = cmd
	}
}

func (ctx *CommandContext) GetCommand(cmdName string) (Command, bool) {
	if cmd, ok := ctx.commandByName[cmdName]; ok {
		return cmd, true
	}
	return nil, false
}

func (ctx *CommandContext) GetCommandSuggests() []prompt.Suggest {
	suggests := make([]prompt.Suggest, 0)
	for _, v := range ctx.commandByName {
		suggests = append(suggests, prompt.Suggest{
			Text:        v.Text(),
			Description: v.Description(),
		})
	}
	return suggests
}

func (ctx CommandContext) GetHost() string {
	return ctx.host
}

func (ctx *CommandContext) SetHost(host string) {
	ctx.host = host
}

func (ctx CommandContext) GetPort() int16 {
	return ctx.port.port
}

func (ctx *CommandContext) SetServiceOrPort(port ServiceOrPort) error {
	if port.port != 0 {
		ctx.port = port
		return nil
	}

	return svcOrPortAssignmentError{
		msg: "Port to assign was the default port",
	}
}

func (ctx *CommandContext) SetCommand(command Command) {
	if command != nil {
		ctx.command = command
	}
}

func (ctx *CommandContext) GetCurrentCommand() Command {
	return ctx.command
}
