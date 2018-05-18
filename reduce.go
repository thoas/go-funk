package funk

import (
	"reflect"
)

func Reduce(arr, reduceFunc, acc interface{}) float64 {
	arrValue := redirectValue(reflect.ValueOf(arr))

	if !IsIteratee(arrValue.Interface()) {
		panic("First parameter must be an iteratee")
	}

	returnType := reflect.TypeOf(Reduce).Out(0)

	isFunc := IsFunction(reduceFunc, 2, 1)
	isRune := reflect.TypeOf(reduceFunc).Kind() == reflect.Int32

	if !(isFunc || isRune) {
		panic("Second argument must be a function or rune")
	}

	accValue := reflect.ValueOf(acc)
	sliceElemType := sliceElem(arrValue.Type())

	if isRune {
		if arrValue.Kind() == reflect.Slice && sliceElemType.Kind() == reflect.Interface {
			accValue = accValue.Convert(returnType)
		} else {
			accValue = accValue.Convert(sliceElemType)
		}
	} else {
		accValue = accValue.Convert(reflect.TypeOf(reduceFunc).In(0))
	}

	accType := accValue.Type()

	// Generate reduce function if was passed as rune
	if isRune {
		in := []reflect.Type{accType, sliceElemType}
		out := []reflect.Type{accType}

		funcType := reflect.FuncOf(in, out, false)
		reduceSign := reduceFunc.(int32)

		reduceFunc = reflect.MakeFunc(funcType, func(args []reflect.Value) []reflect.Value {
			acc := args[0].Interface()
			elem := args[1].Interface()

			var result float64
			params := []interface{}{acc, elem}
			switch reduceSign {
			case '+':
				result = Sum(params)
			case '*':
				result = Product(params)
			}

			return []reflect.Value{reflect.ValueOf(result).Convert(accType)}
		}).Interface()
	}

	funcValue := reflect.ValueOf(reduceFunc)
	funcType := funcValue.Type()

	for i := 0; i < arrValue.Len(); i++ {
		if accType.ConvertibleTo(funcType.In(0)) {
			accValue = accValue.Convert(funcType.In(0))
		}

		arrElementValue := arrValue.Index(i)
		if sliceElemType.ConvertibleTo(funcType.In(1)) {
			arrElementValue = arrElementValue.Convert(funcType.In(1))
		}

		result := funcValue.Call([]reflect.Value{accValue, arrElementValue})
		accValue = result[0]
	}

	resultInterface := accValue.Convert(returnType).Interface()
	result, _ := resultInterface.(float64)
	return result
}
