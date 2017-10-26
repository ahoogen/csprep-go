package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	a := New("I am a test")
	assert.NotNil(t, a)
	assert.NotNil(t, a.data)
	assert.Nil(t, a.next)

	s := a.data.(string)
	assert.Equal(t, "I am a test", s)
}

func TestPush(t *testing.T) {
	a := New(1)
	assert.NotNil(t, a)

	b := a.Push(2)
	assert.NotEqual(t, a, b)
	assert.NotNil(t, b.next)
}

func TestPop(t *testing.T) {
	a := New(1)
	assert.NotNil(t, a)
	assert.Nil(t, a.next)
	a = a.Push(2)
	assert.NotNil(t, a.next)
	a = a.Push(3)

	// b == Elem(3)
	// a == Elems(2, 1)
	b, a := a.Pop()
	assert.NotNil(t, b)
	assert.NotNil(t, a)
	assert.NotNil(t, a.next)
	assert.Nil(t, b.next)
	assert.Equal(t, 3, b.data.(int))
	assert.Equal(t, 2, a.data.(int))

	// b == Elem(2)
	// a == Elems(1, nil)
	b, a = a.Pop()
	assert.NotNil(t, b)
	assert.NotNil(t, a)
	assert.Nil(t, a.next)
	assert.Nil(t, b.next)
	assert.Equal(t, 2, b.data.(int))
	assert.Equal(t, 1, a.data.(int))

	// b == Elem(1)
	// a == nil
	b, a = a.Pop()
	assert.NotNil(t, b)
	assert.Nil(t, a)
	assert.Nil(t, b.next)
	assert.Equal(t, 1, b.data.(int))
}
