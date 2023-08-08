package funk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partition_valid_int(t *testing.T) {
	result := Partition([]int{1, 2, 3, 4}, func(n int) bool {
		return n%2 == 0
	})

	assert.Equal(t, [][]int{{2, 4}, {1, 3}}, result)
}

func Test_partition_valid_float64(t *testing.T) {
	result := Partition([]float64{1.1, 2.2, 3.3, 4.4}, func(n float64) bool {
		return n > float64(2)
	})

	assert.Equal(t, [][]float64{{2.2, 3.3, 4.4}, {1.1}}, result)
}

func Test_partition_valid_string(t *testing.T) {
	result := Partition([]string{"a", "b", "c"}, func(n string) bool {
		return n > "a"
	})

	assert.Equal(t, [][]string{{"b", "c"}, {"a"}}, result)
}

func Test_partition_valid_struct(t *testing.T) {
	result := Partition([]*Foo{
		{
			FirstName: "Kakalot",
			Age:       26,
		},
		{
			FirstName: "Vegeta",
			Age:       27,
		},
		{
			FirstName: "Trunk",
			Age:       10,
		},
	}, func(n *Foo) bool {
		return n.Age%2 == 0
	})

	assert.Equal(t, [][]*Foo{{{FirstName: "Kakalot", Age: 26}, {FirstName: "Trunk", Age: 10}}, {{FirstName: "Vegeta", Age: 27}}}, result)
}
