package funk

import (
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
