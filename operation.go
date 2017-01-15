package funk

import (
	"fmt"
	"reflect"
)

// Sum computes the sum of the values in array.
func Sum(arr interface{}) float64 {
	value := redirectValue(reflect.ValueOf(arr))
	valueType := value.Type()

	kind := value.Kind()

	if kind == reflect.Array || kind == reflect.Slice {
		length := value.Len()

		if length == 0 {
			return 0
		}

		result := 0.0

		for i := 0; i < length; i++ {
			elem := redirectValue(value.Index(i)).Interface()

			switch elem.(type) {
			case int:
				result += float64(elem.(int))
			case int64:
				result += float64(elem.(int64))
			case float32:
				result += float64(elem.(float32))
			case float64:
				result += elem.(float64)
			}
		}

		return result
	}

	panic(fmt.Sprintf("Type %s is not supported by Sum", valueType.String()))
}
