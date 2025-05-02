package uid_test

import (
	"testing"

	"github.com/mdigger/uid"
)

func BenchmarkGeneration(b *testing.B) {
	gen := uid.NewGenerator()
	b.ResetTimer()

	for b.Loop() {
		_ = gen()
	}
}

func BenchmarkParallelGeneration(b *testing.B) {
	gen := uid.NewGenerator()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = gen()
		}
	})
}

func BenchmarkNew(b *testing.B) {
	for b.Loop() {
		_ = uid.New()
	}
}

func BenchmarkParse(b *testing.B) {
	testID := uid.New()
	b.ResetTimer()

	for b.Loop() {
		_, _ = uid.Parse(testID)
	}
}

func BenchmarkFullCycle(b *testing.B) {
	gen := uid.NewGenerator()
	b.ResetTimer()

	for b.Loop() {
		id := gen()
		_, _ = uid.Parse(id)
	}
}

func BenchmarkBatchGeneration(b *testing.B) {
	gen := uid.NewGenerator()
	const batchSize = 1000
	ids := make([]string, batchSize)

	b.ResetTimer()

	for b.Loop() {
		for j := range batchSize {
			ids[j] = gen()
		}
	}
}
