package ch_07_maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{
		"test": "this is just a test",
	}

	t.Run("known word", func(t *testing.T) {
		assertDefinition(t, dictionary, "test", "this is just a test")
	})
	t.Run("unknown word", func(t *testing.T) {
		assertNoDefinition(t, dictionary, "unknown word")
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
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)
		assert.NoError(t, err)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("new word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)
		assert.Error(t, err, ErrWordDoesNotExists)
		assertNoDefinition(t, dictionary, word)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "just a test definition"}

	dictionary.Delete(word)

	assertNoDefinition(t, dictionary, word)
}

func assertDefinition(t *testing.T, dictionary Dictionary, word string, definition string) {
	got, err := dictionary.Search(word)
	assert.NoError(t, err)
	assert.Equal(t, got, definition)
}

func assertNoDefinition(t *testing.T, dictionary Dictionary, word string) {
	_, err := dictionary.Search(word)
	assert.Error(t, err, ErrNotFound)
}
