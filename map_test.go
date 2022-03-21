package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapSet(t *testing.T) {
	s := newSetViaMap[int]()
	s.Add(1)
	assert.True(t, s.Has(1))
	assert.Equal(t, 1, s.Size())
	s.Delete(1)
	assert.False(t, s.Has(1))
	assert.Equal(t, 0, s.Size())
}
