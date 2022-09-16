package timestamp

import (
	"time"
)

// From returns a timestamp string
// from an incoming time.Time value.
func From(now time.Time) string {
	return now.Format("2006/01/02 15:04:05.000")
}

// Now returns a timestamp string
// from the current time.
func Now() string {
	return DepGetTime().Format("2006/01/02 15:04:05.000")
}
