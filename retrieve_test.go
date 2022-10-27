package funk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSlice(t *testing.T) {
	is := assert.New(t)

	is.Equal(Get(SliceOf(foo), "ID"), []int{1})
	is.Equal(Get(SliceOf(foo), "Bar.Name"), []string{"Test"})
	is.Equal(Get(SliceOf(foo), "Bar"), []*Bar{bar})
	is.Equal(Get(([]Foo)(nil), "Bar.Name"), []string{})
	is.Equal(Get([]Foo{}, "Bar.Name"), []string{})
	is.Equal(Get([]*Foo{}, "Bar.Name"), []string{})
}

func TestGetSliceMultiLevel(t *testing.T) {
	is := assert.New(t)

	is.Equal(Get(foo, "Bar.Bars.Bar.Name"), []string{"Level2-1", "Level2-2"})
	is.Equal(Get(SliceOf(foo), "Bar.Bars.Bar.Name"), []string{"Level2-1", "Level2-2"})
}

func TestGetNull(t *testing.T) {
	is := assert.New(t)

	is.Equal(Get(foo, "EmptyValue.Int64"), int64(10))
	is.Equal(Get(foo, "ZeroValue"), nil)
	is.Equal(false, Get(foo, "ZeroBoolValue", WithAllowZero()))
	is.Equal(nil, Get(fooUnexported, "unexported", WithAllowZero()))
	is.Equal(nil, Get(fooUnexported, "unexported", WithAllowZero()))
	is.Equal(Get(foo, "ZeroIntValue", WithAllowZero()), 0)
	is.Equal(Get(foo, "ZeroIntPtrValue", WithAllowZero()), nil)
	is.Equal(Get(foo, "EmptyValue.Int64", WithAllowZero()), int64(10))
	is.Equal(Get(SliceOf(foo), "EmptyValue.Int64"), []int64{10})
}

func TestGetNil(t *testing.T) {
	is := assert.New(t)
	is.Equal(Get(foo2, "Bar.Name"), nil)
	is.Equal(Get(foo2, "Bar.Name", WithAllowZero()), "")
	is.Equal(Get([]*Foo{foo, foo2}, "Bar.Name"), []string{"Test"})
	is.Equal(Get([]*Foo{foo, foo2}, "Bar"), []*Bar{bar})
}

func TestGetMap(t *testing.T) {
	is := assert.New(t)
	m := map[string]interface{}{
		"bar": map[string]interface{}{
			"name": "foobar",
		},
	}

	is.Equal("foobar", Get(m, "bar.name"))
	is.Equal(nil, Get(m, "foo.name"))
	is.Equal([]interface{}{"dark", "dark"}, Get([]map[string]interface{}{m1, m2}, "firstname"))
	is.Equal([]interface{}{"test"}, Get([]map[string]interface{}{m1, m2}, "bar.name"))
}

func TestGetThroughInterface(t *testing.T) {
	is := assert.New(t)

	is.Equal(Get(foo, "BarInterface.Bars.Bar.Name"), []string{"Level2-1", "Level2-2"})
	is.Equal(Get(foo, "BarPointer.Bars.Bar.Name"), []string{"Level2-1", "Level2-2"})
}

func TestGetNotFound(t *testing.T) {
	is := assert.New(t)

	is.Equal(nil, Get(foo, "id"))
	is.Equal(nil, Get(foo, "id.id"))
	is.Equal(nil, Get(foo, "Bar.id"))
	is.Equal(nil, Get(foo, "Bars.id"))
}

func TestGetSimple(t *testing.T) {
	is := assert.New(t)

	is.Equal(Get(foo, "ID"), 1)

	is.Equal(Get(foo, "Bar.Name"), "Test")

	result := Get(foo, "Bar.Bars.Name")

	is.Equal(result, []string{"Level1-1", "Level1-2"})
}

func TestGetOrElse(t *testing.T) {
	is := assert.New(t)

	str := "hello world"
	is.Equal("hello world", GetOrElse(&str, "foobar"))
	is.Equal("hello world", GetOrElse(str, "foobar"))
	is.Equal("foobar", GetOrElse(nil, "foobar"))

	t.Run("nil with type", func(t *testing.T) {
		// test GetOrElse covers this case
		is.Equal("foobar", GetOrElse((*string)(nil), "foobar"))
	})
}

func TestEmbeddedStruct(t *testing.T) {
	is := assert.New(t)

	root := RootStruct{}
	is.Equal(Get(root, "EmbeddedField"), nil)
	is.Equal(Get(root, "EmbeddedStruct.EmbeddedField"), nil)
}
