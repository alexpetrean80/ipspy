package lib

import (
	"fmt"
	"net"
)

func ParseIP(ipStr string) (net.IP, error) {
	ip := net.ParseIP(ipStr)

	if err := ValidateIP(ip); err != nil {
		return nil, err
	}

	return ip, nil
}

func ValidateIP(ip net.IP) (err error) {
	getErr := func(msg string) error {
		return fmt.Errorf("IP address %s %s.", ip, msg)
	}

	switch {
	case ip == nil:
		err = getErr("is invalid")
	case ip.IsPrivate():
		err = getErr("is private")
	case ip.IsLoopback():
		err = getErr("is loopback")
	case ip.IsUnspecified():
		err = getErr("is unspecified")
	}

	return
}
