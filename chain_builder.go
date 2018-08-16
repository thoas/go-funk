package funk

import "reflect"

type chainBuilder struct {
	collection interface{}
}

func (b *chainBuilder) Chunk(size int) FunkBuilder {
	return &chainBuilder{Chunk(b.collection, size)}
}
func (b *chainBuilder) Compact() FunkBuilder {
	return &chainBuilder{Compact(b.collection)}
}
func (b *chainBuilder) Drop(in interface{}, n int) FunkBuilder {
	return &chainBuilder{Drop(b.collection, n)}
}
func (b *chainBuilder) Filter(predicate interface{}) FunkBuilder {
	return &chainBuilder{Filter(b.collection, predicate)}
}
func (b *chainBuilder) FlattenDeep() FunkBuilder {
	return &chainBuilder{FlattenDeep(b.collection)}
}
func (b *chainBuilder) ForEach(predicate interface{}) FunkBuilder {
	c := deepClone(b.collection)
	ForEach(c, predicate)
	return &chainBuilder{c}
}
func (b *chainBuilder) ForEachRight(predicate interface{}) FunkBuilder {
	c := deepClone(b.collection)
	ForEachRight(c, predicate)
	return &chainBuilder{c}
}
func (b *chainBuilder) Initial() FunkBuilder {
	return &chainBuilder{Initial(b.collection)}
}
func (b *chainBuilder) Intersect(y interface{}) FunkBuilder {
	return &chainBuilder{Intersect(b.collection, y)}
}
func (b *chainBuilder) Map(mapFunc interface{}) FunkBuilder {
	return &chainBuilder{Map(b.collection, mapFunc)}
}
func (b *chainBuilder) Reverse() FunkBuilder {
	return &chainBuilder{Reverse(b.collection)}
}
func (b *chainBuilder) Shuffle() FunkBuilder {
	return &chainBuilder{Shuffle(b.collection)}
}
func (b *chainBuilder) Uniq() FunkBuilder {
	return &chainBuilder{Uniq(b.collection)}
}

func (b *chainBuilder) All() bool {
	v := reflect.ValueOf(b.collection)
	c := make([]interface{}, v.Len())

	for i := 0; i < v.Len(); i++ {
		c[i] = v.Index(i).Interface()
	}
	return All(c...)
}
func (b *chainBuilder) Any() bool {
	v := reflect.ValueOf(b.collection)
	c := make([]interface{}, v.Len())

	for i := 0; i < v.Len(); i++ {
		c[i] = v.Index(i).Interface()
	}
	return Any(c...)
}
func (b *chainBuilder) Contains(elem interface{}) bool {
	return Contains(b.collection, elem)
}
func (b *chainBuilder) Every(elements ...interface{}) bool {
	return Every(b.collection, elements...)
}
func (b *chainBuilder) Find(predicate interface{}) interface{} {
	return Find(b.collection, predicate)
}
func (b *chainBuilder) Get(path string) interface{} {
	return Get(b.collection, path)
}
func (b *chainBuilder) Head() interface{} {
	return Head(b.collection)
}
func (b *chainBuilder) Keys() interface{} {
	return Keys(b.collection)
}
func (b *chainBuilder) In(v interface{}) bool {
	return b.Contains(v)
}
func (b *chainBuilder) IndexOf(elem interface{}) int {
	return IndexOf(b.collection, elem)
}
func (b *chainBuilder) IsEmpty() bool {
	return IsEmpty(b.collection)
}
func (b *chainBuilder) IsType(actual interface{}) bool {
	return IsType(b.collection, actual)
}
func (b *chainBuilder) Last() interface{} {
	return Last(b.collection)
}
func (b *chainBuilder) NotEmpty() bool {
	return NotEmpty(b.collection)
}
func (b *chainBuilder) Product() float64 {
	return Product(b.collection)
}
func (b *chainBuilder) Reduce(reduceFunc, acc interface{}) float64 {
	return Reduce(b.collection, reduceFunc, acc)
}
func (b *chainBuilder) Sum() float64 {
	return Sum(b.collection)
}
func (b *chainBuilder) Tail() interface{} {
	return Tail(b.collection)
}
func (b *chainBuilder) Type() reflect.Type {
	return reflect.TypeOf(b.collection)
}
func (b *chainBuilder) Value() interface{} {
	return b.collection
}
func (b *chainBuilder) Values() interface{} {
	return Values(b.collection)
}
