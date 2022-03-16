package funk

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet_EmptyPath(t *testing.T) {
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

func TestSet_StructBasicOneLevel(t *testing.T) {
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

func TestSetStruct_MultiLevels(t *testing.T) {

	var testCases = []struct {
		Original Bar
		Path     string
		SetVal   interface{}
		Expected Bar
	}{
		// Set slice in 4th level
		{
			Original: Bar{
				Name: "1", // name indicates level
				Bar: &Bar{
					Name: "2",
					Bars: []*Bar{
						{Name: "3-1", Bars: []*Bar{{Name: "4-1"}, {Name: "4-2"}, {Name: "4-3"}}},
						{Name: "3-2", Bars: []*Bar{{Name: "4-1"}, {Name: "4-2"}}},
					},
				},
			},
			Path:   "Bar.Bars.Bars.Name",
			SetVal: "val",
			Expected: Bar{
				Name: "1",
				Bar: &Bar{
					Name: "2",
					Bars: []*Bar{
						{Name: "3-1", Bars: []*Bar{{Name: "val"}, {Name: "val"}, {Name: "val"}}},
						{Name: "3-2", Bars: []*Bar{{Name: "val"}, {Name: "val"}}},
					},
				},
			},
		},
		// Set multilevel uninitialized ptr
		{
			Original: Bar{
				Name: "1", // name indicates level
				Bar:  nil,
			},
			Path:   "Bar.Bar.Bar.Name",
			SetVal: "val",
			Expected: Bar{
				Name: "1",
				Bar: &Bar{
					Name: "", // level 2
					Bar: &Bar{
						Bar: &Bar{
							Name: "val", //level 3
						},
					},
				},
			},
		},
		// mix of uninitialized ptr and slices
		{
			Original: Bar{
				Name: "1", // name indicates level
				Bar: &Bar{
					Name: "2",
					Bars: []*Bar{
						{Name: "3-1", Bars: []*Bar{{Name: "4-1"}, {Name: "4-2"}, {Name: "4-3"}}},
						{Name: "3-2", Bars: []*Bar{{Name: "4-1"}, {Name: "4-2"}}},
					},
				},
			},
			Path:   "Bar.Bars.Bars.Bar.Name",
			SetVal: "val",
			Expected: Bar{
				Name: "1", // name indicates level
				Bar: &Bar{
					Name: "2",
					Bars: []*Bar{
						{Name: "3-1", Bars: []*Bar{{Name: "4-1", Bar: &Bar{Name: "val"}},
							{Name: "4-2", Bar: &Bar{Name: "val"}}, {Name: "4-3", Bar: &Bar{Name: "val"}}}},
						{Name: "3-2", Bars: []*Bar{{Name: "4-1", Bar: &Bar{Name: "val"}}, {Name: "4-2", Bar: &Bar{Name: "val"}}}},
					},
				},
			},
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

func TestSet_StructWithCyclicStruct(t *testing.T) {
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

func TestSet_StructWithFieldNotInitialized(t *testing.T) {
	is := assert.New(t)
	myFoo := &Foo{
		Bar: nil, // we will try to set bar's field
	}
	err := Set(myFoo, "name", "Bar.Name")
	is.NoError(err)
	is.Equal("name", myFoo.Bar.Name)
}

func TestSet_SlicePassByPtr(t *testing.T) {

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

func TestSet_SlicePassDirectly(t *testing.T) {
	var testCases = []struct {
		Original interface{} // slice or array
		Path     string
		SetVal   interface{}
		Expected interface{}
	}{
		// Set Slice itself does not work here since not passing by ptr

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
		// Array of ptr
		{
			Original: [2]*Bar{{Name: "a"}, {Name: "b"}},
			Path:     "Name",
			SetVal:   "val",
			Expected: [2]*Bar{{Name: "val"}, {Name: "val"}},
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

	var testCases = []struct {
		OriginalFoo Foo
		Path        string
		SetVal      interface{}
		ExpectedFoo Foo
	}{
		// set string field
		{
			Foo{FirstName: ""},
			"FirstName",
			"hi",
			Foo{FirstName: "hi"},
		},
		// set interface{} field
		{
			Foo{FirstName: "", GeneralInterface: nil},
			"GeneralInterface",
			"str",
			Foo{FirstName: "", GeneralInterface: "str"},
		},
		// set field of the interface{} field
		// Note: set uninitialized interface{} should fail
		// Note: interface of struct (not ptr to struct) should fail
		{
			Foo{FirstName: "", GeneralInterface: &Foo{FirstName: ""}}, // if Foo is not ptr this will fail
			"GeneralInterface.FirstName",
			"foo",
			Foo{FirstName: "", GeneralInterface: &Foo{FirstName: "foo"}},
		},
		// interface two level
		{
			Foo{FirstName: "", GeneralInterface: &Foo{GeneralInterface: nil}},
			"GeneralInterface.GeneralInterface",
			"val",
			Foo{FirstName: "", GeneralInterface: &Foo{GeneralInterface: "val"}},
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			err := Set(&tc.OriginalFoo, tc.SetVal, tc.Path)
			is.NoError(err)
			is.Equal(tc.ExpectedFoo, tc.OriginalFoo)
		})
	}

}

func TestSet_ErrorCaces(t *testing.T) {

	var testCases = []struct {
		OriginalFoo Foo
		Path        string
		SetVal      interface{}
	}{
		// uninit interface
		// Itf is not initialized so Set cannot properly allocate type
		{
			Foo{BarInterface: nil},
			"BarInterface.Name",
			"val",
		},
		{
			Foo{GeneralInterface: &Foo{BarInterface: nil}},
			"GeneralInterface.BarInterface.Name",
			"val",
		},
		// type mismatch
		{
			Foo{FirstName: ""},
			"FirstName",
			20,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			err := Set(&tc.OriginalFoo, tc.SetVal, tc.Path)
			is.Error(err)
		})
	}

	t.Run("not pointer", func(t *testing.T) {
		is := assert.New(t)
		baz := Bar{Name: "dummy"}
		err := Set(baz, Bar{Name: "dummy2"}, "Name")
		is.Error(err)
	})

	t.Run("Unexported field", func(t *testing.T) {
		is := assert.New(t)
		s := struct {
			name string
		}{name: "dummy"}
		err := Set(&s, s, "name")
		is.Error(err)
	})
}

func TestMustSet_Basic(t *testing.T) {
	t.Run("Variable", func(t *testing.T) {
		is := assert.New(t)
		s := 1
		MustSet(&s, 2, "")
		is.Equal(2, s)
	})

	t.Run("Struct", func(t *testing.T) {
		is := assert.New(t)
		s := Bar{Name: "a"}
		MustSet(&s, "b", "Name")
		is.Equal("b", s.Name)
	})
}

// Examples

func ExampleSet() {

	var bar Bar = Bar{
		Name: "level-0",
		Bar: &Bar{
			Name: "level-1",
			Bars: []*Bar{
				{Name: "level2-1"},
				{Name: "level2-2"},
			},
		},
	}

	_ = Set(&bar, "level-0-new", "Name")
	fmt.Println(bar.Name)

	// discard error use MustSet
	MustSet(&bar, "level-1-new", "Bar.Name")
	fmt.Println(bar.Bar.Name)

	_ = Set(&bar, "level-2-new", "Bar.Bars.Name")
	fmt.Println(bar.Bar.Bars[0].Name)
	fmt.Println(bar.Bar.Bars[1].Name)

	// Output:
	// level-0-new
	// level-1-new
	// level-2-new
	// level-2-new
}
