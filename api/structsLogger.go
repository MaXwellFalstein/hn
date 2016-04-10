package hnapi

import "sync"

// Logger is the logger used by HN API when exposing verbose mode.
type Logger struct {
	Verbose bool
	mutex   *sync.Mutex
}
