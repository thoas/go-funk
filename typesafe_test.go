package funk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsInt(t *testing.T) {
	assert := assert.New(t)

	assert.True(ContainsInt([]int{1, 2, 3, 4}, 4))
	assert.False(ContainsInt([]int{1, 2, 3, 4}, 5))

	assert.True(ContainsInt64([]int64{1, 2, 3, 4}, 4))
	assert.False(ContainsInt64([]int64{1, 2, 3, 4}, 5))
}

func TestContainsString(t *testing.T) {
	assert := assert.New(t)

	assert.True(ContainsString([]string{"flo", "gilles"}, "flo"))
	assert.False(ContainsString([]string{"flo", "gilles"}, "alex"))
}

func TestContainsFloat(t *testing.T) {
	assert := assert.New(t)

	assert.True(ContainsFloat64([]float64{0.1, 0.2}, 0.1))
	assert.False(ContainsFloat64([]float64{0.1, 0.2}, 0.3))

	assert.True(ContainsFloat32([]float32{0.1, 0.2}, 0.1))
	assert.False(ContainsFloat32([]float32{0.1, 0.2}, 0.3))
}

func TestSumNumeral(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(SumInt([]int{1, 2, 3}), 6)
	assert.Equal(SumInt64([]int64{1, 2, 3}), int64(6))

	assert.Equal(SumFloat32([]float32{0.1, 0.2, 0.1}), float32(0.4))
	assert.Equal(SumFloat64([]float64{0.1, 0.2, 0.1}), float64(0.4))
}

func TestTypesafeReverse(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(ReverseInt([]int{1, 2, 3, 4}), []int{4, 3, 2, 1})
	assert.Equal(ReverseInt64([]int64{1, 2, 3, 4}), []int64{4, 3, 2, 1})
	assert.Equal(ReverseString([]string{"flo", "gilles"}), []string{"gilles", "flo"})
	assert.Equal(ReverseFloat64([]float64{0.1, 0.2, 0.3}), []float64{0.3, 0.2, 0.1})
	assert.Equal(ReverseFloat32([]float32{0.1, 0.2, 0.3}), []float32{0.3, 0.2, 0.1})
}

func TestTypesafeIndexOf(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(IndexOfString([]string{"foo", "bar"}, "bar"), 1)
	assert.Equal(IndexOfString([]string{"foo", "bar"}, "flo"), -1)

	assert.Equal(IndexOfInt([]int{0, 1, 2}, 1), 1)
	assert.Equal(IndexOfInt([]int{0, 1, 2}, 3), -1)

	assert.Equal(IndexOfInt64([]int64{0, 1, 2}, 1), 1)
	assert.Equal(IndexOfInt64([]int64{0, 1, 2}, 3), -1)

	assert.Equal(IndexOfFloat64([]float64{0.1, 0.2, 0.3}, 0.2), 1)
	assert.Equal(IndexOfFloat64([]float64{0.1, 0.2, 0.3}, 0.4), -1)
}
