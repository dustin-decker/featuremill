package extractor

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestText(t *testing.T) {

	text := "testing some shit"

	expected := []uint32{
		uint32(164984899),
		uint32(650894796),
		uint32(232538530),
	}

	got := TransformText(text)

	cmp.Equal(expected, got)
}
