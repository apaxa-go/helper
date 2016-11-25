package neth

import "net"

// IsIPv4 return true if given ip is IPv4 address (even if it stored as IPv6).
func IsIPv4(ip net.IP) bool {
	return (len(ip) == net.IPv4len) || (len(ip) == net.IPv6len &&
		ip[0] == 0x00 &&
		ip[1] == 0x00 &&
		ip[2] == 0x00 &&
		ip[3] == 0x00 &&
		ip[4] == 0x00 &&
		ip[5] == 0x00 &&
		ip[6] == 0x00 &&
		ip[7] == 0x00 &&
		ip[8] == 0x00 &&
		ip[9] == 0x00 &&
		ip[10] == 0xff &&
		ip[11] == 0xff)
}

// IsIPv6 return true if given ip is IPv6 address (and it is not IPv4 address stored as v6).
func IsIPv6(ip net.IP) bool {
	return len(ip) == net.IPv6len && !IsIPv4(ip)
}

// IPLen return length of given ip in bytes (4 for IPv4 even if it stored as IPv6).
// IPLen return 0 if given address is not valid IP v4 nor v6.
func IPLen(ip net.IP) int {
	switch {
	case IsIPv4(ip):
		return net.IPv4len
	case IsIPv6(ip):
		return net.IPv6len
	default:
		return 0
	}
}
