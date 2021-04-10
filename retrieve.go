package funk

import (
	"reflect"
	"strings"
)

// Get retrieves the value from given path, retriever can be modified with available RetrieverOptions
func Get(out interface{}, path string, opts ...RetrieverOption) interface{} {
	return newRetriever(opts...).Get(out, path)
}

// GetOrElse retrieves the value of the pointer or default.
func GetOrElse(v interface{}, def interface{}) interface{} {
	val := reflect.ValueOf(v)
	if v == nil || (val.Kind() == reflect.Ptr && val.IsNil()) {
		return def
	} else if val.Kind() != reflect.Ptr {
		return v
	}
	return val.Elem().Interface()
}

type RetrieverOption func(retriever *retriever)

func WithAllowZero() RetrieverOption {
	return func(r *retriever) {
		r.AllowZero = true
	}
}

type retriever struct {
	AllowZero bool
}

func newRetriever(opts ...RetrieverOption) *retriever {
	r := &retriever{
		AllowZero: false,
	}
	for _, opt := range opts {
		opt(r)
	}
	// return the modified retriever instance
	return r
}

// Get retrieves the value at path of struct(s).
func (r retriever) Get(out interface{}, path string) interface{} {
	result := r.get(reflect.ValueOf(out), path)
	// valid kind and we can return a result.Interface() without panic
	if result.Kind() != reflect.Invalid && result.CanInterface() {
		// if we don't allow zero and the result is a zero value return nil
		if !r.AllowZero && result.IsZero() {
			return nil
		}
		// if the result kind is a pointer and its nil return nil
		if result.Kind() == reflect.Ptr && result.IsNil() {
			return nil
		}
		// return the result interface (i.e the zero value of it)
		return result.Interface()
	}

	return nil
}

func (r retriever) get(value reflect.Value, path string) reflect.Value {
	if value.Kind() == reflect.Slice || value.Kind() == reflect.Array {
		var resultSlice reflect.Value

		length := value.Len()

		if length == 0 {
			zeroElement := reflect.Zero(value.Type().Elem())
			pathValue := r.get(zeroElement, path)
			value = reflect.MakeSlice(reflect.SliceOf(pathValue.Type()), 0, 0)

			return value
		}

		for i := 0; i < length; i++ {
			item := value.Index(i)

			resultValue := r.get(item, path)

			if resultValue.Kind() == reflect.Invalid || resultValue.IsZero() {
				continue
			}

			resultType := resultValue.Type()

			if resultSlice.Kind() == reflect.Invalid {
				resultType := reflect.SliceOf(resultType)

				resultSlice = reflect.MakeSlice(resultType, 0, 0)
			}

			resultSlice = reflect.Append(resultSlice, resultValue)
		}

		// if the result is a slice of a slice, we need to flatten it
		if resultSlice.Kind() != reflect.Invalid && resultSlice.Type().Elem().Kind() == reflect.Slice {
			return flattenDeep(resultSlice)
		}

		return resultSlice
	}

	parts := strings.Split(path, ".")

	for _, part := range parts {
		value = redirectValue(value)
		kind := value.Kind()

		switch kind {
		case reflect.Invalid:
			continue
		case reflect.Struct:
			value = value.FieldByName(part)
		case reflect.Map:
			value = value.MapIndex(reflect.ValueOf(part))
		case reflect.Slice, reflect.Array:
			value = r.get(value, part)
		default:
			return reflect.ValueOf(nil)
		}
	}

	return value
}
