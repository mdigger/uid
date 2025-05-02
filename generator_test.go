package uid

import (
	"encoding/base32"
	"encoding/binary"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestGenerator(t *testing.T) {
	t.Run("BasicProperties", func(t *testing.T) {
		gen := NewGenerator()
		uid := gen()

		if len(uid) != 13 {
			t.Errorf("expected length 13, got %d", len(uid))
		}

		if !strings.ContainsAny(uid, "0123456789ABCDEFGHIJKLMNOPQRSTUV") {
			t.Errorf("invalid characters in UID: %s", uid)
		}
	})

	t.Run("Uniqueness", func(t *testing.T) {
		gen := NewGenerator()
		const iterations = 1000
		uids := make(map[string]bool, iterations)

		for i := 0; i < iterations; i++ {
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

		decoded, err := base32.HexEncoding.WithPadding(base32.NoPadding).DecodeString(uid)
		if err != nil {
			t.Fatalf("failed to decode UID: %v", err)
		}

		if len(decoded) != 8 {
			t.Fatalf("expected 8 bytes, got %d", len(decoded))
		}

		// check that the timestamp approximately corresponds to the current time
		timestamp := binary.BigEndian.Uint32(decoded[0:4])
		expected := uint32(time.Now().Unix())
		if diff := expected - timestamp; diff > 1 {
			t.Errorf("timestamp mismatch, got %d, expected ~%d (diff %d)",
				timestamp, expected, diff)
		}

		// check that counter is not null
		counter := uint32(decoded[5])<<16 | uint32(decoded[6])<<8 | uint32(decoded[7])
		if counter == 0 {
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
}
