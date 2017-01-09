package funk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	assert := assert.New(t)

	assert.True(Contains([]string{"foo", "bar"}, "bar"))

	f := &Foo{
		ID:        1,
		FirstName: "Drew",
		LastName:  "Olson",
		Age:       30,
		Bar: &Bar{
			Name: "Test",
		},
	}

	b := &Foo{
		ID:        2,
		FirstName: "Florent",
		LastName:  "Messa",
		Age:       28,
	}

	results := []*Foo{f}

	assert.True(Contains(results, f))
	assert.False(Contains(results, nil))
	assert.False(Contains(results, b))

	assert.True(Contains("florent", "rent"))
	assert.False(Contains("florent", "gilles"))

	mapping := ToMap(results, "ID")

	assert.True(Contains(mapping, 1))
	assert.False(Contains(mapping, 2))
}

func TestIndexOf(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(IndexOf([]string{"foo", "bar"}, "bar"), 1)

	f := &Foo{
		ID:        1,
		FirstName: "Drew",
		LastName:  "Olson",
		Age:       30,
		Bar: &Bar{
			Name: "Test",
		},
	}

	b := &Foo{
		ID:        2,
		FirstName: "Florent",
		LastName:  "Messa",
		Age:       28,
	}

	results := []*Foo{f}

	assert.Equal(IndexOf(results, f), 0)
	assert.Equal(IndexOf(results, b), -1)
}

func TestLastIndexOf(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(LastIndexOf([]string{"foo", "bar", "bar"}, "bar"), 2)
	assert.Equal(LastIndexOf([]int{1, 2, 2, 3}, 2), 2)
	assert.Equal(LastIndexOf([]int{1, 2, 2, 3}, 4), -1)
}

func TestFilter(t *testing.T) {
	assert := assert.New(t)

	r := Filter([]int{1, 2, 3, 4}, func(x int) bool {
		return x%2 == 0
	})

	assert.Equal(r, []int{2, 4})
}

func TestFind(t *testing.T) {
	assert := assert.New(t)

	r := Find([]int{1, 2, 3, 4}, func(x int) bool {
		return x%2 == 0
	})

	assert.Equal(r, 2)
}
