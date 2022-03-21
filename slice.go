package collections

type dequeViaRingSlice[T any] struct {
	v []T

	front, back, length int
}

var _ interface{ Deque[int] } = (*dequeViaRingSlice[int])(nil)

func newDequeViaRingSlice[T any]() *dequeViaRingSlice[T] {
	return &dequeViaRingSlice[T]{}
}

func (d *dequeViaRingSlice[T]) Clear() {
	d.v = nil
	d.front, d.back, d.length = 0, 0, 0
}

func (d *dequeViaRingSlice[T]) PeekBack() (value T, ok bool) {
	if d.length == 0 {
		return value, false
	}

	return d.v[mod(d.back-1, len(d.v))], true
}

func (d *dequeViaRingSlice[T]) PeekFront() (value T, ok bool) {
	if d.length == 0 {
		return value, false
	}

	return d.v[d.front], true
}

func (d *dequeViaRingSlice[T]) PopBack() (value T, ok bool) {
	value, ok = d.PeekBack()
	if !ok {
		return value, ok
	}

	d.back = mod(d.back-1, len(d.v))
	d.length--
	return value, true
}

func (d *dequeViaRingSlice[T]) PopFront() (value T, ok bool) {
	value, ok = d.PeekFront()
	if !ok {
		return value, ok
	}

	d.front = mod(d.front+1, len(d.v))
	d.length--
	return value, true
}

func (d *dequeViaRingSlice[T]) PushBack(value T) {
	d.maybeGrow()
	d.v[d.back] = value
	d.back = mod(d.back+1, len(d.v))
	d.length++
}

func (d *dequeViaRingSlice[T]) PushFront(value T) {
	d.maybeGrow()
	d.front = mod(d.front-1, len(d.v))
	d.v[d.front] = value
	d.length++
}

func (d *dequeViaRingSlice[T]) Size() int {
	return d.length
}

func (d *dequeViaRingSlice[T]) isEmpty() bool {
	return d.length == 0
}

func (d *dequeViaRingSlice[T]) isFull() bool {
	return d.length == len(d.v)
}

func (d *dequeViaRingSlice[T]) isSparse() bool {
	return 1 < d.length && d.length < len(d.v)/4
}

func (d *dequeViaRingSlice[T]) maybeGrow() {
	if d.length == 0 {
		d.resize(1)
		return
	}

	if d.isFull() {
		d.resize(len(d.v) * 2)
	}
}

func (d *dequeViaRingSlice[T]) maybeShrink() {
	if d.isSparse() {
		d.resize(len(d.v) / 2)
	}
}

func (d *dequeViaRingSlice[T]) resize(size int) {
	v := make([]T, size)
	for i := 0; i < d.length && i < size; i++ {
		v[i] = d.v[(d.front+i)%len(d.v)]
	}
	d.v = v
	d.front = 0
	d.back = d.length
}

type queueViaSlice[T any] struct {
	v []T
}

var _ interface{ Queue[int] } = (*queueViaSlice[int])(nil)

func newQueueViaSlice[T any]() *queueViaSlice[T] {
	return &queueViaSlice[T]{}
}

func (q *queueViaSlice[T]) Clear() {
	q.v = nil
}

func (q *queueViaSlice[T]) Peek() (value T, ok bool) {
	if len(q.v) > 0 {
		return q.v[0], true
	}
	return value, false
}

func (q *queueViaSlice[T]) Pop() (value T, ok bool) {
	if len(q.v) > 0 {
		value, q.v = q.v[0], q.v[1:]
		return value, true
	}
	return value, false
}

func (q *queueViaSlice[T]) Push(value T) {
	q.v = append(q.v, value)
}

func (q *queueViaSlice[T]) Size() int {
	return len(q.v)
}

type stackViaSlice[T any] struct {
	v []T
}

var _ interface{ Stack[int] } = (*stackViaSlice[int])(nil)

func newStackViaSlice[T any]() *stackViaSlice[T] {
	return &stackViaSlice[T]{}
}

func (q *stackViaSlice[T]) Clear() {
	q.v = nil
}

func (q *stackViaSlice[T]) Peek() (value T, ok bool) {
	if len(q.v) > 0 {
		return q.v[len(q.v)-1], true
	}
	return value, false
}

func (q *stackViaSlice[T]) Pop() (value T, ok bool) {
	if len(q.v) > 0 {
		value, q.v = q.v[len(q.v)-1], q.v[:len(q.v)-1]
		return value, true
	}
	return value, false
}

func (q *stackViaSlice[T]) Push(value T) {
	q.v = append(q.v, value)
}

func (q *stackViaSlice[T]) Size() int {
	return len(q.v)
}
