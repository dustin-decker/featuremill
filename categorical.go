package featuremill

import (
	"fmt"

	"github.com/spaolacci/murmur3"
)

// ExtractCategorical returns a vector that is a positive boolean for a
// deterministic categorical feature ID
func ExtractCategorical(field, category string) string {
	featureID := int32(murmur3.Sum32([]byte(uniqueHashPrefixStr + field + category)))
	return fmt.Sprintf("%d:%d", featureID, 1)
}
