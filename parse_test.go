package uid

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	gen := NewGenerator()
	uid := gen()

	parsed, err := Parse(uid)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	now := time.Now()
	if diff := now.Sub(parsed.Timestamp); diff > 2*time.Second {
		t.Errorf("Timestamp is too far from current time: %v (diff %v)",
			parsed.Timestamp, diff)
	}

	if parsed.Counter == 0 {
		t.Error("Counter should not be zero")
	}
}
