package featuremill

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTransformCategorical(t *testing.T) {

	expected := "2633207889:1"

	got := TransformCategorical("username or something", "sodunn")

	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("unexpected difference: (-got +want)\n%s", diff)
	}
}
