# Go UID Generator

[![Go Reference](https://pkg.go.dev/badge/github.com/mdigger/uid.svg)](https://pkg.go.dev/github.com/mdigger/uid)

A lightweight, high-performance unique identifier generator for Go, featuring:

- **Time-ordered** identifiers (lexicographically sortable)
- **Base32Hex** encoding (URL-safe, case-insensitive)
- **Compact 7-byte binary format** with optimized components
- **Thread-safe** generation using atomic operations

## Features

- 🚀 Generate 12-character UIDs (7-byte binary → Base32Hex)
- ⏳ Embedded timestamp with millisecond precision
- 🔢 Built-in atomic counter for uniqueness (16-bit)
- 🔄 Bidirectional parsing (encode/decode UIDs)
- 🧪 100% test coverage

## Installation

```shell
go get github.com/mdigger/uid
```

## Usage

### Basic Generation

```go
import "github.com/mdigger/uid"

// Generate a new UID
id := uid.New() // e.g. "A5F3D9E2B4C1"
```

### Advanced Usage

```go
// Create a dedicated generator
generator := uid.NewGenerator()

// Generate multiple IDs
id1 := generator()
id2 := generator()

// Parse existing UID
parsed, err := uid.Parse("A5F3D9E2B4C1")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Timestamp: %s\n", parsed.Timestamp)
fmt.Printf("Millisecond: %d\n", parsed.Millisecond)
fmt.Printf("Counter: %d\n", parsed.Counter)
```

## UID Structure

Each 7-byte UID contains:

| Byte Range | Content               | Size  | Description                     |
|------------|-----------------------|-------|---------------------------------|
| 0-3        | Unix timestamp        | 4 bytes | Seconds since Unix epoch       |
| 4          | Millisecond fraction  | 1 byte  | Nanoseconds/1e6 (0-255)        |
| 5-6        | Atomic counter        | 2 bytes | Process-unique sequence number |

Encoded as Base32Hex (RFC 4648) without padding → 12 characters.

## Performance Considerations

- 🔄 **65,536 unique values/ms** - 16-bit counter provides good uniqueness for most use cases
- 📅 **Timestamp validity** until 2106 (32-bit Unix time)
- ⚠ **High-load systems** - Consider original 3-byte counter version if generating >65K IDs/ms

## Why This Package?

- ✅ **No dependencies** - Pure Go standard library
- ✅ **Production-ready** - Thread-safe and battle-tested
- ✅ **Embeddable metadata** - Extract timestamps without DB lookups
- ✅ **Compact size** - Smaller than UUID while being sortable
- 🔄 **Backward compatible** - Same API as original version

## License

MIT License - See [LICENSE](LICENSE) for details.
