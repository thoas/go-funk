package funk

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
)

func Set(in interface{}, val interface{}, path string) error {
	inValue := reflect.ValueOf(in)
	inKind := inValue.Type().Kind()

	if inKind == reflect.Ptr {
		inValue = inValue.Elem() // if it is ptr we set its content not ptr its self
	} else if inKind != reflect.Array && inKind != reflect.Slice {
		panic(fmt.Sprintf("Type %s not supported by Set", inValue.Type().String()))
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

	log.Println(parts)
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
				return fmt.Errorf("field name %v is not found in struct %v", parts[i], kind.String())
			}
			if !fValue.CanSet() {
				panic(fmt.Sprintf("Type %s cannot be set", inValue.Type().String()))
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
				// set the nil pointer to be the pointer to zero value of the type
				inValue.Set(reflect.New(inValue.Type().Elem()))
			}
			inValue = inValue.Elem()
		case reflect.Interface:
			inValue = inValue.Elem()
		default:
			// TODO handle interface{} case
			panic(fmt.Sprintf("kind %v in path is not supported", kind))
		}

	}

	// inValue holds the value we need to set
	if !inValue.CanSet() {
		panic("field not addressable or unexported")
	}

	// change value of
	if inValue.Kind() != setValue.Kind() && inValue.Kind() != reflect.Interface {
		panic(fmt.Sprintf("type not match: target %v, arg %v", inValue.Kind(), setValue.Kind()))
	}

	inValue.Set(setValue)

	return nil
}
