package naiveping

import (
	"net"
	"time"
)

// Ping - just naive realization, no comments
func Ping(addr string, timeout time.Duration) (time.Duration, error) {
	timeBefore := time.Now()
	_, err := net.DialTimeout("ip:1", addr, timeout)
	timeAfter := time.Now()
	if err != nil {
		return 0, err
	}

	return timeAfter.Sub(timeBefore), nil

}
