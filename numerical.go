package extractor

import (
	"fmt"

	"github.com/spaolacci/murmur3"
)

// TransformNumerical returns a min/max scalled vector to a deterministic feature ID
func TransformNumerical(field string, num float32, min float32, max float32) string {
	featureID := murmur3.Sum32([]byte(uniqueHashPrefixStr + field))
	return fmt.Sprintf("%d:%f", featureID, (num-min)/(max-min))
}
