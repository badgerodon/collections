package collections

type dictionaryViaMap[TKey comparable, TValue any] struct {
	m map[TKey]TValue
}

var _ interface{ Dictionary[int, int] } = (*dictionaryViaMap[int, int])(nil)

func newDictionaryViaMap[TKey comparable, TValue any]() *dictionaryViaMap[TKey, TValue] {
	return &dictionaryViaMap[TKey, TValue]{
		m: make(map[TKey]TValue),
	}
}

func (d *dictionaryViaMap[TKey, TValue]) Clear() {
	d.m = make(map[TKey]TValue)
}

func (d *dictionaryViaMap[TKey, TValue]) Delete(key TKey) {
	delete(d.m, key)
}

func (d *dictionaryViaMap[TKey, TValue]) ForEach(callback func(entry Pair[TKey, TValue]) bool) {
	for k, v := range d.m {
		if !callback(NewPair(k, v)) {
			return
		}
	}
}

func (d *dictionaryViaMap[TKey, TValue]) Get(key TKey) (value TValue, ok bool) {
	value, ok = d.m[key]
	return value, ok
}

func (d *dictionaryViaMap[TKey, TValue]) Keys() Collection[TKey] {
	return Map[Pair[TKey, TValue]](d, func(pair Pair[TKey, TValue]) TKey {
		return pair.First
	})
}

func (d *dictionaryViaMap[TKey, TValue]) Set(key TKey, value TValue) {
	d.m[key] = value
}

func (d *dictionaryViaMap[TKey, TValue]) Size() int {
	return len(d.m)
}

func (d *dictionaryViaMap[TKey, TValue]) Values() Collection[TValue] {
	return Map[Pair[TKey, TValue]](d, func(pair Pair[TKey, TValue]) TValue {
		return pair.Second
	})
}

type setViaMap[T comparable] struct {
	m map[T]struct{}
}

var _ interface{ Set[int] } = (*setViaMap[int])(nil)

func newSetViaMap[T comparable]() *setViaMap[T] {
	return &setViaMap[T]{m: make(map[T]struct{})}
}

func (s *setViaMap[T]) Add(value T) {
	s.m[value] = struct{}{}
}

func (s *setViaMap[T]) Clear() {
	s.m = make(map[T]struct{})
}

func (s *setViaMap[T]) Delete(value T) {
	delete(s.m, value)
}

func (s *setViaMap[T]) ForEach(callback func(T) bool) {
	for v := range s.m {
		if !callback(v) {
			break
		}
	}
}

func (s *setViaMap[T]) Has(value T) bool {
	_, ok := s.m[value]
	return ok
}

func (s *setViaMap[T]) Size() int {
	return len(s.m)
}
