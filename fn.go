package fn

import (
	"fmt"
	"reflect"
	"strings"
)

func equal(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	return reflect.DeepEqual(expected, actual)

}

// Chunk is ...
func Chunk(arr interface{}, size int) interface{} {
	return nil
}

// ForEach is ...
func ForEach(arr interface{}, mapFunc interface{}) {
	if !IsIterable(arr) {
		panic("First parameter must be neither array nor slice")
	}

	var (
		funcValue = reflect.ValueOf(mapFunc)
		arrValue  = reflect.ValueOf(arr)
		arrType   = arrValue.Type()
		funcType  = funcValue.Type()
	)

	if arrType.Kind() == reflect.Slice || arrType.Kind() == reflect.Array {
		if !IsFunction(mapFunc, 1, 0) {
			panic("Second argument must be a function with one parameter")
		}

		arrElemType := arrValue.Type().Elem()

		// Checking whether element type is convertible to function's first argument's type.
		if !arrElemType.ConvertibleTo(funcType.In(0)) {
			panic("Map function's argument is not compatible with type of array.")
		}

		for i := 0; i < arrValue.Len(); i++ {
			funcValue.Call([]reflect.Value{arrValue.Index(i)})
		}
	}

	if arrType.Kind() == reflect.Map {
		if !IsFunction(mapFunc, 2, 0) {
			panic("Second argument must be a function with two parameters")
		}

		// Type checking for Map<key, value> = (key, value)
		keyType := arrType.Key()
		valueType := arrType.Elem()

		if !keyType.ConvertibleTo(funcType.In(0)) {
			panic(fmt.Sprintf("function first argument is not compatible with %s", keyType.String()))
		}

		if !valueType.ConvertibleTo(funcType.In(1)) {
			panic(fmt.Sprintf("function second argument is not compatible with %s", valueType.String()))
		}

		for _, key := range arrValue.MapKeys() {
			funcValue.Call([]reflect.Value{key, arrValue.MapIndex(key)})
		}
	}
}

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

func IsIterable(in interface{}) bool {
	arrType := reflect.TypeOf(in)

	kind := arrType.Kind()

	return kind == reflect.Array || kind == reflect.Slice || kind == reflect.Map
}

// Filter is ...
func Filter(arr interface{}, mapFunc interface{}) interface{} {
	if !IsIterable(arr) {
		panic("First parameter must be neither array nor slice")
	}

	if !IsFunction(mapFunc, 1, 1) {
		panic("Second argument must be function")
	}

	funcValue := reflect.ValueOf(mapFunc)

	funcType := funcValue.Type()

	if funcType.Out(0).Kind() != reflect.Bool {
		panic("Return argument should be a boolean")
	}

	arrValue := reflect.ValueOf(arr)

	arrType := arrValue.Type()

	// Get slice type corresponding to array type
	resultSliceType := reflect.SliceOf(arrType.Elem())

	// MakeSlice takes a slice kind type, and makes a slice.
	resultSlice := reflect.MakeSlice(resultSliceType, 0, 0)

	for i := 0; i < arrValue.Len(); i++ {
		elem := arrValue.Index(i)

		result := funcValue.Call([]reflect.Value{elem})[0].Interface().(bool)

		if result == true {
			resultSlice = reflect.Append(resultSlice, elem)
		}
	}

	// Convering resulting slice back to generic interface.
	return resultSlice.Interface()
}

// Find is ...
func Find(arr interface{}, mapFunc interface{}) interface{} {
	return nil
}

// Contains is ...
func Contains(in interface{}, elem interface{}) bool {
	inValue := reflect.ValueOf(in)

	elemValue := reflect.ValueOf(elem)

	inType := inValue.Type()

	if inType.Kind() == reflect.String {
		return strings.Contains(inValue.String(), elemValue.String())
	}

	if inType.Kind() == reflect.Map {
		keys := inValue.MapKeys()
		for i := 0; i < len(keys); i++ {
			if equal(keys[i].Interface(), elem) {
				return true
			}
		}
	}

	if inType.Kind() == reflect.Slice {
		for i := 0; i < inValue.Len(); i++ {
			if equal(inValue.Index(i).Interface(), elem) {
				return true
			}
		}

	}

	return false
}

// ToMap transforms a slice of instances to a Map
// []*Foo => Map<int, *Foo>
func ToMap(in interface{}, pivot string) interface{} {
	value := reflect.ValueOf(in)

	// input value must be a slice
	if value.Kind() != reflect.Slice {
		panic(fmt.Sprintf("%v must be a slice", in))
	}

	inType := value.Type()

	structType := inType.Elem()

	// retrieve the struct in the slice to deduce key type
	if structType.Kind() == reflect.Ptr {
		structType = structType.Elem()
	}

	field, _ := structType.FieldByName(pivot)

	// value of the map will be the input type
	collectionType := reflect.MapOf(field.Type, inType.Elem())

	// create a map from scratch
	collection := reflect.MakeMap(collectionType)

	for i := 0; i < value.Len(); i++ {
		instance := value.Index(i)
		var field reflect.Value

		if instance.Kind() == reflect.Ptr {
			field = instance.Elem().FieldByName(pivot)
		} else {
			field = instance.FieldByName(pivot)
		}

		collection.SetMapIndex(field, instance)
	}

	return collection.Interface()
}

// Map is ...
func Map(arr interface{}, mapFunc interface{}) interface{} {
	if !IsIterable(arr) {
		panic("First parameter must be neither array nor slice")
	}

	if !IsFunction(mapFunc) {
		panic("Second argument must be function")
	}

	funcValue := reflect.ValueOf(mapFunc)

	funcType := funcValue.Type()

	arrValue := reflect.ValueOf(arr)

	arrType := arrValue.Type()

	arrElemType := arrType.Elem()

	if arrType.Kind() == reflect.Slice || arrType.Kind() == reflect.Array {
		if funcType.NumIn() != 1 || funcType.NumOut() == 0 {
			panic("Map function with an array must have one parameter and must return at least one parameter")
		}

		// Checking whether element type is convertible to function's first argument's type.
		if !arrElemType.ConvertibleTo(funcType.In(0)) {
			panic("Map function's argument is not compatible with type of array.")
		}

		// Get slice type corresponding to function's return value's type.
		resultSliceType := reflect.SliceOf(funcType.Out(0))

		// MakeSlice takes a slice kind type, and makes a slice.
		resultSlice := reflect.MakeSlice(resultSliceType, 0, 0)

		for i := 0; i < arrValue.Len(); i++ {
			result := funcValue.Call([]reflect.Value{arrValue.Index(i)})[0]

			resultSlice = reflect.Append(resultSlice, result)
		}

		return resultSlice.Interface()
	}

	if arrType.Kind() == reflect.Map {
		if funcType.NumIn() != 2 {
			panic("Map function with an array must have one parameter")
		}

		// Only one returned parameter, should be a slice
		if funcType.NumOut() == 1 {
			// Get slice type corresponding to function's return value's type.
			resultSliceType := reflect.SliceOf(funcType.Out(0))

			// MakeSlice takes a slice kind type, and makes a slice.
			resultSlice := reflect.MakeSlice(resultSliceType, 0, 0)

			for _, key := range arrValue.MapKeys() {
				results := funcValue.Call([]reflect.Value{key, arrValue.MapIndex(key)})

				result := results[0]

				resultSlice = reflect.Append(resultSlice, result)
			}

			return resultSlice.Interface()
		}

		// two parameters, should be a map
		if funcType.NumOut() == 2 {
			// value of the map will be the input type
			collectionType := reflect.MapOf(funcType.Out(0), funcType.Out(1))

			// create a map from scratch
			collection := reflect.MakeMap(collectionType)

			for _, key := range arrValue.MapKeys() {
				results := funcValue.Call([]reflect.Value{key, arrValue.MapIndex(key)})

				collection.SetMapIndex(results[0], results[1])

			}

			return collection.Interface()
		}
	}

	panic(fmt.Sprintf("Type %s is not supported by Map", arrType.String()))
}

// SliceOf is ...
func SliceOf(in interface{}) interface{} {
	value := reflect.ValueOf(in)

	sliceType := reflect.SliceOf(reflect.TypeOf(in))
	slice := reflect.New(sliceType)
	sliceValue := reflect.MakeSlice(sliceType, 0, 0)
	sliceValue = reflect.Append(sliceValue, value)
	slice.Elem().Set(sliceValue)

	return slice.Elem().Interface()
}

// ZeroOf is ...
func ZeroOf(in interface{}) interface{} {
	return reflect.Zero(reflect.TypeOf(in)).Interface()
}
