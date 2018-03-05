package extractor

import (
	"fmt"

	"github.com/spaolacci/murmur3"
)

// TransformCategorical returns a positive boolean for a
// deterministic categorical feature ID
func TransformCategorical(field, category string) string {
	featureID := murmur3.Sum32([]byte(uniqueHashPrefixStr + field + category))
	return fmt.Sprintf("%d:%d", featureID, 1)
}
