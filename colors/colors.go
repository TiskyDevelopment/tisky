package colors

import "fmt"

/* Colors used by Tisky. */

const (
	Red    = "\u001b[31m"
	Yellow = "\u001b[33m"
	Green  = "\u001b[32m"
	White  = "\u001b[37m"
)

/* Print red text. */
func PrintRed(content string) {
	fmt.Printf("%v%v%v\n", Red, content, White)
}

/* Print yellow text. */
func PrintYellow(content string) {
	fmt.Printf("%v%v%v\n", Yellow, content, White)
}

/* Print green text. */
func PrintGreen(content string) {
	fmt.Printf("%v%v%v\n", Green, content, White)
}
