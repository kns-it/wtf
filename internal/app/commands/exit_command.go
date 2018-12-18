package commands

import (
	"flag"
	"os"
)

type ExitCommand struct {
	code *int
}

func (ec *ExitCommand) Text() string {
	return "exit"
}

func (ec *ExitCommand) Description() string {
	return "Exit the application"
}

func (ec *ExitCommand) Flags() *flag.FlagSet {
	flags := flag.NewFlagSet("Exit", flag.ContinueOnError)

	ec.code = flags.Int("c", 0, "Optional exit code")

	return flags
}

func (ec *ExitCommand) Call() {
	os.Exit(*ec.code)
}