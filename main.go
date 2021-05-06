package main

import (
	"fmt"
	"os"
	"tisky/covid"
	"tisky/errs"
	"tisky/scrape"
)

func main() {

	/*

		If command doesn't have a subcommand will return this.
		TODO: Instead of replying the NotEnoughArgs() error, reply a list of subcommands.

	*/
	if len(os.Args) < 2 {
		fmt.Println("tisky\ntisky covid -help")
		os.Exit(0)
	}

	/* Handle the subcommands. */
	switch os.Args[1] {
	case "covid":
		covid.Handle()
	case "scrape":
		scrape.Handle()
	default:
		errs.CommandNotFound()
		os.Exit(0)
	}
}
