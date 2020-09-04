package naiveping

import (
	"net"
	"time"
)

// AddrPing using to return address and ping value as result of Ping function
type AddrPing struct {
	Addr string
	Ping time.Duration
}

// Ping - just naive realization, no comments
func Ping(addr string, timeout time.Duration) (AddrPing, error) {
	timeBefore := time.Now()
	_, err := net.DialTimeout("ip:1", addr, timeout)
	timeAfter := time.Now()
	if err != nil {
		return AddrPing{"", 0}, err
	}

	return AddrPing{Addr: addr, Ping: timeAfter.Sub(timeBefore)}, nil

}
