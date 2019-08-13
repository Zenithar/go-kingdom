package helpers

import "time"

// TimeFunc is clock indirection function for testing purpose
var TimeFunc = func() time.Time {
	return time.Now().UTC()
}
