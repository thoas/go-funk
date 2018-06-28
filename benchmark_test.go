package funk

import (
	"math/rand"
	"testing"
)

const (
	seed      = 918234565
	sliceSize = 3614562
)

func sliceGenerator(size uint, r *rand.Rand) (out []int64) {
	for i := uint(0); i < size; i++ {
		out = append(out, rand.Int63())
	}
	return
}

func BenchmarkContains(b *testing.B) {
	r := rand.New(rand.NewSource(seed))
	testData := sliceGenerator(sliceSize, r)
	what := r.Int63()

	b.Run("ContainsInt64", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			ContainsInt64(testData, what)
		}
	})

	b.Run("IndexOfInt64", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			IndexOfInt64(testData, what)
		}
	})

	b.Run("Contains", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Contains(testData, what)
		}
	})
}

func BenchmarkUniq(b *testing.B) {
	r := rand.New(rand.NewSource(seed))
	testData := sliceGenerator(sliceSize, r)

	b.Run("UniqInt64", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			UniqInt64(testData)
		}
	})

	b.Run("Uniq", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Uniq(testData)
		}
	})
}

func BenchmarkSum(b *testing.B) {
	r := rand.New(rand.NewSource(seed))
	testData := sliceGenerator(sliceSize, r)

	b.Run("SumInt64", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			SumInt64(testData)
		}
	})

	b.Run("Sum", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Sum(testData)
		}
	})
}

func BenchmarkDrop(b *testing.B) {
	r := rand.New(rand.NewSource(seed))
	testData := sliceGenerator(sliceSize, r)

	b.Run("DropInt64", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			DropInt64(testData, 1)
		}
	})

	b.Run("Drop", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			Drop(testData, 1)
		}
	})
}
