package funk

import (
	"reflect"
	"strings"
)

func Set(in interface{}, val interface{}, path string) error {
	inValue := reflect.ValueOf(in)

	// TODO if it is not ptr then panic?
	if inValue.Type().Kind() == reflect.Ptr {
		inValue = inValue.Elem()
	}

	return set(inValue, reflect.ValueOf(val), path)
}

// Set assigns struct field with val at path
// i.e. in.path = val
func set(inValue reflect.Value, setValue reflect.Value, path string) error {

	inKind := inValue.Type().Kind()

	if inKind == reflect.Slice || inKind == reflect.Array {
		panic("TODO array")
	}

	currRest := strings.SplitN(path, ".", 2 /* num of substr */)
	if len(currRest) == 0 {
		// set in to be val
		panic("TODO no path case")
	}

	if len(currRest) == 1 {
		// set the in's field curr to be val
		setStructField(inValue, setValue, currRest[0])
		return nil
	}

	// len must be 2
	curr := currRest[0]
	rest := currRest[1]
	// recursive
	fValue := inValue.FieldByName(curr)
	if !fValue.IsValid() {
		panic("field not found")
	}

	return set(fValue, setValue, rest)

}

// s is struct
func setStructField(s reflect.Value, setValue reflect.Value, name string) {

	// s := ps.Elem()
	if s.Kind() != reflect.Struct {
		panic("s is not struct")
	}

	// exported field
	f := s.FieldByName(name)
	if !f.IsValid() {
		panic("field not found")
	}
	// A Value can be changed only if it is
	// addressable and was not obtained by
	// the use of unexported struct fields.
	if !f.CanSet() {
		panic("field not addressable or unexported")
	}

	// change value of
	if f.Kind() != setValue.Kind() {
		panic("type not match")
	}

	f.Set(setValue)

}
