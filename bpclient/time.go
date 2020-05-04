package bpclient

import (
	"time"
)

// CurrentDate returns the current date in the format expected by the API
func CurrentDate() string {
	return time.Now().Format("2006-01-02T15:04:05.999Z")
}
