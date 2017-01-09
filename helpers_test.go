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
		FirstName: "Dark",
		LastName:  "Vador",
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

func TestRandomInt(t *testing.T) {
	assert := assert.New(t)

	assert.True(RandomInt(0, 10) <= 10)
}

func TestShard(t *testing.T) {
	assert := assert.New(t)

	tokey := "e89d66bdfdd4dd26b682cc77e23a86eb"

	assert.Equal(Shard(tokey, 1, 2, false), []string{"e", "8", "e89d66bdfdd4dd26b682cc77e23a86eb"})
	assert.Equal(Shard(tokey, 2, 2, false), []string{"e8", "9d", "e89d66bdfdd4dd26b682cc77e23a86eb"})
	assert.Equal(Shard(tokey, 2, 3, true), []string{"e8", "9d", "66", "bdfdd4dd26b682cc77e23a86eb"})
}

func TestRandomString(t *testing.T) {
	assert := assert.New(t)

	assert.Len(RandomString(10), 10)

	result := RandomString(10, []rune("abcdefg"))

	assert.Len(result, 10)

	for _, char := range result {
		assert.True(char >= []rune("a")[0] && char <= []rune("g")[0])
	}
}
