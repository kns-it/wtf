package commands

import (
	"flag"
)

type Command interface {
	Text() string
	Description() string
	Flags() *flag.FlagSet
	Call()
}
