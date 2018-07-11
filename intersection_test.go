package funk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	is := assert.New(t)

	r := Intersect([]int{1, 2, 3, 4}, []int{2, 4, 6})
	is.Equal(r, []int{2, 4})

	r = Intersect([]string{"foo", "bar", "hello", "bar"}, []string{"foo", "bar"})
	is.Equal(r, []string{"foo", "bar"})

}

func TestIntersectString(t *testing.T) {
	is := assert.New(t)

	r := IntersectString([]string{"foo", "bar", "hello", "bar"}, []string{"foo", "bar"})
	is.Equal(r, []string{"foo", "bar"})

}
