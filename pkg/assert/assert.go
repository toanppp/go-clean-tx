package assert

import (
	"encoding/json"
	"testing"
)

func StatusCode(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("unexpected status code: got: %d, want: %d", got, want)
	}
}

func JSON(t *testing.T, got string, want any) {
	p, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("cannot marshal expected json: %v", err)
	}

	if got != string(p) {
		t.Errorf("unexpected json: got: %s, want: %s", got, string(p))
	}
}
