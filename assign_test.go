package funk

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetEmptyPath(t *testing.T) {
	// it is supposed to change the var passed in
	var testCases = []struct {
		// will use path = ""
		Original interface{}
		SetVal   interface{}
	}{
		// int
		{
			Original: 100,
			SetVal:   1,
		},
		// string
		{
			Original: "",
			SetVal:   "val",
		},
		// struct
		{
			Original: Bar{Name: "bar"},
			SetVal:   Bar{Name: "val"},
		},
		// slice
		{
			Original: []Bar{{Name: "bar"}},
			SetVal:   []Bar{{Name: "val1"}, {Name: "val2"}},
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)
			// use empty path
			// must take the addr of the variable to be set
			err := Set(&tc.Original, tc.SetVal, "")
			is.NoError(err)
			is.Equal(tc.Original, tc.SetVal) // original should be set to SetVal
		})
	}
}

func TestSetStructBasicOneLevel(t *testing.T) {
	is := assert.New(t)
	// we set field one by one of baz with expected value
	baz := Foo{
		ID:        100,
		FirstName: "firstname",
		LastName:  "lastname",
		Age:       23,
		Bar:       &Bar{Name: "bar"},
		Bars:      []*Bar{{Name: "1"}},
		EmptyValue: sql.NullInt64{
			Int64: 64,
			Valid: false,
		},
	}
	expected := Foo{
		ID:        1,
		FirstName: "firstname1",
		LastName:  "lastname1",
		Age:       24,
		Bar:       &Bar{Name: "b1", Bar: &Bar{Name: "b2"}},
		Bars:      []*Bar{{Name: "1"}, {Name: "2"}},
		EmptyValue: sql.NullInt64{
			Int64: 11,
			Valid: true,
		},
	}
	err := Set(&baz, 1, "ID")
	is.NoError(err)
	err = Set(&baz, expected.FirstName, "FirstName")
	is.NoError(err)
	err = Set(&baz, expected.LastName, "LastName")
	is.NoError(err)
	err = Set(&baz, expected.Age, "Age")
	is.NoError(err)
	err = Set(&baz, expected.Bar, "Bar")
	is.NoError(err)
	err = Set(&baz, expected.Bars, "Bars")
	is.NoError(err)
	err = Set(&baz, expected.EmptyValue, "EmptyValue")
	is.NoError(err)
	is.Equal(baz, expected)
}

func TestSetStructTwoLevels(t *testing.T) {
	is := assert.New(t)

	// copy here because we need to modify
	fooCopy := *foo

	err := Set(&fooCopy, int64(2), "EmptyValue.Int64")
	is.NoError(err)
	is.Equal(int64(2), fooCopy.EmptyValue.Int64)

}

func TestSetStructWithCyclicStruct(t *testing.T) {
	is := assert.New(t)

	testBar := Bar{
		Name: "testBar",
		Bar:  nil,
	}
	testBar.Bar = &testBar

	err := Set(&testBar, "val", "Bar.Bar.Name")
	is.NoError(err)
	is.Equal("val", testBar.Name)
}

// func TestPointerCycle(t *testing.T) {
// 	is := assert.New(t)

// 	x := 10

// 	intPtr := &x
// 	intPtrPtr := &intPtr
// 	intPtrPtr = intPtrPtr

// }

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

func TestSetSlicePassByPtr(t *testing.T) {

	var testCases = []struct {
		Original interface{} // slice or array
		Path     string
		SetVal   interface{}
		Expected interface{}
	}{
		// Set Slice itself
		{
			Original: []*Bar{},
			Path:     "", // empty path means set the passed in ptr itself
			SetVal:   []*Bar{{Name: "bar"}},
			Expected: []*Bar{{Name: "bar"}},
		},
		// empty slice
		{
			Original: []*Bar{},
			Path:     "Name",
			SetVal:   "val",
			Expected: []*Bar{},
		},
		// slice of ptr
		{
			Original: []*Bar{{Name: "a"}, {Name: "b"}},
			Path:     "Name",
			SetVal:   "val",
			Expected: []*Bar{{Name: "val"}, {Name: "val"}},
		},
		// slice of struct
		{
			Original: []Bar{{Name: "a"}, {Name: "b"}},
			Path:     "Name",
			SetVal:   "val",
			Expected: []Bar{{Name: "val"}, {Name: "val"}},
		},
		// slice of empty ptr
		{
			Original: []*Bar{nil, nil},
			Path:     "Name",
			SetVal:   "val",
			Expected: []*Bar{{Name: "val"}, {Name: "val"}},
		},
		// mix of init ptr and nil ptr
		{
			Original: []*Bar{{Name: "bar"}, nil},
			Path:     "Name",
			SetVal:   "val",
			Expected: []*Bar{{Name: "val"}, {Name: "val"}},
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)
			// take the addr and then pass it in
			err := Set(&tc.Original, tc.SetVal, tc.Path)
			is.NoError(err)
			is.Equal(tc.Expected, tc.Original)
		})
	}
}

func TestSetSlicePassDirectly(t *testing.T) {
	// TODO merge with above
	var testCases = []struct {
		Original interface{} // slice or array
		Path     string
		SetVal   interface{}
		Expected interface{}
	}{
		// Set Slice itself
		// {
		// 	Original: []*Bar{},
		// 	Path:     "", // empty path means set the passed in ptr itself
		// 	SetVal:   []*Bar{{Name: "bar"}},
		// 	Expected: []*Bar{{Name: "bar"}},
		// }, // this will fail
		// empty slice
		{
			Original: []*Bar{},
			Path:     "Name",
			SetVal:   "val",
			Expected: []*Bar{},
		},
		// slice of ptr
		{
			Original: []*Bar{{Name: "a"}, {Name: "b"}},
			Path:     "Name",
			SetVal:   "val",
			Expected: []*Bar{{Name: "val"}, {Name: "val"}},
		},
		// slice of struct
		{
			Original: []Bar{{Name: "a"}, {Name: "b"}},
			Path:     "Name",
			SetVal:   "val",
			Expected: []Bar{{Name: "val"}, {Name: "val"}},
		},
		// slice of empty ptr
		{
			Original: []*Bar{nil, nil},
			Path:     "Name",
			SetVal:   "val",
			Expected: []*Bar{{Name: "val"}, {Name: "val"}},
		},
		// mix of init ptr and nil ptr
		{
			Original: []*Bar{{Name: "bar"}, nil},
			Path:     "Name",
			SetVal:   "val",
			Expected: []*Bar{{Name: "val"}, {Name: "val"}},
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)
			// Not take ptr, pass directly
			err := Set(tc.Original, tc.SetVal, tc.Path)
			is.NoError(err)
			is.Equal(tc.Expected, tc.Original)
		})
	}
}

func TestInterface(t *testing.T) {

	type Baz struct {
		Name string
		Itf  interface{}
	}

	var testCases = []struct {
		OriginalBaz Baz
		Path        string
		SetVal      interface{}
		ExpectedBaz Baz
	}{
		// set string field
		{
			Baz{Name: "", Itf: nil},
			"Name",
			"hi",
			Baz{Name: "hi", Itf: nil},
		},
		// set interface{} field
		{
			Baz{Name: "", Itf: nil},
			"Itf",
			"str",
			Baz{Name: "", Itf: "str"},
		},
		// set field of the interface{} field
		// TODO: set uninitialized interface{} should fail
		// Note: interface of struct (not ptr to struct) should fail
		{
			Baz{Name: "", Itf: &Baz{Name: "", Itf: nil}},
			"Itf.Name",
			"Baz2",
			Baz{Name: "", Itf: &Baz{Name: "Baz2", Itf: nil}},
		},
		// interface two level
		{
			Baz{Name: "", Itf: &Baz{Name: "", Itf: nil}},
			"Itf.Itf",
			"val",
			Baz{Name: "", Itf: &Baz{Name: "", Itf: "val"}},
		},
		// uninit interface
		// {
		// 	Baz{Name: "", Itf: &Baz{Name: "", Itf: nil}},
		// 	"Itf.Itf.Name",
		// 	"val",
		// 	Baz{Name: "", Itf: &Baz{Name: "", Itf: &Baz{Name: "val"}}},
		// },

	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			err := Set(&tc.OriginalBaz, tc.SetVal, tc.Path)
			is.NoError(err)
			is.Equal(tc.ExpectedBaz, tc.OriginalBaz)
		})
	}

}

func TestInterfaceTry(t *testing.T) {
	// type Baz struct {
	// 	Name string
	// 	Itf  interface{}
	// }

	type Baz struct {
		Name string
		Itf  Bar
	}
	baz := Baz{Name: "", Itf: Bar{Name: "dummy"}}

	f := reflect.ValueOf(&baz).Elem().FieldByName("Itf")
	if !f.CanSet() {
		t.Error("cannot set f")
	}

	ff := f.FieldByName("Name")
	if !ff.CanAddr() {
		t.Error("ff cannot addr")
	}

	if !ff.CanSet() {
		t.Error(ff.String())
		t.Error("cannot set ff")
	}
}
