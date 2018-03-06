package featuremill

import (
	"fmt"
	"strings"

	"github.com/spaolacci/murmur3"
)

// ExtractText returns a slice of "featureID:1" strings for each token in the string
func ExtractText(text, delim string) []string {
	features := []string{}
	texts := strings.Split(text, delim)
	for _, v := range texts {
		// feature id per word
		// this works okay for sparse technical logs
		// otherwise you might want to use an IDF transformation
		// to under-weight less meaningfull words/tokens
		fID := int32(murmur3.Sum32([]byte(v)))
		features = append(features, fmt.Sprintf("%d:1", fID))
	}

	return features
}
