package featuremill

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"

	"github.com/spaolacci/murmur3"
)

// TransformIP returns a vector that is a scaled integer representation of IPv4 and IPv6 IPs
// with a deterministic feature ID
func TransformIP(field, addr string) (string, error) {
	ip := net.ParseIP(addr)
	if ip != nil {
		if len(ip) == 16 {
			intRep := binary.BigEndian.Uint32(ip[12:16])
			// deterministic collision resistant feature id
			featureID := murmur3.Sum32([]byte(uniqueHashPrefixStr + field))
			// TODO: need to scale intRep
			res := fmt.Sprintf("%d:%d", featureID, intRep)
			return res, nil
		}
		return "", nil
	}
	return "", errors.New("not a valid IP given")
}
