package neth

import "net"

// IsTimeout return true if given err is timeout (err.Timeout() is true).
func IsTimeout(err error) bool {
	if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
		return true
	}
	return false
}
