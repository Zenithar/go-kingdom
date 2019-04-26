package helpers

import "time"

// TimeFunc is clock indirection function for testing purpose
var TimeFunc = time.Now().UTC
