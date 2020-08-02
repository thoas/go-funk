package funk

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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
		t.Run(fmt.Sprintf("pointer test case #%v", idx), func(t *testing.T) {
			is := assert.New(t)
			res, err := Prune(tc.OriginalFoo, tc.Paths)
			is.NoError(err)

			fooPrune := res.(*Foo)
			is.Equal(tc.ExpectedFoo, fooPrune)
		})
	}

	// pass to prune by struct directly
	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("non pointer test case #%v", idx), func(t *testing.T) {
			is := assert.New(t)
			fooNonPtr := *tc.OriginalFoo
			res, err := Prune(fooNonPtr, tc.Paths)
			is.NoError(err)

			fooPrune := res.(Foo)
			is.Equal(*tc.ExpectedFoo, fooPrune)
		})
	}

	t.Run("Bar Slice", func(t *testing.T) {
		is := assert.New(t)
		barSlice := []*Bar{bar, bar}
		barSlicePruned, err := Prune(barSlice, []string{"Name"})
		is.NoError(err)
		is.Equal([]*Bar{{Name: bar.Name}, {Name: bar.Name}}, barSlicePruned)
	})

	t.Run("Bar Array", func(t *testing.T) {
		is := assert.New(t)
		barArr := [2]*Bar{bar, bar}
		barArrPruned, err := Prune(barArr, []string{"Name"})
		is.NoError(err)
		is.Equal([2]*Bar{{Name: bar.Name}, {Name: bar.Name}}, barArrPruned)
	})

	t.Run("Copy Value", func(t *testing.T) {
		is := assert.New(t)
		fooTest := &Foo{
			Bar: &Bar{
				Name: "bar",
			},
		}
		res, err := Prune(fooTest, []string{"Bar.Name"})
		is.NoError(err)
		fooTestPruned := res.(*Foo)
		is.Equal(fooTest, fooTestPruned)

		// change pruned
		fooTestPruned.Bar.Name = "changed bar"
		// check original is unchanged
		is.Equal(fooTest.Bar.Name, "bar")
	})
}
