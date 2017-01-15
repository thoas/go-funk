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
