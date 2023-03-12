package main

import (
	"fmt"
)

// LogLevel is a logging level
type LogLevel uint8

// Possible log levels
const (
	DebugLevel   LogLevel = iota + 1 // iota starts at 0 so debug level is 1
	WarningLevel                     // warning level is 2
	ErrorLevel                       // error level is 3
)

// String implements the fmt.Stringer interface
func (l LogLevel) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case WarningLevel:
		return "warning"
	case ErrorLevel:
		return "error"
	}

	return fmt.Sprintf("unknown log level: %d", l)
}

func main() {
	fmt.Println(WarningLevel) // warning

	lvl := LogLevel(19)
	fmt.Println(lvl) // unknown log level: 19
}
