package funk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSlice(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(Get(SliceOf(foo), "ID"), []int{1})
	assert.Equal(Get(SliceOf(foo), "Bar.Name"), []string{"Test"})
	assert.Equal(Get(SliceOf(foo), "Bar"), []*Bar{bar})
}

func TestGetSliceMultiLevel(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(Get(foo, "Bar.Bars.Bar.Name"), []string{"Level2-1", "Level2-2"})
	assert.Equal(Get(SliceOf(foo), "Bar.Bars.Bar.Name"), []string{"Level2-1", "Level2-2"})
}

func TestGetNull(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(Get(foo, "EmptyValue.Int64"), int64(10))
	assert.Equal(Get(SliceOf(foo), "EmptyValue.Int64"), []int64{10})
}

func TestGetNil(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(Get(foo2, "Bar.Name"), nil)
	assert.Equal(Get([]*Foo{foo, foo2}, "Bar.Name"), []string{"Test"})
}

func TestGetSimple(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(Get(foo, "ID"), 1)

	assert.Equal(Get(foo, "Bar.Name"), "Test")

	result := Get(foo, "Bar.Bars.Name")

	assert.Equal(result, []string{"Level1-1", "Level1-2"})
}
