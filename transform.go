package funk

import (
	"fmt"
	"reflect"
)

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
