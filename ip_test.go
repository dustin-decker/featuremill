package featuremill

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExtractIP(t *testing.T) {

	text := "127.0.0.1"

	expected := "1799088460:2130706433"

	got, _ := ExtractIP("src_ip", text)

	if diff := cmp.Diff(expected, got); diff != "" {
		t.Errorf("unexpected difference: (-got +want)\n%s", diff)
	}

	_, err := ExtractIP("src_ip", "invalid IP")
	if err == nil {
		t.Error("expected error with invalid IP")
	}
}
