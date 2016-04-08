package hnapi

import (
	"fmt"
)

// VerbosePrintln allows you to use println while verifying whether you are or
// are not in verbose mode.
func (l *Logger) VerbosePrintln(a ...interface{}) {
	if l.Verbose {
		fmt.Println(a)
	}
}

// VerbosePrintfln allows you to use printfln while verifying whether you are or
// are not in verbose mode.
func (l *Logger) VerbosePrintfln(format string, a ...interface{}) {
	if l.Verbose {
		fmt.Printf(format+"\n", a)
	}
}

// VerbosePrintf allows you to use printf while verifying whether you are or
// are not in verbose mode.
func (l *Logger) VerbosePrintf(format string, a ...interface{}) {
	if l.Verbose {
		fmt.Printf(format, a)
	}
}
