package funk

import (
	"math/rand"
	"reflect"
	"time"
)

var numericZeros = []interface{}{
	int(0),
	int8(0),
	int16(0),
	int32(0),
	int64(0),
	uint(0),
	uint8(0),
	uint16(0),
	uint32(0),
	uint64(0),
	float32(0),
	float64(0),
}

// PtrOf makes a copy of the given interface and returns a pointer.
func PtrOf(itf interface{}) interface{} {
	t := reflect.TypeOf(itf)

	cp := reflect.New(t)
	cp.Elem().Set(reflect.ValueOf(itf))

	// Avoid double pointers if itf is a pointer
	if t.Kind() == reflect.Ptr {
		return cp.Elem().Interface()
	}

	return cp.Interface()
}

// IsFunction will return if the argument is a function.
func IsFunction(in interface{}, num ...int) bool {
	funcType := reflect.TypeOf(in)

	result := funcType.Kind() == reflect.Func

	if len(num) >= 1 {
		result = result && funcType.NumIn() == num[0]
	}

	if len(num) == 2 {
		result = result && funcType.NumOut() == num[1]
	}

	return result
}

// IsIteratee will return if the argument is an iteratee.
func IsIteratee(in interface{}) bool {
	arrType := reflect.TypeOf(in)

	kind := arrType.Kind()

	return kind == reflect.Array || kind == reflect.Slice || kind == reflect.Map
}

// SliceOf will return a slice which contains the element.
func SliceOf(in interface{}) interface{} {
	value := reflect.ValueOf(in)

	sliceType := reflect.SliceOf(reflect.TypeOf(in))
	slice := reflect.New(sliceType)
	sliceValue := reflect.MakeSlice(sliceType, 0, 0)
	sliceValue = reflect.Append(sliceValue, value)
	slice.Elem().Set(sliceValue)

	return slice.Elem().Interface()
}

// IsEmtpty will return if the object is considered as empty or not.
func IsEmpty(obj interface{}) bool {
	if obj == nil || obj == "" || obj == false {
		return true
	}

	for _, v := range numericZeros {
		if obj == v {
			return true
		}
	}

	objValue := reflect.ValueOf(obj)

	switch objValue.Kind() {
	case reflect.Map:
		fallthrough
	case reflect.Slice, reflect.Chan:
		return (objValue.Len() == 0)
	case reflect.Struct:
		return reflect.DeepEqual(obj, ZeroOf(obj))
	case reflect.Ptr:
		if objValue.IsNil() {
			return true
		}

		return reflect.DeepEqual(redirectValue(objValue).Interface(), ZeroOf(obj))
	}
	return false
}

// ZeroOf returns a zero value of an element.
func ZeroOf(in interface{}) interface{} {
	return reflect.Zero(reflect.TypeOf(in)).Interface()
}

// RandomInt generates a random int, based on a min and max values
func RandomInt(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

// Shard will shard a string name
func Shard(str string, width int, depth int, restOnly bool) []string {
	var results []string

	for i := 0; i < depth; i++ {
		results = append(results, str[(width*i):(width*(i+1))])
	}

	if restOnly {
		results = append(results, str[(width*depth):])
	} else {
		results = append(results, str)
	}

	return results
}

var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandomString returns a random string with a fixed length
func RandomString(n int, allowedChars ...[]rune) string {
	var letters []rune

	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
