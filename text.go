package extractor

import (
	"fmt"
	"strings"

	"github.com/spaolacci/murmur3"
)

// TransformText returns word vector indices via hashing vectorizer
func TransformText(text string) []string {
	features := []string{}
	texts := strings.Split(text, " ")
	for _, v := range texts {
		// feature id per word
		// this works okay for sparse technical logs
		// otherwise you might want to use an IDF transformation
		// to under-weight less meaningfull words/tokens
		fID := murmur3.Sum32([]byte(v))
		features = append(features, fmt.Sprintf("%d:1", fID))
	}

	return features
}
