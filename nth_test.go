package funk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_nth_valid_n_great_than_array_length(t *testing.T) {
	result := Nth([]int{1, 2, 3}, 4)

	assert.Equal(t, nil, result)
}

func Test_nth_valid_n_less_than_zero_one(t *testing.T) {
	result := Nth([]int32{1, 2, 3}, -3)

	assert.Equal(t, int32(1), result)
}

func Test_nth_valid_n_less_than_zero_two(t *testing.T) {
	result := Nth([]int32{1, 2, 3}, -4)

	assert.Equal(t, nil, result)
}

func Test_nth_valid_int64(t *testing.T) {
	result := Nth([]int64{1, 2, 3}, 2)

	assert.Equal(t, int64(3), result)
}

func Test_nth_valid_float32(t *testing.T) {
	result := Nth([]float32{1, 2, 3}, 1)

	assert.Equal(t, float32(2), result)
}

func Test_nth_valid_float64(t *testing.T) {
	result := Nth([]float64{1.1, 2.2, 3.3}, 2)

	assert.Equal(t, 3.3, result)
}

func Test_nth_valid_bool(t *testing.T) {
	result := Nth([]bool{true, true, true, false, true}, 2)

	assert.Equal(t, true, result)
}

func Test_nth_valid_string(t *testing.T) {
	result := Nth([]string{"a", "b", "c", "d"}, -2)

	assert.Equal(t, "c", result)
}

func Test_nth_valid_interface(t *testing.T) {
	result := Nth([]interface{}{"1.1", true, 2.2, float32(3.3)}, 2)

	assert.Equal(t, 2.2, result)
}
