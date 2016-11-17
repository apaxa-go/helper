package httph

import "time"

// CookieDeleteTime returns Time near beginning of UNIX epoch.
// This time should be used as cookie expiration time to remove cookie.
func CookieDeleteTime() time.Time {
	return time.Unix(1, 0)
}
