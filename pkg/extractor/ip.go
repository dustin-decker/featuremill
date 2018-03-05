package extractor

import (
	"encoding/binary"
	"errors"
	"net"
)

// TransformIP returns integer representation of IPv4 and IPv6 IPs
func TransformIP(addr string) (uint32, error) {
	ip := net.ParseIP(addr)
	if ip != nil {
		if len(ip) == 16 {
			return binary.BigEndian.Uint32(ip[12:16]), nil
		}
		return binary.BigEndian.Uint32(ip), nil
	}
	return 0, errors.New("not a valid IP given")
}
