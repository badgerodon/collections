package collections

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapViaBTree(t *testing.T) {
	m := newDictionaryViaBTree[int, int](func(a, b int) bool {
		return a < b
	})

	ordered := rand.Perm(50)
	sort.Ints(ordered)
	unordered := rand.Perm(50)

	for _, i := range unordered {
		m.Set(i, i)
	}

	s := NewSlice(m.Keys())
	assert.Equal(t, ordered, s)
}
