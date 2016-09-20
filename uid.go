// Package uid is a generator of globally unique identifiers.
//
// The algorithm used to generate globally unique identifiers based on the same
// principle that is used to generate unique IDs in MongoDB. The unique ID is a
// 12 byte sequence consisting of the time of generation, computer ID, and
// process, as well as counter.
//
// The main difference is that the identifier is represented as a string using
// base64-encoding. It also supports a function to quickly parse this string,
// which returns information about all the values, which are assembled from the
// identifier.
package uid

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

// objectIDCounter is a counter which is automatically incremented after
// each generation of a new unique key. The initial value of the given key
// set random.
var objectIDCounter = randInt()

// randInt returns a random uint32 number.
func randInt() uint32 {
	b := make([]byte, 3)
	if _, err := rand.Reader.Read(b); err != nil {
		panic(fmt.Errorf("Cannot generate random number: %v;", err))
	}
	return uint32(b[0])<<16 | uint32(b[1])<<8 | uint32(b[2])
}

// machineId ID stores of the unique machine ID. Used when generating random ID.
var machineID = readMachineID()

// readMachineId initialisere the value of the computer ID.
func readMachineID() []byte {
	id := make([]byte, 3)
	if hostname, err := os.Hostname(); err == nil {
		hw := md5.New()
		hw.Write([]byte(hostname))
		copy(id, hw.Sum(nil))
	} else {
		// Fallback to rand number if machine id can't be gathered
		if _, randErr := rand.Reader.Read(id); randErr != nil {
			panic(
				fmt.Errorf("Cannot get hostname nor generate a random number: %v; %v",
					err, randErr))
		}
	}
	return id
}

// New generates a string with a globally unique identifier.
//
// To generate a unique ID using the same algorithm as at MongoDB. The only
// difference is in the format generated identifier in string form: to do this,
// the library uses base64 instead of hex representation.
func New() string {
	id := make([]byte, 12)
	// Timestamp, 4 bytes, big endian
	binary.BigEndian.PutUint32(id, uint32(time.Now().Unix()))
	// Machine, first 3 bytes of md5(hostname)
	id[4] = machineID[0]
	id[5] = machineID[1]
	id[6] = machineID[2]
	// Pid, 2 bytes, specs don't specify endianness, but we use big endian.
	pid := os.Getpid()
	id[7] = byte(pid >> 8)
	id[8] = byte(pid)
	// Increment, 3 bytes, big endian
	i := atomic.AddUint32(&objectIDCounter, 1)
	id[9] = byte(i >> 16)
	id[10] = byte(i >> 8)
	id[11] = byte(i)
	return base64.RawURLEncoding.EncodeToString(id)
}
