package funk

import (
	"fmt"
	"log"
	"reflect"
)

func calculate(arr interface{}, name string, operation rune) float64 {
	value := redirectValue(reflect.ValueOf(arr))
	valueType := value.Type()

	kind := value.Kind()

	if kind == reflect.Array || kind == reflect.Slice {
		length := value.Len()

		if length == 0 {
			return 0
		}

		result := map[rune]float64{
			'+': 0.0,
			'*': 1,
		}[operation]

		for i := 0; i < length; i++ {
			elem := redirectValue(value.Index(i)).Interface()

			var value float64
			switch elem.(type) {
			case int:
				value = float64(elem.(int))
			case int8:
				value = float64(elem.(int8))
			case int16:
				value = float64(elem.(int16))
			case int32:
				value = float64(elem.(int32))
			case int64:
				value = float64(elem.(int64))
			case float32:
				value = float64(elem.(float32))
			case float64:
				value = elem.(float64)
			}

			switch operation {
			case '+':
				result += value
			case '*':
				result *= value
			}
		}

		return result
	}

	panic(fmt.Sprintf("Type %s is not supported by %s", valueType.String(), name))
}

// Sum computes the sum of the values in array.
func Sum(arr interface{}) float64 {
	return calculate(arr, "Sum", '+')
}

// Product computes the product of the values in array.
func Product(arr interface{}) float64 {
	return calculate(arr, "Product", '*')
}

// SumBy computes the sum of the values of the key in array
func SumBy(arr interface{}, key string) float64 {
	arrayValue := redirectValue(reflect.ValueOf(arr))

	kind := arrayValue.Kind()
	if kind == reflect.Array || kind == reflect.Slice {
		length := arrayValue.Len()

		if length == 0 {
			return 0
		}

		var sum float64 = 0
		for i := 0; i < length; i++ {
			itemValue := redirectValue(reflect.ValueOf(arrayValue.Index(i).Interface())).FieldByName(key)
			if isInt(itemValue) {
				sum += float64(itemValue.Int())
			} else if isFloat(itemValue) {
				sum += itemValue.Float()
			}
		}
		log.Println("sum is", sum)
		return sum
	}
	return 0
}

func isInt(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.Int:
		return true
	case reflect.Int8:
		return true
	case reflect.Int16:
		return true
	case reflect.Int32:
		return true
	case reflect.Int64:
		return true
	default:
		return false
	}
}

func isFloat(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.Float32:
		return true
	case reflect.Float64:
		return true
	default:
		return false
	}
}
