package uid

import (
	"encoding/base64"
	"encoding/binary"
	"time"
)

// Info describes the parsed information about the unique identifier.
type Info struct {
	Time    time.Time // time ID creation
	Machine []byte    // three bytes with the identifier of the computer
	Pid     uint16    // the process ID
	Counter uint32    // counter
}

// Parse parses and returns information about the unique identifier. If the
// string is not a unique identifier, it returns nil.
func Parse(uid string) *Info {
	if len(uid) != 16 {
		return nil
	}
	data, err := base64.RawURLEncoding.DecodeString(uid)
	if err != nil {
		return nil
	}
	counter := data[9:12]
	return &Info{
		Time:    time.Unix(int64(binary.BigEndian.Uint32(data[0:4])), 0),
		Machine: data[4:7],
		Pid:     binary.BigEndian.Uint16(data[7:9]),
		Counter: uint32(uint32(counter[0])<<16 | uint32(counter[1])<<8 |
			uint32(counter[2])),
	}
}
