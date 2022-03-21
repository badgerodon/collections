package collections

type (
	Collection[T any] interface {
		ForEach(callback func(T) bool)
		Size() int
	}

	// A Deque is a double-ended queue.
	Deque[T any] interface {
		Clear()
		PeekBack() (value T, ok bool)
		PeekFront() (value T, ok bool)
		PopBack() (value T, ok bool)
		PopFront() (value T, ok bool)
		PushBack(value T)
		PushFront(value T)
		Size() int
	}

	// A Dictionary is a collection of key value pairs.
	Dictionary[TKey, TValue any] interface {
		Clear()
		Delete(key TKey)
		ForEach(callback func(Pair[TKey, TValue]) bool)
		Get(key TKey) (value TValue, ok bool)
		Keys() Collection[TKey]
		Set(key TKey, value TValue)
		Size() int
		Values() Collection[TValue]
	}

	// A Queue is a collection that supports FIFO operations.
	Queue[T any] interface {
		Clear()
		Peek() (value T, ok bool)
		Pop() (value T, ok bool)
		Push(value T)
		Size() int
	}

	// A Set is a unique collection of values.
	Set[T any] interface {
		Add(value T)
		Clear()
		Delete(value T)
		ForEach(callback func(T) bool)
		Has(value T) bool
		Size() int
	}

	// A Stack is a collection that supports LIFO operations.
	Stack[T any] interface {
		Clear()
		Peek() (value T, ok bool)
		Pop() (value T, ok bool)
		Push(value T)
		Size() int
	}
)

// NewDeque creates a new Deque implemented using a slice as a ring buffer.
func NewDeque[T any]() Deque[T] {
	return newDequeViaRingSlice[T]()
}

// NewDictionary creates a new Dictionary implemented using a map.
func NewDictionary[TKey comparable, TValue any]() Dictionary[TKey, TValue] {
	return newDictionaryViaMap[TKey, TValue]()
}

// NewQueue creates a new Queue implemented using a slice.
func NewQueue[T any]() Queue[T] {
	return newQueueViaSlice[T]()
}

// NewSet creates a new Set implemented using a map.
func NewSet[T comparable]() Set[T] {
	return newSetViaMap[T]()
}

// NewSlice creates a new slice from a collection.
func NewSlice[T any](collection Collection[T]) []T {
	arr := make([]T, 0, collection.Size())
	collection.ForEach(func(item T) bool {
		arr = append(arr, item)
		return true
	})
	return arr
}

// NewSortedDictionary creates a new Dictionary implemented using a btree, providing sorted iteration by the less function.
func NewSortedDictionary[TKey, TValue any](less func(TKey, TKey) bool) Dictionary[TKey, TValue] {
	return newDictionaryViaBTree[TKey, TValue](less)
}

// NewSortedSet creates a new Set implemented using a btree, providing sorted iteration by the less function.
func NewSortedSet[T any](less func(T, T) bool) Set[T] {
	return newSetViaBTree(less)
}

// NewStack creates a new Stack implemented using a slice.
func NewStack[T any]() Stack[T] {
	return newStackViaSlice[T]()
}

type integer interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64
}

func mod[T integer](x, y T) T {
	return (x%y + y) % y
}
