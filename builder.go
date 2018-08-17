package funk

import (
	"fmt"
	"reflect"
)

// Builder ...
type Builder interface {
	Chunk(size int) Builder
	Compact() Builder
	Drop(n int) Builder
	Filter(predicate interface{}) Builder
	FlattenDeep() Builder
	Initial() Builder
	Intersect(y interface{}) Builder
	Map(mapFunc interface{}) Builder
	Reverse() Builder
	Shuffle() Builder
	Tail() Builder
	Uniq() Builder

	All() bool
	Any() bool
	Contains(elem interface{}) bool
	Every(elements ...interface{}) bool
	Find(predicate interface{}) interface{}
	ForEach(predicate interface{})
	ForEachRight(predicate interface{})
	Head() interface{}
	Keys() interface{}
	IndexOf(elem interface{}) int
	IsEmpty() bool
	IsType(actual interface{}) bool
	Last() interface{}
	LastIndexOf(elem interface{}) int
	NotEmpty() bool
	Product() float64
	Reduce(reduceFunc, acc interface{}) float64
	Sum() float64
	Type() reflect.Type
	Value() interface{}
	Values() interface{}
}

// Chain ...
func Chain(v interface{}) Builder {
	isNotNil(v, "Chain")

	valueType := reflect.TypeOf(v)
	if isValidBuilderEntry(valueType) ||
		(valueType.Kind() == reflect.Ptr && isValidBuilderEntry(valueType.Elem())) {
		return &chainBuilder{v}
	}

	panic(fmt.Sprintf("Type %s is not supported by Chain", valueType.String()))
}

// LazyChain ...
func LazyChain(v interface{}) Builder {
	isNotNil(v, "LazyChain")

	valueType := reflect.TypeOf(v)
	if isValidBuilderEntry(valueType) ||
		(valueType.Kind() == reflect.Ptr && isValidBuilderEntry(valueType.Elem())) {
		return &lazyBuilder{func() interface{} { return v }}
	}

	panic(fmt.Sprintf("Type %s is not supported by LazyChain", valueType.String()))

}

// LazyChainWith ...
func LazyChainWith(generator func() interface{}) Builder {
	isNotNil(generator, "LazyChainWith")
	return &lazyBuilder{func() interface{} {
		isNotNil(generator, "LazyChainWith")

		v := generator()
		valueType := reflect.TypeOf(v)
		if isValidBuilderEntry(valueType) ||
			(valueType.Kind() == reflect.Ptr && isValidBuilderEntry(valueType.Elem())) {
			return v
		}

		panic(fmt.Sprintf("Type %s is not supported by LazyChainWith generator", valueType.String()))
	}}
}

func isNotNil(v interface{}, from string) {
	if v == nil {
		panic(fmt.Sprintf("nil value is not supported by %s", from))
	}
}

func isValidBuilderEntry(valueType reflect.Type) bool {
	return valueType.Kind() == reflect.Slice || valueType.Kind() == reflect.Array ||
		valueType.Kind() == reflect.Map ||
		valueType.Kind() == reflect.String
}
