package funk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnion(t *testing.T) {
	is := assert.New(t)

	r := Union([]int{1, 2, 3, 4}, []int{2, 4, 6})
	is.Equal(r, []int{1, 2, 3, 4, 2, 4, 6})

	r = Union(map[int]int{1: 1, 2: 2}, map[int]int{1: 0, 3: 3})
	is.Equal(r, map[int]int{1: 0, 2: 2, 3: 3})
}

func TestUnionShortcut(t *testing.T) {
	is := assert.New(t)

	r := Union(nil)
	is.Nil(r)

	r = Union([]int{1, 2})
	is.Equal(r, []int{1, 2})
}

func TestUnionStringMap(t *testing.T) {
	is := assert.New(t)

	r := Union(map[string]string{"a": "a", "b": "b"}, map[string]string{"a": "z", "z": "a"}, map[string]string{"z": "z"})
	is.Equal(r, map[string]string{"a": "z", "b": "b", "z": "z"})
}
