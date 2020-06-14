package funk

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func Set(in interface{}, val interface{}, path string) error {
	inValue := reflect.ValueOf(in)
	inKind := inValue.Type().Kind()

	if inKind == reflect.Ptr {
		inValue = inValue.Elem()
	}

	// todo change checks
	if !inValue.CanSet() && !IsIteratee(in) {
		panic(fmt.Sprintf("Type %s cannot be set", inValue.Type().String()))
	}

	parts := []string{}
	if path != "" {
		parts = strings.Split(path, ".")
	}

	return set(inValue, reflect.ValueOf(val), parts)
}

// Set assigns struct field with val at path
// i.e. in.path = val
func set(inValue reflect.Value, setValue reflect.Value, parts []string) error {

	//inKind := inValue.Type().Kind()

	// traverse the path to get the inValue we need to set
	for _, part := range parts {
		inValue = redirectValue(inValue)
		kind := inValue.Kind()

		if kind == reflect.Invalid {
			// TODO: decide if initilize the struct and continue
			return errors.New("nil pointer found along the path")
		}

		if kind == reflect.Struct {
			inValue = inValue.FieldByName(part)
			if !inValue.IsValid() {
				return fmt.Errorf("field name %v is not found in struct %v", part, kind.String())
			}
			if !inValue.CanSet() {
				panic(fmt.Sprintf("Type %s cannot be set", inValue.Type().String()))
			}
			continue
		}

		if kind == reflect.Slice || kind == reflect.Array {
			// set all its elements
			length := inValue.Len()
			for i := 0; i < length; i++ {
				err := set(inValue.Index(i), setValue, parts)
				if err != nil {
					return err
				}
			}
			return nil
		}

		return fmt.Errorf("field name %s not found", part)
	}

	// inValue holds the value we need to set
	if !inValue.CanSet() {
		panic("field not addressable or unexported")
	}

	// change value of
	if inValue.Kind() != setValue.Kind() {
		panic("type not match")
	}

	inValue.Set(setValue)

	return nil
}
