package extractor

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTransformBoolean(t *testing.T) {

	text := "no"
	expected := "3059673168:0"
	got, _ := TransformBoolean("IsBad", text)
	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("unexpected difference: (-got +want)\n%s", diff)
	}

	text = "true"
	expected = "3059673168:1"
	got, _ = TransformBoolean("IsBad", text)
	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("unexpected difference: (-got +want)\n%s", diff)
	}
}
