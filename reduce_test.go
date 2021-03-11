package funk

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	testCases := []struct {
		Arr    interface{}
		Func   interface{}
		Acc    interface{}
		Result interface{}
	}{
		{
			[]int{1, 2, 3, 4},
			func(acc, elem float64) float64 { return acc + elem },
			0,
			float64(10),
		},
		{
			&[]int16{1, 2, 3, 4},
			'+',
			5,
			int16(15),
		},
		{
			[]float64{1.1, 2.2, 3.3},
			'+',
			0,
			float64(6.6),
		},
		{
			&[]int{1, 2, 3, 5},
			func(acc int8, elem int16) int32 { return int32(acc) * int32(elem) },
			1,
			int32(30),
		},
		{
			[]interface{}{1, 2, 3.3, 4},
			'*',
			1,
			float64(26.4),
		},
		{
			[]string{"1", "2", "3", "4"},
			func(acc string, elem string) string { return acc + elem },
			"",
			"1234",
		},
	}

	for idx, test := range testCases {
		t.Run(fmt.Sprintf("test case #%d", idx+1), func(t *testing.T) {
			is := assert.New(t)
			result := Reduce(test.Arr, test.Func, test.Acc)
			if !is.Equal(result, test.Result) {
				t.Errorf("%#v doesn't eqal to %#v", result, test.Result)
			}
		})
	}
}
