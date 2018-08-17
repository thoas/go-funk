package funk

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	Test à faire:

	ForEach(predicate interface{}) Builder
	ForEachRight(predicate interface{}) Builder
	Initial() Builder

	All() bool
	Any() bool
	Get(path string) interface{}
	Head() interface{}
	Keys() interface{}
	In(v interface{}) bool
	IsEmpty() bool
	IsType(actual interface{}) bool
	Last() interface{}
	NotEmpty() bool
	Product() float64
	Reduce(reduceFunc, acc interface{}) float64
	Sum() float64
	Tail() interface{}
	Type() reflect.Type
	Value() interface{}
	Values() interface{}
*/

func TestChainChunk(t *testing.T) {
	testCases := []struct {
		In   interface{}
		Size int
	}{
		{
			In:   []int{0, 1, 2, 3, 4},
			Size: 2,
		},
		{
			In:   []int{},
			Size: 2,
		},
		{
			In:   []int{1},
			Size: 2,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			expected := Chunk(tc.In, tc.Size)
			actual := Chain(tc.In).Chunk(tc.Size).Value()

			is.Equal(expected, actual)
		})
	}
}

func TestChainCompact(t *testing.T) {
	var emptyFunc func() bool
	emptyFuncPtr := &emptyFunc

	nonEmptyFunc := func() bool { return true }
	nonEmptyFuncPtr := &nonEmptyFunc

	nonEmptyMap := map[int]int{1: 2}
	nonEmptyMapPtr := &nonEmptyMap

	var emptyMap map[int]int
	emptyMapPtr := &emptyMap

	var emptyChan chan bool
	nonEmptyChan := make(chan bool, 1)
	nonEmptyChan <- true

	emptyChanPtr := &emptyChan
	nonEmptyChanPtr := &nonEmptyChan

	var emptyString string
	emptyStringPtr := &emptyString

	nonEmptyString := "42"
	nonEmptyStringPtr := &nonEmptyString

	testCases := []struct {
		In interface{}
	}{
		// Check with nils
		{
			In: []interface{}{42, nil, (*int)(nil)},
		},

		// Check with functions
		{
			In: []interface{}{42, emptyFuncPtr, emptyFunc, nonEmptyFuncPtr},
		},

		// Check with slices, maps, arrays and channels
		{
			In: []interface{}{
				42, [2]int{}, map[int]int{}, []string{}, nonEmptyMapPtr, emptyMap,
				emptyMapPtr, nonEmptyMap, nonEmptyChan, emptyChan, emptyChanPtr, nonEmptyChanPtr,
			},
		},

		// Check with strings, numbers and booleans
		{
			In: []interface{}{true, 0, float64(0), "", "42", emptyStringPtr, nonEmptyStringPtr, false},
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			expected := Compact(tc.In)
			actual := Chain(tc.In).Compact().Value()

			is.Equal(expected, actual)
		})
	}
}

func TestChainDrop(t *testing.T) {
	testCases := []struct {
		In interface{}
		N  int
	}{
		{
			In: []int{0, 1, 1, 2, 3, 0, 0, 12},
			N:  3,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			expected := Drop(tc.In, tc.N)
			actual := Chain(tc.In).Drop(tc.N).Value()

			is.Equal(expected, actual)
		})
	}
}

func TestChainFlattenDeep(t *testing.T) {
	testCases := []struct {
		In interface{}
	}{
		{
			In: [][]int{{1, 2}, {3, 4}},
		},
		{
			In: [][][]int{{{1, 2}, {3, 4}}, {{5, 6}}},
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			expected := FlattenDeep(tc.In)
			actual := Chain(tc.In).FlattenDeep().Value()

			is.Equal(expected, actual)
		})
	}
}

func TestChainFilter(t *testing.T) {
	testCases := []struct {
		In        interface{}
		Predicate interface{}
	}{
		{
			In:        []int{1, 2, 3, 4},
			Predicate: func(x int) bool { return x%2 == 0 },
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			expected := Filter(tc.In, tc.Predicate)
			actual := Chain(tc.In).Filter(tc.Predicate).Value()

			is.Equal(expected, actual)
		})
	}
}

func TestChainIntersect(t *testing.T) {
	testCases := []struct {
		In  interface{}
		Sec interface{}
	}{
		{
			In:  []int{1, 2, 3, 4},
			Sec: []int{2, 4, 6},
		},
		{
			In:  []string{"foo", "bar", "hello", "bar"},
			Sec: []string{"foo", "bar"},
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			expected := Intersect(tc.In, tc.Sec)
			actual := Chain(tc.In).Intersect(tc.Sec).Value()

			is.Equal(expected, actual)
		})
	}
}

func TestChainMap(t *testing.T) {
	testCases := []struct {
		In     interface{}
		MapFnc interface{}
	}{
		{
			In:     []int{1, 2, 3, 4},
			MapFnc: func(x int) string { return "Hello" },
		},
		{
			In:     []int{1, 2, 3, 4},
			MapFnc: func(x int) (int, int) { return x, x },
		},
		{
			In:     map[int]string{1: "Florent", 2: "Gilles"},
			MapFnc: func(k int, v string) int { return k },
		},
		{
			In:     map[int]string{1: "Florent", 2: "Gilles"},
			MapFnc: func(k int, v string) (string, string) { return fmt.Sprintf("%d", k), v },
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			expected := Map(tc.In, tc.MapFnc)
			actual := Chain(tc.In).Map(tc.MapFnc).Value()

			is.Equal(expected, actual)
		})
	}
}

func TestChainReverse(t *testing.T) {
	testCases := []struct {
		In interface{}
	}{
		{
			In: []int{0, 1, 2, 3, 4},
		},
		{
			In: "abcdefg",
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			expected := Reverse(tc.In)
			actual := Chain(tc.In).Reverse().Value()

			is.Equal(expected, actual)
		})
	}
}

func TestChainShuffle(t *testing.T) {
	testCases := []struct {
		In interface{}
	}{
		{
			In: []int{0, 1, 2, 3, 4},
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			expected := Shuffle(tc.In)
			actual := Chain(tc.In).Shuffle().Value()

			is.NotEqual(expected, actual)
			is.Len(actual, len(expected.([]int)))
			is.Contains(actual, expected.([]int)[0])
		})
	}
}

func TestChainUniq(t *testing.T) {
	testCases := []struct {
		In interface{}
	}{
		{
			In: []int{0, 1, 1, 2, 3, 0, 0, 12},
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			expected := Uniq(tc.In)
			actual := Chain(tc.In).Uniq().Value()

			is.Equal(expected, actual)
		})
	}
}

func TestChainContains(t *testing.T) {
	testCases := []struct {
		In       interface{}
		Contains interface{}
	}{
		{
			In:       []string{"foo", "bar"},
			Contains: "bar",
		},
		{
			In:       results,
			Contains: f,
		},
		{
			In:       results,
			Contains: nil,
		},
		{
			In:       results,
			Contains: b,
		},
		{
			In:       "florent",
			Contains: "rent",
		},
		{
			In:       "florent",
			Contains: "gilles",
		},
		{
			In:       map[int]*Foo{1: f, 3: c},
			Contains: 1,
		},
		{
			In:       map[int]*Foo{1: f, 3: c},
			Contains: 2,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			expected := Contains(tc.In, tc.Contains)
			actual := Chain(tc.In).Contains(tc.Contains)

			is.Equal(expected, actual)
		})
	}
}

func TestChainEvery(t *testing.T) {
	testCases := []struct {
		In       interface{}
		Contains []interface{}
	}{
		{
			In:       []string{"foo", "bar", "baz"},
			Contains: []interface{}{"bar", "foo"},
		},
		{
			In:       results,
			Contains: []interface{}{f, c},
		},
		{
			In:       results,
			Contains: []interface{}{nil},
		},
		{
			In:       results,
			Contains: []interface{}{f, b},
		},
		{
			In:       "florent",
			Contains: []interface{}{"rent", "flo"},
		},
		{
			In:       "florent",
			Contains: []interface{}{"rent", "gilles"},
		},
		{
			In:       map[int]*Foo{1: f, 3: c},
			Contains: []interface{}{1, 3},
		},
		{
			In:       map[int]*Foo{1: f, 3: c},
			Contains: []interface{}{2, 3},
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			expected := Every(tc.In, tc.Contains...)
			actual := Chain(tc.In).Every(tc.Contains...)

			is.Equal(expected, actual)
		})
	}
}

func TestChainIndexOf(t *testing.T) {
	testCases := []struct {
		In   interface{}
		Item interface{}
	}{
		{
			In:   []string{"foo", "bar"},
			Item: "bar",
		},
		{
			In:   results,
			Item: f,
		},
		{
			In:   results,
			Item: b,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			expected := IndexOf(tc.In, tc.Item)
			actual := Chain(tc.In).IndexOf(tc.Item)

			is.Equal(expected, actual)
		})
	}
}

func TestChainLastIndexOf(t *testing.T) {
	testCases := []struct {
		In   interface{}
		Item interface{}
	}{
		{
			In:   []string{"foo", "bar", "bar"},
			Item: "bar",
		},
		{
			In:   []int{1, 2, 2, 3},
			Item: 2,
		},
		{
			In:   []int{1, 2, 2, 3},
			Item: 4,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			expected := LastIndexOf(tc.In, tc.Item)
			actual := Chain(tc.In).LastIndexOf(tc.Item)

			is.Equal(expected, actual)
		})
	}
}

func TestChainFind(t *testing.T) {
	testCases := []struct {
		In        interface{}
		Predicate interface{}
	}{
		{
			In:        []int{1, 2, 3, 4},
			Predicate: func(x int) bool { return x%2 == 0 },
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)

			expected := Find(tc.In, tc.Predicate)
			actual := Chain(tc.In).Find(tc.Predicate)

			is.Equal(expected, actual)
		})
	}
}
