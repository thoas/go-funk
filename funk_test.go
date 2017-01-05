package funk

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Bar is
type Bar struct {
	Name string
	Bar  *Bar
	Bars []*Bar
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
	Bars      []*Bar
}

func (f Foo) TableName() string {
	return "foo"
}

var bar *Bar = &Bar{
	Name: "Test",
	Bars: []*Bar{
		&Bar{
			Name: "Level1-1",
			Bar: &Bar{
				Name: "Level2-1",
			},
		},
		&Bar{
			Name: "Level1-2",
			Bar: &Bar{
				Name: "Level2-2",
			},
		},
	},
}

var foo *Foo = &Foo{
	ID:        1,
	FirstName: "Drew",
	LastName:  "Olson",
	Age:       30,
	Bar:       bar,
	Bars: []*Bar{
		bar,
		bar,
	},
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
	assert.False(Contains(results, nil))
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

	result, ok := r.([]string)

	assert.True(ok)
	assert.Equal(len(result), 4)

	r = Map([]int{1, 2, 3, 4}, func(x int) (int, int) {
		return x, x
	})

	resultType := reflect.TypeOf(r)

	assert.True(resultType.Kind() == reflect.Map)
	assert.True(resultType.Key().Kind() == reflect.Int)
	assert.True(resultType.Elem().Kind() == reflect.Int)

	mapping := map[int]string{
		1: "Florent",
		2: "Gilles",
	}

	r = Map(mapping, func(k int, v string) int {
		return k
	})

	assert.True(reflect.TypeOf(r).Kind() == reflect.Slice)
	assert.True(reflect.TypeOf(r).Elem().Kind() == reflect.Int)

	r = Map(mapping, func(k int, v string) (string, string) {
		return fmt.Sprintf("%d", k), v
	})

	resultType = reflect.TypeOf(r)

	assert.True(resultType.Kind() == reflect.Map)
	assert.True(resultType.Key().Kind() == reflect.String)
	assert.True(resultType.Elem().Kind() == reflect.String)
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

func TestForEach(t *testing.T) {
	assert := assert.New(t)

	results := []int{}

	ForEach([]int{1, 2, 3, 4}, func(x int) {
		if x%2 == 0 {
			results = append(results, x)
		}
	})

	assert.Equal(results, []int{2, 4})

	mapping := map[int]string{
		1: "Florent",
		2: "Gilles",
	}

	ForEach(mapping, func(k int, v string) {
		assert.Equal(v, mapping[k])
	})
}

func TestChunk(t *testing.T) {
	assert := assert.New(t)

	results := Chunk([]int{0, 1, 2, 3, 4}, 2).([][]int)

	assert.Len(results, 3)
	assert.Len(results[0], 2)
	assert.Len(results[1], 2)
	assert.Len(results[2], 1)

	assert.Len(Chunk([]int{}, 2), 0)
	assert.Len(Chunk([]int{1}, 2), 1)
}

func TestGetSimple(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(Get(foo, "ID"), 1)
	result := Get(foo, "Bar.Bars.Name")

	assert.Equal(result, []string{"Level1-1", "Level1-2"})

	result = Get(foo, "Bar.Bars.Bar.Name")

	assert.Equal(result, []string{"Level2-1", "Level2-2"})
}

func TestGetSlice(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(Get(SliceOf(foo), "ID"), []int{1})
	assert.Equal(Get(SliceOf(foo), "Bar.Name"), []string{"Test"})
	assert.Equal(Get(SliceOf(foo), "Bar"), []*Bar{bar})
}

func TestFlatten(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(Flatten([][]int{[]int{1, 2}, []int{3, 4}}), []int{1, 2, 3, 4})
}

func TestGetSliceMultiLevel(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(Get(SliceOf(foo), "Bar.Bars.Bar.Name"), []string{"Level2-1", "Level2-2"})
}
