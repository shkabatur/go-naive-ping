package naiveping

import (
	"errors"
	"net"
	"time"
)

// Ping - just naive realization, no comments
func Ping(addr string, timeout time.Duration) (time.Duration, error) {
	timeBefore := time.Now()
	_, err := net.DialTimeout("ip:1", addr, timeout)
	timeAfter := time.Now()
	if err != nil {
		return 0, errors.New("Remote machine is note reachable")
	}

	return timeAfter.Sub(timeBefore), nil

}
