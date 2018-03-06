package featuremill

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExtractTimestamp(t *testing.T) {

	text := "2018-03-05T03:12:14"

	expected := []string{
		"-627773727:0.166667",
		"2046480298:0.130435",
		"-939123019:0.203390",
	}

	got, _ := ExtractTimestamp("@timestamp", text)

	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("unexpected difference: (-got +want)\n%s", diff)
	}
}
