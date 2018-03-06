package featuremill

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExtractBoolean(t *testing.T) {

	text := "no"
	expected := "-1235294128:0"
	got, _ := ExtractBoolean("IsBad", text)
	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("unexpected difference: (-got +want)\n%s", diff)
	}

	text = "true"
	expected = "-1235294128:1"
	got, _ = ExtractBoolean("IsBad", text)
	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("unexpected difference: (-got +want)\n%s", diff)
	}
}
