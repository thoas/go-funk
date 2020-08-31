package funk

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	is := assert.New(t)

	r := Map([]int{1, 2, 3, 4}, func(x int) string {
		return "Hello"
	})

	result, ok := r.([]string)

	is.True(ok)
	is.Equal(len(result), 4)

	r = Map([]int{1, 2, 3, 4}, func(x int) (int, int) {
		return x, x
	})

	resultType := reflect.TypeOf(r)

	is.True(resultType.Kind() == reflect.Map)
	is.True(resultType.Key().Kind() == reflect.Int)
	is.True(resultType.Elem().Kind() == reflect.Int)

	mapping := map[int]string{
		1: "Florent",
		2: "Gilles",
	}

	r = Map(mapping, func(k int, v string) int {
		return k
	})

	is.True(reflect.TypeOf(r).Kind() == reflect.Slice)
	is.True(reflect.TypeOf(r).Elem().Kind() == reflect.Int)

	r = Map(mapping, func(k int, v string) (string, string) {
		return fmt.Sprintf("%d", k), v
	})

	resultType = reflect.TypeOf(r)

	is.True(resultType.Kind() == reflect.Map)
	is.True(resultType.Key().Kind() == reflect.String)
	is.True(resultType.Elem().Kind() == reflect.String)
}

func TestToMap(t *testing.T) {
	is := assert.New(t)

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

	is.True(reflect.TypeOf(instanceMap).Kind() == reflect.Map)

	mapping, ok := instanceMap.(map[int]*Foo)

	is.True(ok)

	for _, result := range results {
		item, ok := mapping[result.ID]

		is.True(ok)
		is.True(reflect.TypeOf(item).Kind() == reflect.Ptr)
		is.True(reflect.TypeOf(item).Elem().Kind() == reflect.Struct)

		is.Equal(item.ID, result.ID)
	}
}

func TestChunk(t *testing.T) {
	is := assert.New(t)

	results := Chunk([]int{0, 1, 2, 3, 4}, 2).([][]int)

	is.Len(results, 3)
	is.Len(results[0], 2)
	is.Len(results[1], 2)
	is.Len(results[2], 1)

	is.Len(Chunk([]int{}, 2), 0)
	is.Len(Chunk([]int{1}, 2), 1)
	is.Len(Chunk([]int{1, 2, 3}, 0), 3)
}

func TestFlattenDeep(t *testing.T) {
	is := assert.New(t)

	is.Equal(FlattenDeep([][]int{{1, 2}, {3, 4}}), []int{1, 2, 3, 4})
}

func TestShuffle(t *testing.T) {
	initial := []int{0, 1, 2, 3, 4}

	results := Shuffle(initial)

	is := assert.New(t)

	is.Len(results, 5)

	for _, entry := range initial {
		is.True(Contains(results, entry))
	}
}

func TestReverse(t *testing.T) {
	results := Reverse([]int{0, 1, 2, 3, 4})

	is := assert.New(t)

	is.Equal(Reverse("abcdefg"), "gfedcba")
	is.Len(results, 5)

	is.Equal(results, []int{4, 3, 2, 1, 0})
}

func TestUniq(t *testing.T) {
	is := assert.New(t)

	results := Uniq([]int{0, 1, 1, 2, 3, 0, 0, 12})
	is.Len(results, 5)
	is.Equal(results, []int{0, 1, 2, 3, 12})

	results = Uniq([]string{"foo", "bar", "foo", "bar", "bar"})
	is.Len(results, 2)
	is.Equal(results, []string{"foo", "bar"})
}

func TestConvertSlice(t *testing.T) {
	instances := []*Foo{foo, foo2}

	var raw []Model

	ConvertSlice(instances, &raw)

	is := assert.New(t)

	is.Len(raw, len(instances))
}

func TestDrop(t *testing.T) {
	results := Drop([]int{0, 1, 1, 2, 3, 0, 0, 12}, 3)

	is := assert.New(t)

	is.Len(results, 5)

	is.Equal([]int{2, 3, 0, 0, 12}, results)
}

func TestPrune(t *testing.T) {

	var testCases = []struct {
		OriginalFoo *Foo
		Paths       []string
		ExpectedFoo *Foo
	}{
		{
			foo,
			[]string{"FirstName"},
			&Foo{
				FirstName: foo.FirstName,
			},
		},
		{
			foo,
			[]string{"FirstName", "ID"},
			&Foo{
				FirstName: foo.FirstName,
				ID:        foo.ID,
			},
		},
		{
			foo,
			[]string{"EmptyValue.Int64"},
			&Foo{
				EmptyValue: sql.NullInt64{
					Int64: foo.EmptyValue.Int64,
				},
			},
		},
		{
			foo,
			[]string{"FirstName", "ID", "EmptyValue.Int64"},
			&Foo{
				FirstName: foo.FirstName,
				ID:        foo.ID,
				EmptyValue: sql.NullInt64{
					Int64: foo.EmptyValue.Int64,
				},
			},
		},
		{
			foo,
			[]string{"FirstName", "ID", "EmptyValue.Int64"},
			&Foo{
				FirstName: foo.FirstName,
				ID:        foo.ID,
				EmptyValue: sql.NullInt64{
					Int64: foo.EmptyValue.Int64,
				},
			},
		},
		{
			foo,
			[]string{"FirstName", "ID", "Bar"},
			&Foo{
				FirstName: foo.FirstName,
				ID:        foo.ID,
				Bar:       foo.Bar,
			},
		},
		{
			foo,
			[]string{"Bar", "Bars"},
			&Foo{
				Bar:  foo.Bar,
				Bars: foo.Bars,
			},
		},
		{
			foo,
			[]string{"FirstName", "Bars.Name"},
			&Foo{
				FirstName: foo.FirstName,
				Bars: []*Bar{
					{Name: bar.Name},
					{Name: bar.Name},
				},
			},
		},
		{
			foo,
			[]string{"Bars.Name", "Bars.Bars.Name"},
			&Foo{
				Bars: []*Bar{
					{Name: bar.Name, Bars: []*Bar{{Name: "Level1-1"}, {Name: "Level1-2"}}},
					{Name: bar.Name, Bars: []*Bar{{Name: "Level1-1"}, {Name: "Level1-2"}}},
				},
			},
		},
		{
			foo,
			[]string{"BarInterface", "BarPointer"},
			&Foo{
				BarInterface: bar,
				BarPointer:   &bar,
			},
		},
	}

	// pass to prune by pointer to struct
	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Prune pointer test case #%v", idx), func(t *testing.T) {
			is := assert.New(t)
			res, err := Prune(tc.OriginalFoo, tc.Paths)
			require.NoError(t, err)

			fooPrune := res.(*Foo)
			is.Equal(tc.ExpectedFoo, fooPrune)
		})
	}

	// pass to prune by struct directly
	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Prune non pointer test case #%v", idx), func(t *testing.T) {
			is := assert.New(t)
			fooNonPtr := *tc.OriginalFoo
			res, err := Prune(fooNonPtr, tc.Paths)
			require.NoError(t, err)

			fooPrune := res.(Foo)
			is.Equal(*tc.ExpectedFoo, fooPrune)
		})
	}

	// test PruneByTag
	var TagTestCases = []struct {
		OriginalFoo *Foo
		Paths       []string
		ExpectedFoo *Foo
		Tag         string
	}{
		{
			foo,
			[]string{"tag 1", "tag 4.BarName"},
			&Foo{
				FirstName: foo.FirstName,
				Bar: &Bar{
					Name: bar.Name,
				},
			},
			"tag_name",
		},
	}

	for idx, tc := range TagTestCases {
		t.Run(fmt.Sprintf("PruneByTag test case #%v", idx), func(t *testing.T) {
			is := assert.New(t)
			fooNonPtr := *tc.OriginalFoo
			res, err := PruneByTag(fooNonPtr, tc.Paths, tc.Tag)
			require.NoError(t, err)

			fooPrune := res.(Foo)
			is.Equal(*tc.ExpectedFoo, fooPrune)
		})
	}

	t.Run("Bar Slice", func(t *testing.T) {
		barSlice := []*Bar{bar, bar}
		barSlicePruned, err := pruneByTag(barSlice, []string{"Name"}, nil /*tag*/)
		require.NoError(t, err)
		assert.Equal(t, []*Bar{{Name: bar.Name}, {Name: bar.Name}}, barSlicePruned)
	})

	t.Run("Bar Array", func(t *testing.T) {
		barArr := [2]*Bar{bar, bar}
		barArrPruned, err := pruneByTag(barArr, []string{"Name"}, nil /*tag*/)
		require.NoError(t, err)
		assert.Equal(t, [2]*Bar{{Name: bar.Name}, {Name: bar.Name}}, barArrPruned)
	})

	// test values are copied and not referenced in return result
	// NOTE: pointers at the end of path are referenced. Maybe we need to make a copy
	t.Run("Copy Value Str", func(t *testing.T) {
		is := assert.New(t)
		fooTest := &Foo{
			Bar: &Bar{
				Name: "bar",
			},
		}
		res, err := pruneByTag(fooTest, []string{"Bar.Name"}, nil)
		require.NoError(t, err)
		fooTestPruned := res.(*Foo)
		is.Equal(fooTest, fooTestPruned)

		// change pruned
		fooTestPruned.Bar.Name = "changed bar"
		// check original is unchanged
		is.Equal(fooTest.Bar.Name, "bar")
	})

	// error cases
	var errCases = []struct {
		InputFoo *Foo
		Paths    []string
		TagName  *string
	}{
		{
			foo,
			[]string{"NotExist"},
			nil,
		},
		{
			foo,
			[]string{"FirstName.NotExist", "LastName"},
			nil,
		},
		{
			foo,
			[]string{"LastName", "FirstName.NotExist"},
			nil,
		},
		{
			foo,
			[]string{"LastName", "Bars.NotExist"},
			nil,
		},
		// tags
		{
			foo,
			[]string{"tag 999"},
			&[]string{"tag_name"}[0],
		},
		{
			foo,
			[]string{"tag 1.NotExist"},
			&[]string{"tag_name"}[0],
		},
		{
			foo,
			[]string{"tag 4.NotExist"},
			&[]string{"tag_name"}[0],
		},
		{
			foo,
			[]string{"FirstName"},
			&[]string{"tag_name_not_exist"}[0],
		},
	}

	for idx, errTC := range errCases {
		t.Run(fmt.Sprintf("error test case #%v", idx), func(t *testing.T) {
			_, err := pruneByTag(errTC.InputFoo, errTC.Paths, errTC.TagName)
			assert.Error(t, err)
		})
	}
}

func ExamplePrune() {
	type ExampleFoo struct {
		ExampleFooPtr *ExampleFoo `json:"example_foo_ptr"`
		Name          string      `json:"name"`
		Number        int         `json:"number"`
	}

	exampleFoo := ExampleFoo{
		ExampleFooPtr: &ExampleFoo{
			Name:   "ExampleFooPtr",
			Number: 2,
		},
		Name:   "ExampleFoo",
		Number: 1,
	}

	// prune using struct field name
	res, _ := Prune(exampleFoo, []string{"ExampleFooPtr.Name", "Number"})
	prunedFoo := res.(ExampleFoo)
	fmt.Println(prunedFoo.ExampleFooPtr.Name)
	fmt.Println(prunedFoo.ExampleFooPtr.Number)
	fmt.Println(prunedFoo.Name)
	fmt.Println(prunedFoo.Number)

	// prune using struct json tag
	res2, _ := PruneByTag(exampleFoo, []string{"example_foo_ptr.name", "number"}, "json")
	prunedByTagFoo := res2.(ExampleFoo)
	fmt.Println(prunedByTagFoo.ExampleFooPtr.Name)
	fmt.Println(prunedByTagFoo.ExampleFooPtr.Number)
	fmt.Println(prunedByTagFoo.Name)
	fmt.Println(prunedByTagFoo.Number)
	// output:
	// ExampleFooPtr
	// 0
	//
	// 1
	// ExampleFooPtr
	// 0
	//
	// 1
}
