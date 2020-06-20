package funk

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func Set(in interface{}, val interface{}, path string) error {
	if in == nil {
		return errors.New("Cannot Set nil")
	}
	parts := []string{}
	if path != "" {
		parts = strings.Split(path, ".")
	}
	return setByParts(in, val, parts)
}

func setByParts(in interface{}, val interface{}, parts []string) error {

	if in == nil {
		// nil interface can happen during traversing the path
		return errors.New("Cannot traverse nil interface{}")
	}

	inValue := reflect.ValueOf(in)
	inKind := inValue.Type().Kind()

	// Note: if interface contains a struct (not ptr to struct) then the content of the struct cannot be set.
	// I.e. it is not CanAddr() or CanSet()
	// So we require in interface{} to be a ptr, slice or array
	if inKind == reflect.Ptr {
		inValue = inValue.Elem() // if it is ptr we set its content not ptr its self
	} else if inKind != reflect.Array && inKind != reflect.Slice {
		panic(fmt.Sprintf("Type %s not supported by Set", inValue.Type().String()))
	}

	return set(inValue, reflect.ValueOf(val), parts)
}

// Set assigns struct field with val at path
// i.e. in.path = val
func set(inValue reflect.Value, setValue reflect.Value, parts []string) error {

	//log.Println(parts)
	//inKind := inValue.Type().Kind()

	// traverse the path to get the inValue we need to set
	i := 0
	for i < len(parts) {

		kind := inValue.Kind()

		switch kind {
		case reflect.Invalid:
			return errors.New("nil pointer found along the path")
		case reflect.Struct:
			fValue := inValue.FieldByName(parts[i])
			if !fValue.IsValid() {
				return fmt.Errorf("field name %v is not found in struct %v", parts[i], inValue.Type().String())
			}
			if !fValue.CanSet() {
				panic(fmt.Sprintf("field name %v is not exported in struct %v", parts[i], inValue.Type().String()))
			}
			inValue = fValue
			i++
		case reflect.Slice | reflect.Array:
			// set all its elements
			length := inValue.Len()
			for j := 0; j < length; j++ {
				err := set(inValue.Index(j), setValue, parts[i:])
				if err != nil {
					return err
				}
			}
			return nil
		case reflect.Ptr:
			// only traverse down one level
			if inValue.IsNil() {
				// we initilize nil ptr to ptr to zero value of the type
				// and continue traversing
				inValue.Set(reflect.New(inValue.Type().Elem()))
			}
			// traverse the ptr until it is not pointer any more or is nil again
			inValue = redirectValue(inValue)
		case reflect.Interface:
			// Note: if interface contains a struct (not ptr to struct) then the content of the struct cannot be set.
			// I.e. it is not CanAddr() or CanSet()
			// we treat this as a new call to setByParts, and it will do proper check of the types
			return setByParts(inValue.Interface(), setValue.Interface(), parts[i:])
		default:
			panic(fmt.Sprintf("kind %v in path is not supported", kind))
		}

	}

	// inValue holds the value we need to set
	if !inValue.CanSet() {
		// not expect to hit here
		panic(fmt.Sprintf("field not addressable or unexported of type %v", inValue.Kind()))
	}

	// interface{} can be set to any val
	// other types we ensure the type matches
	if inValue.Kind() != setValue.Kind() && inValue.Kind() != reflect.Interface {
		panic(fmt.Sprintf("type not match: target %v, arg %v", inValue.Kind(), setValue.Kind()))
	}
	inValue.Set(setValue)

	return nil
}
