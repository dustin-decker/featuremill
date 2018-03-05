package extractor

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTransformTimestamp(t *testing.T) {

	text := "2018-03-05T03:12:14"

	expected := []string{
		"3667193569:0.166667",
		"2046480298:0.130435",
		"3355844277:0.203390",
	}

	got, _ := TransformTimestamp("@timestamp", text)

	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("unexpected difference: (-got +want)\n%s", diff)
	}
}
