package ch_19_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		myStackOfInts := new(Stack[int])
		// check stack is empty
		assert.True(t, myStackOfInts.IsEmpty())

		// add a thing, then check it's not empty
		myStackOfInts.Push(123)
		assert.False(t, myStackOfInts.IsEmpty())

		// add another thing, pop it back again
		myStackOfInts.Push(456)
		value, _ := myStackOfInts.Pop()
		assert.Equal(t, value, 456)
		value, _ = myStackOfInts.Pop()
		assert.Equal(t, value, 123)
		assert.True(t, myStackOfInts.IsEmpty())
	})

	t.Run("string stack", func(t *testing.T) {
		myStackOfStrings := new(Stack[string])

		// check stack is empty
		assert.True(t, myStackOfStrings.IsEmpty())

		// add a thing, then check it's not empty
		myStackOfStrings.Push("123")
		assert.False(t, myStackOfStrings.IsEmpty())

		// add another thing, pop it back again
		myStackOfStrings.Push("456")
		value, _ := myStackOfStrings.Pop()
		assert.Equal(t, value, "456")
		value, _ = myStackOfStrings.Pop()
		assert.Equal(t, value, "123")
		assert.True(t, myStackOfStrings.IsEmpty())
	})
}
