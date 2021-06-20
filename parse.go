package uid

import (
	"encoding/base64"
	"encoding/binary"
	"errors"
	"time"
)

// Info describes the parsed information about the unique identifier.
type Info struct {
	Time    time.Time // time ID creation
	Machine [3]byte   // three bytes with the identifier of the computer
	Pid     uint16    // the process ID
	Counter uint32    // counter
}

// Parse parses and returns information about the unique identifier.
func Parse(uid string) (info Info, err error) {
	if len(uid) != 16 {
		err = errors.New("bad uid length")
		return
	}

	data, err := base64.RawURLEncoding.DecodeString(uid)
	if err != nil {
		return
	}

	info.Time = time.Unix(int64(binary.BigEndian.Uint32(data[0:4])), 0)
	copy(info.Machine[:], data[4:7])
	info.Pid = binary.BigEndian.Uint16(data[7:9])
	counter := data[9:12]
	info.Counter = uint32(uint32(counter[0])<<16 | uint32(counter[1])<<8 | uint32(counter[2]))

	return
}
