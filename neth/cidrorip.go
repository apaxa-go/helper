package neth

/*
import (
	"net"
	"strings"
)

func ParseCIDROrIP(s string) (ip net.IP, ipnet *net.IPNet, err error) {
	if strings.Contains(s, "/") < 0 {
		ip  = net.ParseIP(s)
		if ip==nil {
			err = net.ParseError{"IP address", s}
		}
		return
	}else{
		return net.ParseCIDR(s)
	}
}
*/
