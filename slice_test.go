package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDequeViaRingSlice(t *testing.T) {
	d := newDequeViaRingSlice[int]()
	d.PushFront(1)
	d.PushBack(2)
	assert.Equal(t, 2, d.Size())
	v, ok := d.PopBack()
	assert.True(t, ok)
	assert.Equal(t, 2, v)
	v, ok = d.PopBack()
	assert.True(t, ok)
	assert.Equal(t, 1, v)
	v, ok = d.PopBack()
	assert.False(t, ok)
	assert.Equal(t, 0, v)
	d.PushBack(3)
	v, ok = d.PopFront()
	assert.True(t, ok)
	assert.Equal(t, 3, v)
	assert.Equal(t, 0, d.Size())

	for i := 0; i < 1000; i++ {
		d.PushBack(i)
	}
	for i := 0; i < 1000; i++ {
		_, ok := d.PopFront()
		assert.True(t, ok)
	}
}
