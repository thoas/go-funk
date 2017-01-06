package funk

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

func sliceElem(rtype reflect.Type) reflect.Type {
	if rtype.Kind() == reflect.Slice || rtype.Kind() == reflect.Array {
		return sliceElem(rtype.Elem())
	}

	return rtype
}

func redirectValue(value reflect.Value) reflect.Value {
	if value.Kind() == reflect.Ptr {
		return redirectValue(value.Elem())
	}

	return value
}

// FlattenDeep recursively flattens array.
func FlattenDeep(out interface{}) interface{} {
	return flattenDeep(reflect.ValueOf(out)).Interface()
}

func flattenDeep(value reflect.Value) reflect.Value {
	sliceType := sliceElem(value.Type())

	resultSlice := reflect.MakeSlice(reflect.SliceOf(sliceType), 0, 0)

	return flatten(value, resultSlice)
}

func flatten(value reflect.Value, result reflect.Value) reflect.Value {
	length := value.Len()

	for i := 0; i < length; i++ {
		item := value.Index(i)
		kind := item.Kind()

		if kind == reflect.Slice || kind == reflect.Array {
			result = flatten(item, result)
		} else {
			result = reflect.Append(result, item)
		}
	}

	return result
}

// Get retrieves the value at path of struct(s).
func Get(out interface{}, path string) interface{} {
	result := get(reflect.ValueOf(out), path)

	if result.Kind() != reflect.Invalid {
		return result.Interface()
	}

	return nil
}

// Keys creates an array of the own enumerable map keys or struct field names.
func Keys(out interface{}) interface{} {
	value := redirectValue(reflect.ValueOf(out))
	valueType := value.Type()

	if value.Kind() == reflect.Map {
		keys := value.MapKeys()

		length := len(keys)

		resultSlice := reflect.MakeSlice(reflect.SliceOf(valueType.Key()), length, length)

		for i, key := range keys {
			resultSlice.Index(i).Set(key)
		}

		return resultSlice.Interface()
	}

	if value.Kind() == reflect.Struct {
		length := value.NumField()

		resultSlice := make([]string, length)

		for i := 0; i < length; i++ {
			resultSlice[i] = valueType.Field(i).Name
		}

		return resultSlice
	}

	panic(fmt.Sprintf("Type %s is not supported by Keys", valueType.String()))
}

// Values creates an array of the own enumerable map values or struct field values.
func Values(out interface{}) interface{} {
	value := redirectValue(reflect.ValueOf(out))
	valueType := value.Type()

	if value.Kind() == reflect.Map {
		keys := value.MapKeys()

		length := len(keys)

		resultSlice := reflect.MakeSlice(reflect.SliceOf(valueType.Elem()), length, length)

		for i, key := range keys {
			resultSlice.Index(i).Set(value.MapIndex(key))
		}

		return resultSlice.Interface()
	}

	if value.Kind() == reflect.Struct {
		length := value.NumField()

		resultSlice := make([]interface{}, length)

		for i := 0; i < length; i++ {
			resultSlice[i] = value.Field(i).Interface()
		}

		return resultSlice
	}

	panic(fmt.Sprintf("Type %s is not supported by Keys", valueType.String()))
}

func get(value reflect.Value, path string) reflect.Value {
	if value.Kind() == reflect.Slice || value.Kind() == reflect.Array {
		var resultSlice reflect.Value

		length := value.Len()

		for i := 0; i < length; i++ {
			item := value.Index(i)

			resultValue := get(item, path)

			if resultValue.Kind() == reflect.Invalid {
				continue
			}

			resultType := resultValue.Type()

			if resultSlice.Kind() == reflect.Invalid {
				resultType := reflect.SliceOf(resultType)

				resultSlice = reflect.MakeSlice(resultType, 0, 0)
			}

			resultSlice = reflect.Append(resultSlice, resultValue)
		}

		// if the result is a slice of a slice, we need to flatten it
		if resultSlice.Type().Elem().Kind() == reflect.Slice {
			return flattenDeep(resultSlice)
		}

		return resultSlice
	}

	parts := strings.Split(path, ".")

	for _, part := range parts {
		value = redirectValue(value)
		kind := value.Kind()

		if kind == reflect.Invalid {
			continue
		}

		if kind == reflect.Struct {
			value = value.FieldByName(part)
			continue
		}

		if kind == reflect.Slice || kind == reflect.Array {
			value = get(value, part)
			continue
		}
	}

	return value
}

// Chunk creates an array of elements split into groups with the length of size.
// If array can't be split evenly, the final chunk will be
// the remaining element.
func Chunk(arr interface{}, size int) interface{} {
	if !IsIteratee(arr) {
		panic("First parameter must be neither array nor slice")
	}

	arrValue := reflect.ValueOf(arr)

	arrType := arrValue.Type()

	resultSliceType := reflect.SliceOf(arrType)

	// Initialize final result slice which will contains slice
	resultSlice := reflect.MakeSlice(resultSliceType, 0, 0)

	itemType := arrType.Elem()

	var itemSlice reflect.Value

	itemSliceType := reflect.SliceOf(itemType)

	length := arrValue.Len()

	for i := 0; i < length; i++ {
		if i%size == 0 || i == 0 {
			if itemSlice.Kind() != reflect.Invalid {
				resultSlice = reflect.Append(resultSlice, itemSlice)
			}

			itemSlice = reflect.MakeSlice(itemSliceType, 0, 0)
		}

		itemSlice = reflect.Append(itemSlice, arrValue.Index(i))

		if i == length-1 {
			resultSlice = reflect.Append(resultSlice, itemSlice)
		}
	}

	return resultSlice.Interface()
}

// ForEach iterates over elements of collection and invokes iteratee
// for each element.
func ForEach(arr interface{}, predicate interface{}) {
	if !IsIteratee(arr) {
		panic("First parameter must be an iteratee")
	}

	var (
		funcValue = reflect.ValueOf(predicate)
		arrValue  = reflect.ValueOf(arr)
		arrType   = arrValue.Type()
		funcType  = funcValue.Type()
	)

	if arrType.Kind() == reflect.Slice || arrType.Kind() == reflect.Array {
		if !IsFunction(predicate, 1, 0) {
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
		if !IsFunction(predicate, 2, 0) {
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

// Filter iterates over elements of collection, returning an array of
// all elements predicate returns truthy for.
func Filter(arr interface{}, predicate interface{}) interface{} {
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

	return resultSlice.Interface()
}

// Find iterates over elements of collection, returning the first
// element predicate returns truthy for.
func Find(arr interface{}, predicate interface{}) interface{} {
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

	for i := 0; i < arrValue.Len(); i++ {
		elem := arrValue.Index(i)

		result := funcValue.Call([]reflect.Value{elem})[0].Interface().(bool)

		if result == true {
			return elem.Interface()
		}
	}

	return nil
}

// IndexOf gets the index at which the first occurrence of value is found in array or return -1
// if the value cannot be found
func IndexOf(in interface{}, elem interface{}) int {
	inValue := reflect.ValueOf(in)

	elemValue := reflect.ValueOf(elem)

	inType := inValue.Type()

	if inType.Kind() == reflect.String {
		return strings.Index(inValue.String(), elemValue.String())
	}

	if inType.Kind() == reflect.Slice {
		for i := 0; i < inValue.Len(); i++ {
			if equal(inValue.Index(i).Interface(), elem) {
				return i
			}
		}
	}

	return -1
}

// Contains returns true if an element is present in a iteratee.
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

// ToMap transforms a slice of instances to a Map.
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

// Map manipulates an iteratee and transforms it to another type.
func Map(arr interface{}, mapFunc interface{}) interface{} {
	if !IsIteratee(arr) {
		panic("First parameter must be an iteratee")
	}

	if !IsFunction(mapFunc) {
		panic("Second argument must be function")
	}

	var (
		funcValue = reflect.ValueOf(mapFunc)
		funcType  = funcValue.Type()
		arrValue  = reflect.ValueOf(arr)
		arrType   = arrValue.Type()
	)

	if arrType.Kind() == reflect.Slice || arrType.Kind() == reflect.Array {
		if funcType.NumIn() != 1 || funcType.NumOut() == 0 {
			panic("Map function with an array must have one parameter and must return at least one parameter")
		}

		arrElemType := arrType.Elem()

		// Checking whether element type is convertible to function's first argument's type.
		if !arrElemType.ConvertibleTo(funcType.In(0)) {
			panic("Map function's argument is not compatible with type of array.")
		}

		if funcType.NumOut() == 1 {
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

		if funcType.NumOut() == 2 {
			// value of the map will be the input type
			collectionType := reflect.MapOf(funcType.Out(0), funcType.Out(1))

			// create a map from scratch
			collection := reflect.MakeMap(collectionType)

			for i := 0; i < arrValue.Len(); i++ {
				results := funcValue.Call([]reflect.Value{arrValue.Index(i)})

				collection.SetMapIndex(results[0], results[1])
			}

			return collection.Interface()
		}
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

// SliceOf will returns a slice which contains the element.
func SliceOf(in interface{}) interface{} {
	value := reflect.ValueOf(in)

	sliceType := reflect.SliceOf(reflect.TypeOf(in))
	slice := reflect.New(sliceType)
	sliceValue := reflect.MakeSlice(sliceType, 0, 0)
	sliceValue = reflect.Append(sliceValue, value)
	slice.Elem().Set(sliceValue)

	return slice.Elem().Interface()
}

// ZeroOf will returns a zero value of an element.
func ZeroOf(in interface{}) interface{} {
	return reflect.Zero(reflect.TypeOf(in)).Interface()
}
