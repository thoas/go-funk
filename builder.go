package funk

import (
	"reflect"
)

type FunkBuilder interface {
	Chunk(size int) FunkBuilder
	Compact() FunkBuilder
	Drop(in interface{}, n int) FunkBuilder
	Filter(predicate interface{}) FunkBuilder
	FlattenDeep() FunkBuilder
	ForEach(predicate interface{}) FunkBuilder
	ForEachRight(predicate interface{}) FunkBuilder
	Initial() FunkBuilder
	Intersect(y interface{}) FunkBuilder
	Map(mapFunc interface{}) FunkBuilder
	Reverse() FunkBuilder
	Shuffle() FunkBuilder
	Uniq() FunkBuilder

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
