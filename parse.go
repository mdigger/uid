package uid

import (
	"encoding/base32"
	"encoding/binary"
	"errors"
	"time"
)

// UID contains the disassembled components of the identifier.
type UID struct {
	Timestamp   time.Time
	Millisecond uint8
	Counter     uint32
}

// ErrInvalidLength means that the string does not match the expected length.
var ErrInvalidLength = errors.New("invalid uid length")

// Parse parses the UID string into its component parts.
func Parse(uid string) (*UID, error) {
	decoded, err := base32.HexEncoding.WithPadding(base32.NoPadding).DecodeString(uid)
	if err != nil {
		return nil, err
	}

	if len(decoded) != 7 {
		return nil, ErrInvalidLength
	}

	var bytes [7]byte
	copy(bytes[:], decoded)

	timestamp := binary.BigEndian.Uint32(bytes[0:4])
	millis := bytes[4]
	counter := uint32(bytes[5])<<8 | uint32(bytes[6])

	return &UID{
		Timestamp:   time.Unix(int64(timestamp), 0),
		Millisecond: millis,
		Counter:     counter,
	}, nil
}
