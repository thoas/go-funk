package funk

import (
	"reflect"
)

// Nth returns the element at the specified index from the array (Slice).
// The index is 1-based, so the first element is at index 1.
func Nth(in interface{}, number int) interface{} {
	if !IsCollection(in) {
		panic("First parameter must be a collection")
	}

	inValue := reflect.ValueOf(in)

	if number >= inValue.Len() {
		return nil
	}

	if number < 0 {
		if number+inValue.Len() < 0 {
			return nil
		}
		number = number + inValue.Len()
	}

	return inValue.Index(number).Interface()
}
