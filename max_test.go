package funk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxWithArrayNumericInput(t *testing.T) {
	//Test Data
	d1 := []int{8, 3, 4, 44, 0}
	n1 := []int{}
	d2 := []int8{3, 3, 5, 9, 1}
	n2 := []int8{}
	d3 := []int16{4, 5, 4, 33, 2}
	n3 := []int16{}
	d4 := []int32{5, 3, 21, 15, 3}
	n4 := []int32{}
	d5 := []int64{9, 3, 9, 1, 2}
	n5 := []int64{}
	//Calls
	r1 := MaxInt(d1)
	r2 := MaxInt8(d2)
	r3 := MaxInt16(d3)
	r4 := MaxInt32(d4)
	r5 := MaxInt64(d5)
	// Assertions
	assert.Equal(t, int(44), r1, "It should return the max value in array")
	assert.Panics(t, func() { MaxInt(n1) }, "It should panic")
	assert.Equal(t, int8(9), r2, "It should return the max value in array")
	assert.Panics(t, func() { MaxInt8(n2) }, "It should panic")
	assert.Equal(t, int16(33), r3, "It should return the max value in array")
	assert.Panics(t, func() { MaxInt16(n3) }, "It should panic")
	assert.Equal(t, int32(21), r4, "It should return the max value in array")
	assert.Panics(t, func() { MaxInt32(n4) }, "It should panic")
	assert.Equal(t, int64(9), r5, "It should return the max value in array")
	assert.Panics(t, func() { MaxInt64(n5) }, "It should panic")

}

func TestMaxWithArrayFloatInput(t *testing.T) {
	//Test Data
	d1 := []float64{2, 38.3, 4, 4.4, 4}
	n1 := []float64{}
	d2 := []float32{2.9, 1.3, 4.23, 4.4, 7.7}
	n2 := []float32{}
	//Calls
	r1 := MaxFloat64(d1)
	r2 := MaxFloat32(d2)
	// Assertions
	assert.Equal(t, float64(38.3), r1, "It should return the max value in array")
	assert.Panics(t, func() { MaxFloat64(n1) }, "It should panic")
	assert.Equal(t, float32(7.7), r2, "It should return the max value in array")
	assert.Panics(t, func() { MaxFloat32(n2) }, "It should panic")
}

func TestMaxWithArrayInputWithStrings(t *testing.T) {
	//Test Data
	d1 := []string{"abc", "abd", "cbd"}
	d2 := []string{"abc", "abd", "abe"}
	d3 := []string{"abc", "foo", " "}
	d4 := []string{"abc", "abc", "aaa"}
	n1 := []string{}
	//Calls
	r1 := MaxString(d1)
	r2 := MaxString(d2)
	r3 := MaxString(d3)
	r4 := MaxString(d4)
	// Assertions
	assert.Equal(t, "cbd", r1, "It should print cbd because its first char is max in the list")
	assert.Equal(t, "abe", r2, "It should print abe because its first different char is max in the list")
	assert.Equal(t, "foo", r3, "It should print foo because its first different char is max in the list")
	assert.Equal(t, "abc", r4, "It should print abc because its first different char is max in the list")
	assert.Panics(t, func() { MaxString(n1) }, "It should panic")
}
