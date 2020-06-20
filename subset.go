
package funk

import (
	"reflect"
)

// Subset returns if collection x is a subset of y.
func Subset(x interface{}, y interface{}) boolean {
	if !IsCollection(x) {
		panic("First parameter must be a collection")
	}
	if !IsCollection(y) {
		panic("Second parameter must be a collection")
	}

	//hash := map[interface{}]struct{}{}

	xValue := reflect.ValueOf(x)
	xType := xValue.Type()

	yValue := reflect.ValueOf(y)
	yType := yValue.Type()

	if NotEqual(xType, yType) {
		panic("Parameters must have the same type")
	}

  if yValue == nil || len(yValue)==0 ||  len(yValue)<len(xValue) {
    return false
  }
  
  if xValue == nil || len(xValue)==0{
    return true
  }
  
  for _, elem := range xValue {
		if !Contains(yValue, elem) {
			return false
		}
	}
	return true
}


func SubsetString(x []string, y []string) boolean {
	if len(x) == 0 {
		return true
	}

  if len(y) == 0 || len(x)>len(y) {
		return false
	}
  
  for _, stringElem := range x {
		if !Contains(y, stringElem) {
			return false
		}
	}
	return true
	
}
