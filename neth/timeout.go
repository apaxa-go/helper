package neth

import "net"

// IsTimeout return true if given err is timeout (opErr.Timeout).
func IsTimeout(err error) bool {
	if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
		return true
	}
	return false
}
