package errs

import (
	"fmt"
	"os"
	"tisky/colors"
)

/* Error for less than needed arguments. */
func NotEnoughArgs() {
	fmt.Println(colors.Red + "[ ! ] Not enough arguments." + colors.White)
	os.Exit(0)
}

/* Error for command not found. */
func CommandNotFound() {
	fmt.Println(colors.Red + "[ ! ] Command not found." + colors.White)
	os.Exit(0)
}

/* Error for failed requests. */
func NoInternet() {
	fmt.Println(colors.Red + "[ ! ] Could not get the COVID data. Make sure you have a connection to the internet." + colors.White)
	os.Exit(0)
}

/* Error for bad usage. */
func BadUsage(example string) {
	fmt.Printf("%vBad usage of the command.%v\nExample:\n%v \n", colors.Red, colors.White, example)
	os.Exit(0)
}

/* Print red colored text. */
func PrRed(content string) {
	fmt.Printf("%v%v%v \n", colors.Red, content, colors.White)
	os.Exit(0)
}
