package extractor

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTransformNumerical(t *testing.T) {

	expected := "359062843:0.839898"

	got := TransformNumerical("error_rate", 88.1, 22, 100.7)

	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("unexpected difference: (-got +want)\n%s", diff)
	}
}
