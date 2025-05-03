package uid

import (
	"encoding/base32"
	"encoding/binary"
	"math/rand"
	"sync/atomic"
	"time"
)

// NewGenerator generates unique identifiers using Base32Hex encoding.
// Returns 12-character strings without padding.
func NewGenerator() func() string {
	counter := rand.Uint32()

	return func() string {
		var uid [7]byte

		now := time.Now()
		binary.BigEndian.PutUint32(uid[0:4], uint32(now.Unix()))
		uid[4] = byte(now.Nanosecond() / 1e6) // 0-255 ms

		c := atomic.AddUint32(&counter, 1)
		uid[5] = byte(c >> 8)
		uid[6] = byte(c)

		return encoder(uid[:])
	}
}

// New generates a unique identifier using Base32Hex encoding.
func New() string {
	return gen()
}

var (
	encoder = base32.HexEncoding.WithPadding(base32.NoPadding).EncodeToString
	gen     = NewGenerator() // default generator
)
