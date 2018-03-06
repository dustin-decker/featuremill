package featuremill

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTransformDate(t *testing.T) {

	text := "2018-03-05T03:12:14"

	expected := []string{
		"3667193569:0.166667",
		"210877497:0.250000",
	}

	got, _ := TransformDate("@timestamp", text)

	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("unexpected difference: (-got +want)\n%s", diff)
	}
}
