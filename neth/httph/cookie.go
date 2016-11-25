package httph

import "time"

// CookieDeleteUnixTime is UNIX time near beginning of UNIX epoch.
// It is used as cookie expiration time to remove cookie.
const CookieDeleteUnixTime = 1

// CookieDeleteTime returns Time near beginning of UNIX epoch.
// This time should be used as cookie expiration time to remove cookie.
func CookieDeleteTime() time.Time {
	return time.Unix(CookieDeleteUnixTime, 0)
}
