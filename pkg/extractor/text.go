package extractor

import (
	"strings"

	"github.com/spaolacci/murmur3"
)

// TransformText returns word vector indices via hashing vectorizer
func TransformText(text string) []uint32 {
	features := []uint32{}
	texts := strings.Split(text, " ")
	for _, v := range texts {
		i := murmur3.Sum32([]byte(v))
		features = append(features, i)
	}

	return features
}
