package funk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubtract(t *testing.T) {
	is := assert.New(t)

	r := Subtract([]int{1, 2, 3, 4, 5}, []int{2, 4, 6})
	is.Equal([]int{1, 3, 5}, r)

	r = Subtract([]string{"foo", "bar", "hello", "bar", "hi"}, []string{"foo", "bar"})
	is.Equal([]string{"hello", "hi"}, r)
}

func TestSubtractString(t *testing.T) {
	is := assert.New(t)

	r := SubtractString([]string{"foo", "bar", "hello", "bar"}, []string{"foo", "bar"})
	is.Equal([]string{"hello"}, r)
}
