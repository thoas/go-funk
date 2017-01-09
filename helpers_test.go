package funk

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceOf(t *testing.T) {
	assert := assert.New(t)

	f := &Foo{
		ID:        1,
		FirstName: "Drew",
		LastName:  "Olson",
		Age:       30,
		Bar: &Bar{
			Name: "Test",
		},
	}

	result := SliceOf(f)

	resultType := reflect.TypeOf(result)

	assert.True(resultType.Kind() == reflect.Slice)
	assert.True(resultType.Elem().Kind() == reflect.Ptr)

	elemType := resultType.Elem().Elem()

	assert.True(elemType.Kind() == reflect.Struct)

	value := reflect.ValueOf(result)

	assert.Equal(value.Len(), 1)

	_, ok := value.Index(0).Interface().(*Foo)

	assert.True(ok)
}
