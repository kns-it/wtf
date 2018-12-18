package commands

import (
	"flag"
	"fmt"
	"github.com/kns-it/wtf/internal/app"
	"io"
	"log"
	"net"
)

type DNSCommand struct {
	context     *app.CommandContext
	writer      io.Writer
	resolveA    *bool
	resolveAAAA *bool
	host        *string
}

func NewDNSCommand(ctx *app.CommandContext, writer io.Writer) Command {
	return &DNSCommand{
		context: ctx,
		writer: writer,
	}
}

func (dnsCom *DNSCommand) Text() string {
	return "nslookup"
}

func (dnsCom *DNSCommand) Description() string {
	return "Run nslookup"
}

func (dnsCom *DNSCommand) Flags() *flag.FlagSet {
	flagSet := flag.NewFlagSet("DNS", flag.ContinueOnError)
	dnsCom.host = flagSet.String("h", "", "Host override")

	return flagSet
}

func (dnsCom *DNSCommand) Call() {

	runLookup := func(host string) {
		if addresses, err := net.LookupHost(host); err == nil {
			fmt.Println("Resolved addresses:")
			for _, addr := range addresses {
				fmt.Println(addr)
			}
		} else {
			if _, err := fmt.Fprintf(dnsCom.writer, "Got error while resolving %s: %s", host, err.Error()); err != nil {
				log.Println(err)
			}
		}
	}

	if dnsCom.host != nil && *dnsCom.host != "" {
		runLookup(*dnsCom.host)
	} else if dnsCom.context.GetHost() != "" {
		runLookup(dnsCom.context.GetHost())
	}
}
