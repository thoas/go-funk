package funk

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFillMismatchedTypes(t *testing.T) {
	_, err := Fill([]string{"a", "b"}, 1)
	assert.EqualError(t, err, "Cannot fill '[]string' with 'int'")
}

func TestFillUnfillableTypes(t *testing.T) {
	var stringVariable string
	var uint32Variable uint32
	var boolVariable bool

	types := [](interface{}){
		stringVariable,
		uint32Variable,
		boolVariable,
	}

	for _, unfillable := range types {
		_, err := Fill(unfillable, 1)
		assert.EqualError(t, err, "Can only fill slices and arrays")
	}
}

func TestFill(t *testing.T) {
	result, err := Fill([]int{1,2,3}, 1)
	assert.NoError(t, err)
	assert.Equal(t, []int{1,1,1}, result)
}