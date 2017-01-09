package funk

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestToMap(t *testing.T) {
	assert := assert.New(t)

	f := &Foo{
		ID:        1,
		FirstName: "Dark",
		LastName:  "Vador",
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

func TestFlattenDeep(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(FlattenDeep([][]int{[]int{1, 2}, []int{3, 4}}), []int{1, 2, 3, 4})
}

func TestShuffle(t *testing.T) {
	results := Shuffle([]int{0, 1, 2, 3, 4})

	assert := assert.New(t)

	assert.Len(results, 5)

	fmt.Println(results)
}
