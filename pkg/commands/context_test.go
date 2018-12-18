package commands_test

import (
	"github.com/kns-it/wtf/internal/app"
	"github.com/kns-it/wtf/pkg/commands"
	"testing"
)

func TestServiceOrPort_SetService(t *testing.T) {
	sop := commands.ServiceOrPort{}
	sop.SetService(app.HTTP)
}

func TestCommandContext_SetServiceOrPortWithoutError(t *testing.T) {

	sop := commands.ServiceOrPort{}
	sop.SetService(app.HTTPS)

	ctx := commands.CommandContext{}
	err := ctx.SetServiceOrPort(sop)

	if err != nil {
		t.Fail()
		t.Error(err)
	}
}

func TestCommandContext_SetServiceOrPortWithError(t *testing.T) {

	sop := commands.ServiceOrPort{}

	ctx := commands.CommandContext{}
	err := ctx.SetServiceOrPort(sop)

	if err == nil {
		t.Fail()
	}
}
