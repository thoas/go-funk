package funk

import (
	"reflect"
	"strings"
)

// Max checks for the type of input, validates the input, compares the elements and returns the maximum element in an array/slice.
// It accepts []int, []int8, []int16, []int32, []int64, []float32, []float64, []string
// For strings, the first different character amongst the elements is compared using go's inbuilt strings.Compare method
// It returns int or string or nil
// It returns nil for the following cases:
//  - input is not an array/slice
//  - input has heterogeneous elements
func Max(input interface{}) interface{} {

	i := reflect.ValueOf(input)

	if (i.Kind() != reflect.Slice && i.Kind() != reflect.Array) || i.Len() == 0 {
		return nil
	}

	var max reflect.Value

	for idx := 0; idx < i.Len(); idx++ {
		item := i.Index(idx)

		if idx == 0 {
			max = item
			continue
		}

		if item.Kind() != i.Index(idx - 1).Kind() {
			return nil
		}

		switch item.Kind() {
			case reflect.Int:
				if int(item.Int()) > int(max.Int()) {
					max = item
				}
				break
			case reflect.Int8:
				if int8(item.Int()) > int8(max.Int()) {
					max = item
				}
				break
			case reflect.Int16:
				if int16(item.Int()) > int16(max.Int()) {
					max = item
				}
				break
			case reflect.Int32:
				if int32(item.Int()) > int32(max.Int()) {
					max = item
				}
				break
			case reflect.Int64:
				if int64(item.Int()) > int64(max.Int()) {
					max = item
				}
				break
			case reflect.Float32:
				if float32(item.Float()) > float32(max.Float()) {
					max = item
				}
				break
			case reflect.Float64:
				if float64(item.Float()) > float64(max.Float()) {
					max = item
				}
				break
			case reflect.String:
				s := compareStrings(max.String(), item.String())
				max = reflect.ValueOf(s)
				break
		}

	}

	switch max.Kind() {

		case reflect.Int:
			return int(max.Int())
		case reflect.Int8:
			return int8(max.Int())
		case reflect.Int16:
			return int16(max.Int())
		case reflect.Int32:
			return int32(max.Int())
		case reflect.Int64:
			return int64(max.Int())
		case reflect.Float32:
			return float32(max.Float())
		case reflect.Float64:
			return float64(max.Float())
		case reflect.String:
			return max.String()
		default:
			return nil

	}

}

func compareStrings(max, current string) string {

	r := strings.Compare(strings.ToLower(max), strings.ToLower(current))

	if r > 0 {
		return max
	} else {
		return current
	}

}
