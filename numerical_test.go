package featuremill

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExtractNumerical(t *testing.T) {

	expected := "359062843:0.839898"

	got := ExtractNumerical("error_rate", 88.1, 22, 100.7)

	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("unexpected difference: (-got +want)\n%s", diff)
	}
}
