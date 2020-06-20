package funk

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSubset(t *testing.T) {
	is := assert.New(t)

	r := Subset([]int{1, 2, 4}, []int{1, 2, 3, 4, 5})
	is.Equal(true, r)

	r = Subset([]string{"foo", "bar"},[]string{"foo", "bar", "hello", "bar", "hi"})
	is.Equal(true, r)

	r = Subset([]string{"hello", "foo", "bar", "hello", "bar", "hi"}, []string{})
	is.Equal(false, r)
  
  r = Subset([]string{}, []string{"hello", "foo", "bar", "hello", "bar", "hi"})
	is.Equal(true, r)
}

func TestSubtractString(t *testing.T) {
	is := assert.New(t)

	r = SubsetString([]string{"foo", "bar"},[]string{"foo", "bar", "hello", "bar", "hi"})
	is.Equal(true, r)

	r = SubsetString([]string{"hello", "foo", "bar", "hello", "bar", "hi"}, []string{})
	is.Equal(false, r)
  
  r = SubsetString([]string{}, []string{})
	is.Equal(true, r)
}
