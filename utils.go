package funk

import "reflect"

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
