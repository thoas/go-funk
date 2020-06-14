package funk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetEmptyPath(t *testing.T) {
	is := assert.New(t)
	x := 1
	err := Set(&x, 2, "")
	is.NoError(err)
	is.Equal(2, x)
}

func TestSetStructOneLevel(t *testing.T) {
	is := assert.New(t)
	// copy here because we need to modify
	fooCopy := *foo
	err := Set(&fooCopy, 2, "ID")
	is.NoError(err)
	is.Equal(2, fooCopy.ID)

}

func TestSetStructTwoLevels(t *testing.T) {
	is := assert.New(t)

	// copy here because we need to modify
	fooCopy := *foo

	err := Set(&fooCopy, int64(2), "EmptyValue.Int64")
	is.NoError(err)
	is.Equal(int64(2), fooCopy.EmptyValue.Int64)

}

func TestSetStructNotPtr(t *testing.T) {
	is := assert.New(t)

	// copy here because we need to modify
	fooCopy := *foo

	is.PanicsWithValue("Type funk.Foo cannot be set", func() { Set(fooCopy, int64(2), "ID") })

}

func TestSetStructWithFieldNotInitialized(t *testing.T) {
	is := assert.New(t)
	myFoo := &Foo{
		Bar: nil, // we will try to set bar's field
	}
	err := Set(myFoo, int64(2), "Bar.Name")
	is.EqualError(err, "nil pointer found along the path")
}

func TestSetSlice(t *testing.T) {
	is := assert.New(t)
	bars := []*Bar{{Name: "a"}, {Name: "b"}}
	// Note: take address is required
	err := Set(bars, "c", "Name")
	is.NoError(err)
	is.Equal([]*Bar{{Name: "c"}, {Name: "c"}}, bars)
}
