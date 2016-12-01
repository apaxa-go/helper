package timeh

import "time"

// SimpleLayout is additional predefined layout for use in Time.Format and Time.Parse.
// The reference time used in the layouts is the specific time: "Mon Jan 2 15:04:05 MST 2006".
const SimpleLayout = "2006-01-02 15:04:05"

// Constants defines correlation between different time units.
const (
	// Base

	MonthsInYear = 12
	HoursInDay   = 24
	MinsInHour   = 60
	SecsInMin    = 60

	// Derived

	MinsInDay  = MinsInHour * HoursInDay
	SecsInDay  = SecsInMin * MinsInDay
	SecsInHour = SecsInMin * MinsInHour

	// Seconds fraction

	MillisecsInSec = 1e3
	MicrosecsInSec = 1e6
	NanosecsInSec  = 1e9
	PicosecsInSec  = 1e12

	MicrosecsInMillisec = 1e3
	NanosecsInMillisec  = 1e6
	PicosecsInMillisec  = 1e9

	NanosecsInMicrosec = 1e3
	PicosecsInMicrosec = 1e6

	PicosecsInNanosec = 1e3
)

// Approximate time constants.
const (
	DaysInMonth = 30
)

// UnixEpoch returns local Time corresponding to the beginning of UNIX epoch.
func UnixEpoch() time.Time {
	return time.Unix(0, 0)
}
