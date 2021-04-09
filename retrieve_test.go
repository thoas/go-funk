package funk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSlice(t *testing.T) {
	is := assert.New(t)

	is.Equal(Get(SliceOf(foo), "ID", false), []int{1})
	is.Equal(Get(SliceOf(foo), "Bar.Name", false), []string{"Test"})
	is.Equal(Get(SliceOf(foo), "Bar", false), []*Bar{bar})
	is.Equal(Get(([]Foo)(nil), "Bar.Name", false), []string{})
	is.Equal(Get([]Foo{}, "Bar.Name", false), []string{})
	is.Equal(Get([]*Foo{}, "Bar.Name", false), []string{})
}

func TestGetSliceMultiLevel(t *testing.T) {
	is := assert.New(t)

	is.Equal(Get(foo, "Bar.Bars.Bar.Name", false), []string{"Level2-1", "Level2-2"})
	is.Equal(Get(SliceOf(foo), "Bar.Bars.Bar.Name", false), []string{"Level2-1", "Level2-2"})
}

func TestGetNull(t *testing.T) {
	is := assert.New(t)

	is.Equal(Get(foo, "EmptyValue.Int64", false), int64(10))
	is.Equal(Get(foo, "ZeroValue", false), nil)
	is.Equal( false, Get(foo, "ZeroBoolValue", true))
	is.Equal( nil, Get(fooUnexported, "unexported", true))
	is.Equal( nil, Get(fooUnexported, "unexported", false))
	is.Equal(Get(foo, "ZeroIntValue", true), 0)
	is.Equal(Get(foo, "ZeroIntPtrValue", true), nil)
	is.Equal(Get(foo, "EmptyValue.Int64", true), int64(10))
	is.Equal(Get(SliceOf(foo), "EmptyValue.Int64", false), []int64{10})
}

func TestGetNil(t *testing.T) {
	is := assert.New(t)
	is.Equal(Get(foo2, "Bar.Name", false), nil)
	is.Equal(Get(foo2, "Bar.Name", true), "")
	is.Equal(Get([]*Foo{foo, foo2}, "Bar.Name", false), []string{"Test"})
	is.Equal(Get([]*Foo{foo, foo2}, "Bar", false), []*Bar{bar})
}

func TestGetMap(t *testing.T) {
	is := assert.New(t)
	m := map[string]interface{}{
		"bar": map[string]interface{}{
			"name": "foobar",
		},
	}

	is.Equal("foobar", Get(m, "bar.name", false))
	is.Equal(nil, Get(m, "foo.name", false))
	is.Equal([]interface{}{"dark", "dark"}, Get([]map[string]interface{}{m1, m2}, "firstname", false))
	is.Equal([]interface{}{"test"}, Get([]map[string]interface{}{m1, m2}, "bar.name", false))
}

func TestGetThroughInterface(t *testing.T) {
	is := assert.New(t)

	is.Equal(Get(foo, "BarInterface.Bars.Bar.Name", false), []string{"Level2-1", "Level2-2"})
	is.Equal(Get(foo, "BarPointer.Bars.Bar.Name", false), []string{"Level2-1", "Level2-2"})
}

func TestGetNotFound(t *testing.T) {
	is := assert.New(t)

	is.Equal(nil, Get(foo, "id", false))
	is.Equal(nil, Get(foo, "id.id", false))
	is.Equal(nil, Get(foo, "Bar.id", false))
	is.Equal(nil, Get(foo, "Bars.id", false))
}

func TestGetSimple(t *testing.T) {
	is := assert.New(t)

	is.Equal(Get(foo, "ID", false), 1)

	is.Equal(Get(foo, "Bar.Name", false), "Test")

	result := Get(foo, "Bar.Bars.Name", false)

	is.Equal(result, []string{"Level1-1", "Level1-2"})
}

func TestGetOrElse(t *testing.T) {
	is := assert.New(t)

	str := "hello world"
	is.Equal("hello world", GetOrElse(&str, "foobar"))
	is.Equal("hello world", GetOrElse(str, "foobar"))
	is.Equal("foobar", GetOrElse(nil, "foobar"))

	t.Run("nil with type", func(t *testing.T) {
		// simple nil comparison is not sufficient for nil with type.
		is.False(interface{}((*string)(nil)) == nil)
		// test GetOrElse coveers this case
		is.Equal("foobar", GetOrElse((*string)(nil), "foobar"))
	})
}
