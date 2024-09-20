package errors

import "fmt"

type LoxError interface{}

type ErrorLexer int
type ErrorTerminal int

// Lexer error enum
const (
	L0000 ErrorLexer = iota
)

// Terminal error enum
const (
	T0000 ErrorTerminal = iota
)

func SendError(error_code any) {
	fmt.Printf("Error %v\n", error_code)
}

func ReportErrorLexer(line int, message string)
