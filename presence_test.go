package funk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var f = &Foo{
	ID:        1,
	FirstName: "Dark",
	LastName:  "Vador",
	Age:       30,
	Bar: &Bar{
		Name: "Test",
	},
}

var b = &Foo{
	ID:        2,
	FirstName: "Florent",
	LastName:  "Messa",
	Age:       28,
}
var c = &Foo{
	ID:        3,
	FirstName: "Harald",
	LastName:  "Nordgren",
	Age:       27,
}

var results = []*Foo{f, c}

func TestContains(t *testing.T) {
	is := assert.New(t)

	is.True(Contains([]string{"foo", "bar"}, "bar"))
	is.True(Contains([...]string{"foo", "bar"}, "bar"))
	is.Panics(func() { Contains(1, 2) })

	is.True(Contains(results, f))
	is.False(Contains(results, nil))
	is.False(Contains(results, b))

	is.True(Contains("florent", "rent"))
	is.False(Contains("florent", "gilles"))

	mapping := ToMap(results, "ID")

	is.True(Contains(mapping, 1))
	is.False(Contains(mapping, 2))
}

func TestEvery(t *testing.T) {
	is := assert.New(t)

	is.True(Every([]string{"foo", "bar", "baz"}, "bar", "foo"))

	is.True(Every(results, f, c))
	is.False(Every(results, nil))
	is.False(Every(results, f, b))

	is.True(Every("florent", "rent", "flo"))
	is.False(Every("florent", "rent", "gilles"))

	mapping := ToMap(results, "ID")

	is.True(Every(mapping, 1, 3))
	is.False(Every(mapping, 2, 3))
}

func TestIndexOf(t *testing.T) {
	is := assert.New(t)

	is.Equal(IndexOf([]string{"foo", "bar"}, "bar"), 1)

	is.Equal(IndexOf(results, f), 0)
	is.Equal(IndexOf(results, b), -1)
}

func TestLastIndexOf(t *testing.T) {
	is := assert.New(t)

	is.Equal(LastIndexOf([]string{"foo", "bar", "bar"}, "bar"), 2)
	is.Equal(LastIndexOf([]int{1, 2, 2, 3}, 2), 2)
	is.Equal(LastIndexOf([]int{1, 2, 2, 3}, 4), -1)
}

func TestFilter(t *testing.T) {
	is := assert.New(t)

	r := Filter([]int{1, 2, 3, 4}, func(x int) bool {
		return x%2 == 0
	})

	is.Equal(r, []int{2, 4})
}

func TestFind(t *testing.T) {
	is := assert.New(t)

	r := Find([]int{1, 2, 3, 4}, func(x int) bool {
		return x%2 == 0
	})

	is.Equal(r, 2)
}
