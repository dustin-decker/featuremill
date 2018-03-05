package extractor

import (
	"errors"
	"fmt"

	"github.com/spaolacci/murmur3"
)

// TransformBoolean returns integer representation of IPv4 and IPv6 IPs
func TransformBoolean(field, boolean string) (string, error) {

	// deterministic collision resistant feature id
	featureID := murmur3.Sum32([]byte(uniqueHashPrefixStr + field))

	if boolean == "false" ||
		boolean == "False" ||
		boolean == "0" ||
		boolean == "no" ||
		boolean == "No" {
		return fmt.Sprintf("%d:%d", featureID, 0), nil
	} else if boolean == "true" ||
		boolean == "True" ||
		boolean == "1" ||
		boolean == "yes" ||
		boolean == "Yes" {
		return fmt.Sprintf("%d:%d", featureID, 1), nil
	}
	return "", errors.New("could not match bool")
}
