package collections

type collectionImpl[T any] struct {
	forEach func(callback func(T) bool)
	size    func() int
}

func (c collectionImpl[T]) ForEach(callback func(T) bool) {
	c.forEach(callback)
}

func (c collectionImpl[T]) Size() int {
	return c.size()
}

// Map maps all the values in a collection to a new collection.
func Map[TFrom, TTo any](c Collection[TFrom], f func(TFrom) TTo) Collection[TTo] {
	return collectionImpl[TTo]{
		forEach: func(callback func(TTo) bool) {
			c.ForEach(func(v TFrom) bool {
				return callback(f(v))
			})
		},
		size: func() int {
			return c.Size()
		},
	}
}
