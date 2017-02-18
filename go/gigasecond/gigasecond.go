// Package gigaseconds calculates the datetime after a given datetime in which 1e9 seconds have passed
package gigasecond

import "time"

const testVersion = 4

type time struct {
	Time time
}

// API function.  It uses a type from the Go standard library.
func AddGigasecond(time.Time) time.Time {

}