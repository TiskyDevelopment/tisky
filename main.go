package main

import (
	"os"
	"tisky/covid"
	"tisky/errs"
)

func main() {

	/*

		If command doesn't have a subcommand will return this.
		TODO: Instead of replying the NotEnoughArgs() error, reply a list of subcommands.

	*/
	if len(os.Args) < 2 {
		errs.NotEnoughArgs()
	}

	/* Handle the subcommands. */
	switch os.Args[1] {
	case "covid":
		covid.Handle()
	default:
		errs.CommandNotFound()
		os.Exit(0)
	}
}
