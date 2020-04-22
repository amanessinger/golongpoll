package golongpoll

import (
	"strconv"
	"time"
)

// microsecondStringToTime converts a string representation of milliseconds
// into the corresponding Time value in UTC.  This is used to convert
// timestamps sent by browsers that use javascript's Date.getTime() and then
// pass that number as a url param (hence why it winds up a string).
// IF an invalid string input is used, then time will default to Time{} and the
// error return value will be non-nil.
func microsecondStringToTime(ms string) (time.Time, error) {
	msInt, err := strconv.ParseInt(ms, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(0, msInt*int64(time.Microsecond)).In(time.UTC), nil
}

// timeToEpochMicroseconds converts a Time type to the corresponding
// number of milliseconds since epoch (Jan 1 1970) in UTC.  Note if the Time
// input is before 1970, the corresponding milliseconds value is negative.
func timeToEpochMicroseconds(t time.Time) int64 {
	return t.UnixNano() / int64(time.Microsecond)
}
