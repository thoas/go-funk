package funk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxWithArrayNumericInput(t *testing.T) {

	//Test Data
	d1 := []int{8, 3, 4, 44, 0}
	d2 := []int8{3, 3, 5, 9, 1}
	d3 := []int16{4, 5, 4, 33, 2}
	d4 := []int32{5, 3, 21, 15, 3}
	d5 := []int64{9, 3, 9, 1, 2}

	r1 := Max(d1)
	r2 := Max(d2)
	r3 := Max(d3)
	r4 := Max(d4)
	r5 := Max(d5)

	// Assertions
	assert.Equal(t, int(44), r1, "It should return the max value in array")
	assert.Equal(t, int8(9), r2, "It should return the max value in array")
	assert.Equal(t, int16(33), r3, "It should return the max value in array")
	assert.Equal(t, int32(21), r4, "It should return the max value in array")
	assert.Equal(t, int64(9), r5, "It should return the max value in array")

}

func TestMaxWithArrayFloatInput(t *testing.T) {

	//Test Data
	d1 := []float64{2, 38.3, 4, 4.4, 4}
	d2 := []float32{2.9, 1.3, 4.23, 4.4, 7.7}

	r1 := Max(d1)
	r2 := Max(d2)

	// Assertions
	assert.Equal(t, float64(38.3), r1, "It should return the max value in array")
	assert.Equal(t, float32(7.7), r2, "It should return the max value in array")

}

func TestMaxWithArrayInputWithStrings(t *testing.T) {

	//Test Data
	d1 := []string{"abc", "abd", "cbd"}
	d2 := []string{"abc", "abd", "abe"}

	r1 := Max(d1)
	r2 := Max(d2)

	// Assertions
	assert.Equal(t, "cbd", r1, "It should print cbd because its first char is max in the list")
	assert.Equal(t, "abe", r2, "It should print abe because its first different char is max in the list")
}

func TestMaxWithNilInput(t *testing.T) {

	r := Max(nil)

	// Assertions
	assert.Nil(t, r, "If nil is passed, the function should return nil")

}

func TestMaxWithDifferentTypesInput(t *testing.T) {


	d := []interface{}{"abc", 4, 5.6}

	r := Max(d)

	// Assertions
	assert.Nil(t, r, "If different types are passed in a slice, the function should return nil")

}

