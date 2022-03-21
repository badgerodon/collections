package collections

import "github.com/tidwall/btree"

type dictionaryViaBTree[TKey, TValue any] struct {
	b    *btree.Generic[Pair[TKey, TValue]]
	less func(TKey, TKey) bool
}

var _ interface{ Dictionary[int, int] } = (*dictionaryViaBTree[int, int])(nil)

func newDictionaryViaBTree[TKey, TValue any](less func(TKey, TKey) bool) *dictionaryViaBTree[TKey, TValue] {
	return &dictionaryViaBTree[TKey, TValue]{
		b: btree.NewGeneric(func(a, b Pair[TKey, TValue]) bool {
			return less(a.First, b.First)
		}),
		less: less,
	}
}

func (d *dictionaryViaBTree[TKey, TValue]) Clear() {
	less := d.less
	d.b = btree.NewGeneric(func(a, b Pair[TKey, TValue]) bool {
		return less(a.First, b.First)
	})
	d.less = less
}

func (d *dictionaryViaBTree[TKey, TValue]) Delete(key TKey) {
	var zero TValue
	d.b.Delete(NewPair(key, zero))
}

func (d *dictionaryViaBTree[TKey, TValue]) ForEach(callback func(Pair[TKey, TValue]) bool) {
	d.b.Scan(func(item Pair[TKey, TValue]) bool {
		return callback(item)
	})
}

func (d *dictionaryViaBTree[TKey, TValue]) Get(key TKey) (value TValue, ok bool) {
	item, ok := d.b.Get(NewPair(key, value))
	if !ok {
		return value, ok
	}

	return item.Second, true
}

func (d *dictionaryViaBTree[TKey, TValue]) Keys() Collection[TKey] {
	return Map[Pair[TKey, TValue]](d, func(pair Pair[TKey, TValue]) TKey {
		return pair.First
	})
}

func (d *dictionaryViaBTree[TKey, TValue]) Set(key TKey, value TValue) {
	d.b.Set(NewPair(key, value))
}

func (d *dictionaryViaBTree[TKey, TValue]) Size() int {
	return d.b.Len()
}

func (d *dictionaryViaBTree[TKey, TValue]) Values() Collection[TValue] {
	return Map[Pair[TKey, TValue]](d, func(pair Pair[TKey, TValue]) TValue {
		return pair.Second
	})
}

type setViaBTree[T any] struct {
	b    *btree.Generic[T]
	less func(T, T) bool
}

var _ interface{ Set[int] } = (*setViaBTree[int])(nil)

func newSetViaBTree[T any](less func(T, T) bool) *setViaBTree[T] {
	return &setViaBTree[T]{
		b:    btree.NewGeneric(less),
		less: less,
	}
}

func (s *setViaBTree[T]) Add(value T) {
	s.b.Set(value)
}

func (s *setViaBTree[T]) Clear() {
	s.b = btree.NewGeneric(s.less)
}

func (s *setViaBTree[T]) Delete(value T) {
	s.b.Delete(value)
}

func (s *setViaBTree[T]) ForEach(callback func(T) bool) {
	s.b.Scan(callback)
}

func (s *setViaBTree[T]) Has(value T) bool {
	_, ok := s.b.Get(value)
	return ok
}

func (s *setViaBTree[T]) Size() int {
	return s.b.Len()
}
