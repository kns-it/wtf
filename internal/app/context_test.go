package app_test

import (
	"github.com/kns-it/wtf/internal/app"
	"testing"
)

func TestServiceOrPort_SetService(t *testing.T) {
	sop := app.ServiceOrPort{}
	sop.SetService(app.HTTP)
}

func TestCommandContext_SetServiceOrPortWithoutError(t *testing.T) {

	sop := app.ServiceOrPort{}
	sop.SetService(app.HTTPS)

	ctx := app.CommandContext{}
	err := ctx.SetServiceOrPort(sop)

	if err != nil {
		t.Fail()
		t.Error(err)
	}
}

func TestCommandContext_SetServiceOrPortWithError(t *testing.T) {

	sop := app.ServiceOrPort{}

	ctx := app.CommandContext{}
	err := ctx.SetServiceOrPort(sop)

	if err == nil {
		t.Fail()
	}
}