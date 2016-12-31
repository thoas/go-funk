package fn

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Bar is
type Bar struct {
	Name string
}

func (b Bar) TableName() string {
	return "bar"
}

// Foo is
type Foo struct {
	ID        int
	FirstName string `tag_name:"tag 1"`
	LastName  string `tag_name:"tag 2"`
	Age       int    `tag_name:"tag 3"`
	Bar       *Bar
}

func (f Foo) TableName() string {
	return "foo"
}

func TestSliceOf(t *testing.T) {
	assert := assert.New(t)

	f := &Foo{
		ID:        1,
		FirstName: "Drew",
		LastName:  "Olson",
		Age:       30,
		Bar: &Bar{
			Name: "Test",
		},
	}

	result := SliceOf(f)

	resultType := reflect.TypeOf(result)

	assert.True(resultType.Kind() == reflect.Slice)
	assert.True(resultType.Elem().Kind() == reflect.Ptr)

	elemType := resultType.Elem().Elem()

	assert.True(elemType.Kind() == reflect.Struct)

	value := reflect.ValueOf(result)

	assert.Equal(value.Len(), 1)

	_, ok := value.Index(0).Interface().(*Foo)

	assert.True(ok)
}

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
	assert.False(Contains(results, b))

	assert.True(Contains("florent", "rent"))
	assert.False(Contains("florent", "gilles"))

	mapping := ToMap(results, "ID")

	assert.True(Contains(mapping, 1))
	assert.False(Contains(mapping, 2))
}

func TestToMap(t *testing.T) {
	assert := assert.New(t)

	f := &Foo{
		ID:        1,
		FirstName: "Drew",
		LastName:  "Olson",
		Age:       30,
		Bar: &Bar{
			Name: "Test",
		},
	}

	results := []*Foo{f}

	instanceMap := ToMap(results, "ID")

	assert.True(reflect.TypeOf(instanceMap).Kind() == reflect.Map)

	mapping, ok := instanceMap.(map[int]*Foo)

	assert.True(ok)

	for _, result := range results {
		item, ok := mapping[result.ID]

		assert.True(ok)
		assert.True(reflect.TypeOf(item).Kind() == reflect.Ptr)
		assert.True(reflect.TypeOf(item).Elem().Kind() == reflect.Struct)

		assert.Equal(item.ID, result.ID)
	}
}

func TestMap(t *testing.T) {
	assert := assert.New(t)

	r := Map([]int{1, 2, 3, 4}, func(x int) string {
		return "Hello"
	})

	_, ok := r.([]string)

	assert.True(ok)
}

func TestFilter(t *testing.T) {
	assert := assert.New(t)

	r := Filter([]int{1, 2, 3, 4}, func(x int) bool {
		return x%2 == 0
	})

	assert.Equal(r, []int{2, 4})
}

func TestForEach(t *testing.T) {
	assert := assert.New(t)

	results := []int{}

	ForEach([]int{1, 2, 3, 4}, func(x int) {
		if x%2 == 0 {
			results = append(results, x)
		}
	})

	assert.Equal(results, []int{2, 4})
}
