package funk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForEach(t *testing.T) {
	assert := assert.New(t)

	results := []int{}

	ForEach([]int{1, 2, 3, 4}, func(x int) {
		if x%2 == 0 {
			results = append(results, x)
		}
	})

	assert.Equal(results, []int{2, 4})

	mapping := map[int]string{
		1: "Florent",
		2: "Gilles",
	}

	ForEach(mapping, func(k int, v string) {
		assert.Equal(v, mapping[k])
	})
}

func TestForEachRight(t *testing.T) {
	assert := assert.New(t)

	results := []int{}

	ForEachRight([]int{1, 2, 3, 4}, func(x int) {
		results = append(results, x*2)
	})

	assert.Equal(results, []int{8, 6, 4, 2})

	mapping := map[int]string{
		1: "Florent",
		2: "Gilles",
	}

	mapKeys := []int{}

	ForEachRight(mapping, func(k int, v string) {
		assert.Equal(v, mapping[k])

		mapKeys = append(mapKeys, k)
	})

	assert.Equal(mapKeys, []int{2, 1})
}
