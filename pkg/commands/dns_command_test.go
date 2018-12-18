package commands_test

import (
	"github.com/kns-it/wtf/pkg/commands"
	"strings"
	"testing"
)

type stringWriter struct {
	stringBuilder strings.Builder
}

func newStringWriter() *stringWriter {
	return &stringWriter{
		stringBuilder: strings.Builder{},
	}
}

func (sw *stringWriter) Write(p []byte) (n int, err error) {
	return sw.stringBuilder.Write(p)
}

func (sw *stringWriter) String() string {
	return sw.stringBuilder.String()
}

func TestDNSCommand_Call_A_Record(t *testing.T) {

	ctx := commands.NewCommandContext()
	ctx.SetHost("www.kns-it.de")

	writer := newStringWriter()
	dnsCommand := commands.NewDNSCommand(ctx, writer)
	flags := dnsCommand.Flags()

	if err := flags.Parse([]string{"-m", "fwd"}); err != nil {
		t.FailNow()
	}

	dnsCommand.Call()

	if !strings.Contains(writer.String(), "136.243.7.115") {
		t.Fail()
	}

	if !strings.Contains(writer.String(), "2a01:4f8:211:fc10:499:2:0:1") {
		t.Fail()
	}
}