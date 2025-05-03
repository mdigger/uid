package uid

import (
	"strings"
	"sync"
	"testing"
	"time"
)

func TestGenerator(t *testing.T) {
	t.Run("BasicProperties", func(t *testing.T) {
		gen := NewGenerator()
		uid := gen()

		if len(uid) != 12 {
			t.Errorf("expected length 12, got %d", len(uid))
		}

		if !strings.ContainsAny(uid, "0123456789ABCDEFGHIJKLMNOPQRSTUV") {
			t.Errorf("invalid characters in UID: %s", uid)
		}
	})

	t.Run("Uniqueness", func(t *testing.T) {
		gen := NewGenerator()
		const iterations = 1000
		uids := make(map[string]bool, iterations)

		for range iterations {
			uid := gen()
			if uids[uid] {
				t.Fatalf("duplicate UID generated: %s", uid)
			}
			uids[uid] = true
		}
	})

	t.Run("Structure", func(t *testing.T) {
		gen := NewGenerator()
		uid := gen()

		parsed, err := Parse(uid)
		if err != nil {
			t.Fatalf("failed to parse UID: %v", err)
		}

		// check that the timestamp approximately corresponds to the current time
		expected := time.Now().Unix()
		if diff := expected - parsed.Timestamp.Unix(); diff > 1 {
			t.Errorf("timestamp mismatch, got %d, expected ~%d (diff %d)",
				parsed.Timestamp.Unix(), expected, diff)
		}

		// check that the milliseconds are in the acceptable range
		if parsed.Millisecond > 255 {
			t.Errorf("milliseconds out of range: %d", parsed.Millisecond)
		}

		// check that the counter is not zero
		if parsed.Counter == 0 {
			t.Error("counter should not be zero")
		}
	})

	t.Run("Concurrency", func(t *testing.T) {
		gen := NewGenerator()
		var wg sync.WaitGroup
		uids := make(chan string, 10000)

		for range 100 {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for range 100 {
					uids <- gen()
				}
			}()
		}

		go func() {
			wg.Wait()
			close(uids)
		}()

		unique := make(map[string]bool)
		for uid := range uids {
			if unique[uid] {
				t.Errorf("duplicate UID detected in concurrent test: %s", uid)
			}
			unique[uid] = true
		}
	})

	t.Run("CounterRollover", func(t *testing.T) {
		gen := NewGenerator()

		// generating the UID before the counter overflows
		var lastCounter uint32
		var uid string

		for range 100000 {
			uid = gen()
			parsed, err := Parse(uid)
			if err != nil {
				t.Fatalf("failed to parse UID: %v", err)
			}

			// if the counter decreased, it means there was an overflow
			if parsed.Counter < lastCounter && lastCounter > 60000 {
				t.Logf("counter rollover detected at %d -> %d", lastCounter, parsed.Counter)
				return
			}
			lastCounter = parsed.Counter
		}

		t.Error("counter rollover not detected after 100000 iterations")
	})
}
