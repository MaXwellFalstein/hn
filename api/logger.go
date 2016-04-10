package hnapi

import (
	"fmt"
	"sync"
)

// NewLogger returns a new Hacker News API Logger struct.
func NewLogger(verbose bool) *Logger {
	return &Logger{
		Verbose: verbose,
		mutex:   new(sync.Mutex),
	}
}

// VerbosePrintln allows you to use println while verifying whether you are or
// are not in verbose mode.
func (l *Logger) VerbosePrintln(a ...interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.Verbose {
		fmt.Println(a...)
	}
}

// VerbosePrintfln allows you to use printfln while verifying whether you are or
// are not in verbose mode.
func (l *Logger) VerbosePrintfln(format string, a ...interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.Verbose {
		fmt.Printf(format+"\n", a...)
	}
}

// VerbosePrintf allows you to use printf while verifying whether you are or
// are not in verbose mode.
func (l *Logger) VerbosePrintf(format string, a ...interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.Verbose {
		fmt.Printf(format, a...)
	}
}
