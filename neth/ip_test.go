package neth

import (
	"net"
	"testing"
)

var tests = []struct {
	ip     net.IP
	isV4   bool
	isV6   bool
	length int
}{
	{isV4: false, isV6: false, length: 0},
	{ip: []byte{1, 2}, isV4: false, isV6: false, length: 0},
	{ip: []byte{1, 2, 3, 4}, isV4: true, isV6: false, length: 4},
	{ip: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff, 1, 2, 3}, isV4: false, isV6: false, length: 0},
	{ip: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff, 1, 2, 3, 4}, isV4: true, isV6: false, length: 4},
	{ip: []byte{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0xff, 0xff, 1, 2, 3, 4}, isV4: false, isV6: true, length: 16},
	{ip: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff, 1, 2, 3, 4, 5}, isV4: false, isV6: false, length: 0},
}

func TestIsIPv4(t *testing.T) {
	for _,test:=range tests{
		if r:=IsIPv4(test.ip); r!=test.isV4{
			t.Errorf("TestIsIPv4: for %v result is %v, but expected %v",test.ip, r, test.isV4)
		}
	}
}

func TestIsIPv6(t *testing.T) {
	for _,test:=range tests{
		if r:=IsIPv6(test.ip); r!=test.isV6{
			t.Errorf("TestIsIPv6: for %v result is %v, but expected %v",test.ip, r, test.isV6)
		}
	}
}

func TestIPLen(t *testing.T) {
	for _,test:=range tests{
		if r:=IPLen(test.ip); r!=test.length{
			t.Errorf("TestIPLen: for %v result is %v, but expected %v",test.ip, r, test.length)
		}
	}
}