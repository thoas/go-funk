package funk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestIndexOfMaxIssue139Example tests the exact example provided in GitHub issue #139.
func TestIndexOfMaxIssue139Example(t *testing.T) {
	t.Parallel()

	nums := []int64{8, 3, 4, 44, 0}
	result := IndexOfMaxInt64(nums)
	assert.Equal(t, 3, result, "Should return index 3 for the example in issue #139")
	assert.Equal(t, int64(44), nums[result], "The element at returned index should be 44")
}

func TestIndexOfMaxWithArrayNumericInput(t *testing.T) {
	t.Parallel()

	// Test Data
	data1 := []int{8, 3, 4, 44, 0}
	data1dup := []int{44, 3, 4, 44, 0} // duplicate max at beginning
	emptyData := []int{}

	// Calls
	result1 := IndexOfMaxInt(data1)
	result1dup := IndexOfMaxInt(data1dup)

	// Assertions
	assert.Equal(t, 3, result1, "It should return the index of max value in array")
	assert.Equal(t, 0, result1dup, "It should return the first index of duplicate max value")
	assert.Panics(t, func() { IndexOfMaxInt(emptyData) }, "It should panic")
}

func TestIndexOfMaxWithArrayFloatInput(t *testing.T) {
	t.Parallel()

	// Test Data
	data1 := []float64{2, 38.3, 4, 4.4, 4}
	data1dup := []float64{38.3, 2, 4, 38.3, 4} // duplicate max at beginning
	emptyData := []float64{}

	// Calls
	result1 := IndexOfMaxFloat64(data1)
	result1dup := IndexOfMaxFloat64(data1dup)

	// Assertions
	assert.Equal(t, 1, result1, "It should return the index of max value in array")
	assert.Equal(t, 0, result1dup, "It should return the first index of duplicate max value")
	assert.Panics(t, func() { IndexOfMaxFloat64(emptyData) }, "It should panic")
}

func TestIndexOfMaxSingleElement(t *testing.T) {
	t.Parallel()

	// Test Data with single elements
	data1 := []int{42}
	data2 := []float64{3.14}
	data3 := []string{"hello"}

	// Calls
	result1 := IndexOfMaxInt(data1)
	result2 := IndexOfMaxFloat64(data2)
	result3 := IndexOfMaxString(data3)

	// Assertions
	assert.Equal(t, 0, result1, "Single element should return index 0")
	assert.Equal(t, 0, result2, "Single element should return index 0")
	assert.Equal(t, 0, result3, "Single element should return index 0")
}
