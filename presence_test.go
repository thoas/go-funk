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

type Person struct {
	name string
	age  int
}

func TestContains(t *testing.T) {
	is := assert.New(t)

	is.True(Contains([]string{"foo", "bar"}, "bar"))

	is.True(Contains(results, f))
	is.False(Contains(results, nil))
	is.False(Contains(results, b))

	is.False(Contains([]string{"florent"}, "gilles"))
}

func TestEvery(t *testing.T) {
	is := assert.New(t)

	is.True(Every([]string{"foo", "bar", "baz"}, "bar", "foo"))

	is.True(Every(results, f, c))
	is.False(Every(results, nil))
	is.False(Every(results, f, b))

	is.True(Every([]string{"florent"}, "rent", "flo"))
	is.False(Every([]string{"florent"}, "rent", "gilles"))
}

func TestSome(t *testing.T) {
	is := assert.New(t)

	is.True(Some([]string{"foo", "bar", "baz"}, "foo"))
	is.True(Some([]string{"foo", "bar", "baz"}, "foo", "qux"))

	is.True(Some(results, f))
	is.False(Some(results, b))
	is.False(Some(results, nil))
	is.True(Some(results, f, b))

	persons := []Person{
		{name: "Zeeshan", age: 23},
		{name: "Bob", age: 26},
	}

	person := Person{"Zeeshan", 23}
	person2 := Person{"Alice", 23}
	person3 := Person{"John", 26}

	is.True(Some(persons, person, person2))
	is.False(Some(persons, person2, person3))
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

	r1, ok1 := Find([]int{1, 2, 3, 4}, func(x int) bool {
		return x%2 == 0
	})
	r2, ok2 := Find([]int{1, 2, 3, 4}, func(x int) bool {
		return x%2 == 5
	})

	is.Equal(r1, 2)
	is.True(ok1)

	is.Equal(r2, 0)
	is.False(ok2)
}

func TestFindKey(t *testing.T) {
	is := assert.New(t)

	k, r := FindKey(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}, func(x int) bool {
		return x == 2
	})

	is.Equal(r, 2)
	is.Equal(k, "b")

	k1, r1 := FindKey([]int{1, 2, 3, 4}, func(x int) bool {
		return x%2 == 0
	})
	is.Equal(r1, 2)
	is.Equal(k1, 1)
}
