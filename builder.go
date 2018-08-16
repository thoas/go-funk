package funk

import (
	"fmt"
	"reflect"
)

// Builder ...
type Builder interface {
	Chunk(size int) Builder
	Compact() Builder
	Drop(in interface{}, n int) Builder
	Filter(predicate interface{}) Builder
	FlattenDeep() Builder
	ForEach(predicate interface{}) Builder
	ForEachRight(predicate interface{}) Builder
	Initial() Builder
	Intersect(y interface{}) Builder
	Map(mapFunc interface{}) Builder
	Reverse() Builder
	Shuffle() Builder
	Uniq() Builder

	All() bool
	Any() bool
	Contains(elem interface{}) bool
	Every(elements ...interface{}) bool
	Find(predicate interface{}) interface{}
	Get(path string) interface{}
	Head() interface{}
	Keys() interface{}
	In(v interface{}) bool
	IndexOf(elem interface{}) int
	IsEmpty() bool
	IsType(actual interface{}) bool
	Last() interface{}
	NotEmpty() bool
	Product() float64
	Reduce(reduceFunc, acc interface{}) float64
	Sum() float64
	Tail() interface{}
	Type() reflect.Type
	Value() interface{}
	Values() interface{}
}

// Chain ...
func Chain(v interface{}) Builder {
	valueType := reflect.TypeOf(v)
	if valueType.Kind() == reflect.Slice || valueType.Kind() == reflect.Array || valueType.Kind() == reflect.Map {
		return &chainBuilder{v}
	}

	panic(fmt.Sprintf("Type %s is not supported by Chain", valueType.String()))
}

// LazyChain ...
func LazyChain(v interface{}) Builder {
	valueType := reflect.TypeOf(v)
	if valueType.Kind() == reflect.Slice || valueType.Kind() == reflect.Array || valueType.Kind() == reflect.Map {
		return &lazyBuilder{func() interface{} { return v }}
	}
	panic(fmt.Sprintf("Type %s is not supported by LazyChain", valueType.String()))

}

// LazyChainWith ...
func LazyChainWith(generator func() interface{}) Builder {
	return &lazyBuilder{func() interface{} {
		v := generator()
		valueType := reflect.TypeOf(v)
		if valueType.Kind() == reflect.Slice || valueType.Kind() == reflect.Array || valueType.Kind() == reflect.Map {
			return v
		}
		panic(fmt.Sprintf("Type %s is not supported by LazyChainWith generator", valueType.String()))
	}}
}
