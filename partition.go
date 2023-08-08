package funk

import (
	"reflect"
)

// Partition separates the elements of the input array into two slices based on the predicate function.
// It takes an array-like data structure and a predicate function that determines the Partition.
// The predicate function should have the signature func(elementType) bool.
// The function returns two new slices: one containing elements that satisfy the predicate,
// and the other containing elements that do not satisfy the predicate.
func Partition(in, predicate interface{}) interface{} {
	if !IsCollection(in) {
		panic("First parameter must be a collection")
	}

	if !IsFunction(predicate, 1, 1) {
		panic("Second argument must be function")
	}

	inValue, funcValue := reflect.ValueOf(in), reflect.ValueOf(predicate)

	funcType := funcValue.Type()

	if funcType.Out(0).Kind() != reflect.Bool {
		panic("Return argument should be a boolean")
	}

	result := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(in)), 0, 0)
	partitionOne, partitionTwo := reflect.MakeSlice(inValue.Type(), 0, 0), reflect.MakeSlice(inValue.Type(), 0, 0)

	for i := 0; i < inValue.Len(); i++ {
		element := inValue.Index(i)

		res := funcValue.Call([]reflect.Value{reflect.ValueOf(element.Interface())})
		if res[0].Interface().(bool) {
			partitionOne = reflect.Append(partitionOne, element)
		} else {
			partitionTwo = reflect.Append(partitionTwo, element)
		}
	}

	if partitionOne.Len() > 0 || partitionTwo.Len() > 0 {
		result = reflect.Append(result, partitionOne, partitionTwo)
	}

	return result.Interface()
}
