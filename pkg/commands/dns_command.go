package commands

import (
	"flag"
	"fmt"
	"io"
	"net"
)

type resolver func(host string) []string

var dnsResolvers = map[string]resolver{
	"fwd": func(host string) []string {
		if addrs, err := net.LookupHost(host); err == nil {
			return addrs
		}
		return []string{}
	},
	"rev": func(host string) []string {
		if addrs, err := net.LookupAddr(host); err == nil {
			return addrs
		}
		return []string{}
	},
	"cname": func(host string) []string {
		if cname, err := net.LookupCNAME(host); err == nil {
			return []string{cname}
		}
		return []string{}
	},
	"mx": func(host string) []string {
		if mx, err := net.LookupMX(host); err == nil {
			mxEntryStrings := make([]string, len(mx))
			for i, mxEntry := range mx {
				mxEntryStrings[i] = fmt.Sprintf("Host: %s, Weight: %d", mxEntry.Host, mxEntry.Pref)
			}
			return mxEntryStrings
		}
		return []string{}
	},
}

type DNSCommand struct {
	context *CommandContext
	writer  io.Writer
	host    *string

	mode *string
}

func NewDNSCommand(ctx *CommandContext, writer io.Writer) Command {
	return &DNSCommand{
		context: ctx,
		writer:  writer,
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
	dnsCom.mode = flagSet.String("m", "fwd", "What to resolve, possible values [fwd|rev|cname|mx|txt|ns]")

	return flagSet
}

func (dnsCom *DNSCommand) Call() {
	runLookup := func(host string) {
		if resolver, ok := dnsResolvers[*dnsCom.mode]; ok {
			_, _ = fmt.Fprintln(dnsCom.writer, "Resolved entries:")
			for _, addr := range resolver(host) {
				_, _ = fmt.Fprintln(dnsCom.writer, addr)
			}
		} else {
			_, _ = fmt.Fprintf(dnsCom.writer, "Given mode %s does not match any resolver\n", *dnsCom.mode)
		}
	}

	if dnsCom.host != nil && *dnsCom.host != "" {
		runLookup(*dnsCom.host)
	} else if dnsCom.context.GetHost() != "" {
		runLookup(dnsCom.context.GetHost())
	}
}
