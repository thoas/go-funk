package funk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	is := assert.New(t)

	is.Equal(Sum([]int{1, 2, 3}), 6.0)
	is.Equal(Sum(&[]int{1, 2, 3}), 6.0)
	is.Equal(Sum([]interface{}{1, 2, 3, 0.5}), 6.5)
}

func TestSumBy(t *testing.T) {
	type NumItem struct {
		Num float64
	}
	is := assert.New(t)
	var numArray []interface{}
	numArray = append(numArray, NumItem{Num: 1})
	numArray = append(numArray, NumItem{Num: 2})
	numArray = append(numArray, NumItem{Num: 3})

	is.Equal(SumBy(numArray, "Num"), 6.0)
	numArray = append(numArray, NumItem{Num: 4.2})
	is.Equal(SumBy(numArray, "Num"), 10.2)
}

func TestProduct(t *testing.T) {
	is := assert.New(t)

	is.Equal(Product([]int{2, 3, 4}), 24.0)
	is.Equal(Product(&[]int{2, 3, 4}), 24.0)
	is.Equal(Product([]interface{}{2, 3, 4, 0.5}), 12.0)
}
