package funk

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	assert := assert.New(t)

	results := Keys(map[string]int{"one": 1, "two": 2}).([]string)
	sort.Strings(results)

	assert.Equal(results, []string{"one", "two"})

	fields := Keys(foo).([]string)

	sort.Strings(fields)

	assert.Equal(fields, []string{"Age", "Bar", "Bars", "EmptyValue", "FirstName", "ID", "LastName"})
}

func TestValues(t *testing.T) {
	assert := assert.New(t)

	results := Values(map[string]int{"one": 1, "two": 2}).([]int)
	sort.Ints(results)

	assert.Equal(results, []int{1, 2})

	values := Values(foo).([]interface{})

	assert.Len(values, 7)
}
