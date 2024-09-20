package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/CrisGui/glox/errors"
	"github.com/CrisGui/glox/lexer"
	"github.com/CrisGui/glox/utils"
)

func runSourceFile(filename string) {
	fmt.Printf("Running source file: %v\n", filename)
}
func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			errors.SendError(1)
		}
		// fmt.Printf("%s\n", scanner.Text())
		lexer.ParsePrompt(scanner.Text())
	}
}

func printHelpMessage() {
	fmt.Printf("Usage: glox [options]\n")
	fmt.Printf("OPTIONS:\n")
	fmt.Printf("\t-s[ource] <FILE>\n")
	fmt.Printf("\t-p[rompt]\n")
	fmt.Printf("\t-h[elp]\n")
}

func parseArguments(arguments []string) {
	if len(arguments) < 1 {
		errors.SendError(errors.TerminalCategory, 0)
	}
	switch arguments[0] {
	case "-p", "--prompt":
		runPrompt()
	case "-s", "--source":
		utils.TODO("Fix error when filename isn't specified")
		runSourceFile(arguments[1])
	case "-h", "--help":
		printHelpMessage()
	default:
		errors.SendError(errors.TerminalCategory, 0)
	}
}

func Run(arguments []string) {
	if len(arguments) > 1 {
		parseArguments(arguments[1:])
	} else {
		printHelpMessage()
	}
}
