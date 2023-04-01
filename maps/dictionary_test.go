package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{
		"test": "this is just a test",
	}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assert.Equal(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assert.Error(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)
		assert.NoError(t, err)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "key 1"
		oldDefinition := "value 1"
		newDefinition := "this is just a test"
		dictionary := Dictionary{word: oldDefinition}

		err := dictionary.Add(word, newDefinition)
		assert.Error(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, oldDefinition)
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{word: definition}
	newDefinition := "new definition"

	dictionary.Update(word, newDefinition)
	assertDefinition(t, dictionary, word, newDefinition)
}

func assertDefinition(t *testing.T, dictionary Dictionary, word string, definition string) {
	got, err := dictionary.Search(word)
	assert.NoError(t, err)
	assert.Equal(t, got, definition)
}
