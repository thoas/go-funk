package funk

import (
	"reflect"
)

// Filter iterates over elements of collection, returning an array of
// all elements predicate returns truthy for.
func Filter[T any](arr []T, predicate func(T) bool) []T {
	res := make([]T, 0, len(arr))
	for _, e := range arr {
		if predicate(e) {
			res = append(res, e)
		}
	}
	return res
}

// Find iterates over elements of collection, returning the first
// element predicate returns truthy for and true as a second argument.
// If no elements found - nil value for given type will be returned with false.
func Find[T any](arr []T, predicate func(T) bool) (def T, ok bool) {
	for _, e := range arr {
		if predicate(e) {
			return e, true
		}
	}
	return def, ok
}

// todo: should only work with maps?
// FindKey iterates over elements of collection, returning the first
// element of an array and random of a map which predicate returns truthy for.
func FindKey(arr interface{}, predicate interface{}) (matchKey, matchEle interface{}) {
	if !IsIteratee(arr) {
		panic("First parameter must be an iteratee")
	}

	if !IsFunction(predicate, 1, 1) {
		panic("Second argument must be function")
	}

	funcValue := reflect.ValueOf(predicate)

	funcType := funcValue.Type()

	if funcType.Out(0).Kind() != reflect.Bool {
		panic("Return argument should be a boolean")
	}

	arrValue := reflect.ValueOf(arr)
	var keyArrs []reflect.Value

	isMap := arrValue.Kind() == reflect.Map
	if isMap {
		keyArrs = arrValue.MapKeys()
	}
	for i := 0; i < arrValue.Len(); i++ {
		var (
			elem reflect.Value
			key  reflect.Value
		)
		if isMap {
			key = keyArrs[i]
			elem = arrValue.MapIndex(key)
		} else {
			key = reflect.ValueOf(i)
			elem = arrValue.Index(i)
		}

		result := funcValue.Call([]reflect.Value{elem})[0].Interface().(bool)

		if result {
			return key.Interface(), elem.Interface()
		}
	}

	return nil, nil
}

// todo: to use callback as a second arg - need to write another method: IndexOfBy
// IndexOf gets the index at which the first occurrence of value is found in array or return -1
// if the value cannot be found
func IndexOf[T comparable](in []T, elem T) int {
	for i, e := range in {
		if e == elem {
			return i
		}
	}
	return -1
}

// todo: to use callback as a second arg - need to write another method: LastIndexOfBy
// LastIndexOf gets the index at which the last occurrence of value is found in array or return -1
// if the value cannot be found
func LastIndexOf[T comparable](in []T, elem T) int {
	for i := len(in) - 1; i >= 0; i-- {
		if in[i] == elem {
			return i
		}
	}
	return -1
}

// Contains returns true if an element is present in a iteratee.
func Contains[T comparable](in []T, elem T) bool {
	for _, e := range in {
		if e == elem {
			return true
		}
	}
	return false
}

// Every returns true if every element is present in an array.
func Every[T comparable](in []T, elements ...T) bool {
	for _, elem := range elements {
		if !Contains(in, elem) {
			return false
		}
	}
	return true
}

// Some returns true if at least one element is present in an array.
func Some[T comparable](in []T, elements ...T) bool {
	for _, elem := range elements {
		if Contains(in, elem) {
			return true
		}
	}
	return false
}
