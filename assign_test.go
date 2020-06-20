package funk

import (
	"fmt"
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

	is.PanicsWithValue("Type funk.Foo not supported by Set", func() { Set(fooCopy, int64(2), "ID") })

}

func TestSetStructWithFieldNotInitialized(t *testing.T) {
	is := assert.New(t)
	myFoo := &Foo{
		Bar: nil, // we will try to set bar's field
	}
	err := Set(myFoo, "name", "Bar.Name")
	is.NoError(err)
	is.Equal("name", myFoo.Bar.Name)
}

func TestSetSlice(t *testing.T) {
	is := assert.New(t)
	bars := []*Bar{{Name: "a"}, {Name: "b"}}
	// Note: take address is required
	err := Set(&bars, "c", "Name")
	is.NoError(err)
	is.Equal([]*Bar{{Name: "c"}, {Name: "c"}}, bars)
}

func TestSetSliceWithNilElements(t *testing.T) {
	is := assert.New(t)
	bars := []*Bar{nil, nil}
	// Case slice
	err := Set(bars, "c", "Name")
	is.NoError(err)
	is.Equal([]*Bar{{Name: "c"}, {Name: "c"}}, bars)

	// Case ptr to slice
	bars2 := []*Bar{nil, nil}
	err = Set(&bars2, "c", "Name")
	is.NoError(err)
	is.Equal([]*Bar{{Name: "c"}, {Name: "c"}}, bars2)
}

func TestInterface(t *testing.T) {

	type Baz struct {
		Name string
		Itf  interface{}
	}

	baz := Baz{
		Name: "Baz1",
		Itf:  nil,
	}

	var testCases = []struct {
		Path        string
		SetVal      interface{}
		ExpectedBaz Baz
	}{
		// set string field
		{
			"Name",
			"hi",
			Baz{Name: "hi", Itf: nil},
		},
		// set interface{} field
		{
			"Itf",
			"str",
			Baz{Name: "Baz1", Itf: "str"},
		},
	}
	//interface{}(interface{}("c"))

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)
			testBaz := baz // make a copy

			err := Set(&testBaz, tc.SetVal, tc.Path)
			is.NoError(err)
			is.Equal(tc.ExpectedBaz, testBaz)
		})
	}

}
