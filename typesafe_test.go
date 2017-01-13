package funk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsInt(t *testing.T) {
	is := assert.New(t)

	is.True(ContainsInt([]int{1, 2, 3, 4}, 4))
	is.False(ContainsInt([]int{1, 2, 3, 4}, 5))

	is.True(ContainsInt64([]int64{1, 2, 3, 4}, 4))
	is.False(ContainsInt64([]int64{1, 2, 3, 4}, 5))
}

func TestContainsString(t *testing.T) {
	is := assert.New(t)

	is.True(ContainsString([]string{"flo", "gilles"}, "flo"))
	is.False(ContainsString([]string{"flo", "gilles"}, "alex"))
}

func TestContainsFloat(t *testing.T) {
	is := assert.New(t)

	is.True(ContainsFloat64([]float64{0.1, 0.2}, 0.1))
	is.False(ContainsFloat64([]float64{0.1, 0.2}, 0.3))

	is.True(ContainsFloat32([]float32{0.1, 0.2}, 0.1))
	is.False(ContainsFloat32([]float32{0.1, 0.2}, 0.3))
}

func TestSumNumeral(t *testing.T) {
	is := assert.New(t)

	is.Equal(SumInt([]int{1, 2, 3}), 6)
	is.Equal(SumInt64([]int64{1, 2, 3}), int64(6))

	is.Equal(SumFloat32([]float32{0.1, 0.2, 0.1}), float32(0.4))
	is.Equal(SumFloat64([]float64{0.1, 0.2, 0.1}), float64(0.4))
}

func TestTypesafeReverse(t *testing.T) {
	is := assert.New(t)

	is.Equal(ReverseInt([]int{1, 2, 3, 4}), []int{4, 3, 2, 1})
	is.Equal(ReverseInt64([]int64{1, 2, 3, 4}), []int64{4, 3, 2, 1})
	is.Equal(ReverseString([]string{"flo", "gilles"}), []string{"gilles", "flo"})
	is.Equal(ReverseFloat64([]float64{0.1, 0.2, 0.3}), []float64{0.3, 0.2, 0.1})
	is.Equal(ReverseFloat32([]float32{0.1, 0.2, 0.3}), []float32{0.3, 0.2, 0.1})
}

func TestTypesafeIndexOf(t *testing.T) {
	is := assert.New(t)

	is.Equal(IndexOfString([]string{"foo", "bar"}, "bar"), 1)
	is.Equal(IndexOfString([]string{"foo", "bar"}, "flo"), -1)

	is.Equal(IndexOfInt([]int{0, 1, 2}, 1), 1)
	is.Equal(IndexOfInt([]int{0, 1, 2}, 3), -1)

	is.Equal(IndexOfInt64([]int64{0, 1, 2}, 1), 1)
	is.Equal(IndexOfInt64([]int64{0, 1, 2}, 3), -1)

	is.Equal(IndexOfFloat64([]float64{0.1, 0.2, 0.3}, 0.2), 1)
	is.Equal(IndexOfFloat64([]float64{0.1, 0.2, 0.3}, 0.4), -1)
}

func TestTypesafeLastIndexOf(t *testing.T) {
	is := assert.New(t)

	is.Equal(LastIndexOfString([]string{"foo", "bar", "bar"}, "bar"), 2)
	is.Equal(LastIndexOfInt([]int{1, 2, 2, 3}, 2), 2)
	is.Equal(LastIndexOfInt64([]int64{1, 2, 2, 3}, 4), -1)
}

func TestTypesafeUniq(t *testing.T) {
	is := assert.New(t)

	is.Equal(UniqInt64([]int64{0, 1, 1, 2, 3, 0, 0, 12}), []int64{0, 1, 2, 3, 12})
	is.Equal(UniqInt([]int{0, 1, 1, 2, 3, 0, 0, 12}), []int{0, 1, 2, 3, 12})
	is.Equal(UniqFloat64([]float64{0.0, 0.1, 0.1, 0.2, 0.3, 0.0, 0.0, 0.12}), []float64{0.0, 0.1, 0.2, 0.3, 0.12})
	is.Equal(UniqString([]string{"foo", "bar", "foo", "bar"}), []string{"foo", "bar"})
}

func TestTypesafeShuffle(t *testing.T) {
	is := assert.New(t)

	initial := []int{1, 2, 3, 5}

	results := ShuffleInt(initial)

	is.Len(results, 4)

	for _, entry := range initial {
		is.True(ContainsInt(results, entry))
	}
}
