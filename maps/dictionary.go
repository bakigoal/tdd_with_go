package maps

type Dictionary map[string]string

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

const (
	ErrNotFound   = DictionaryError("could not find the word you were looking for")
	ErrWordExists = DictionaryError("cannot add word because it already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	record, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return record, nil
}

func (d Dictionary) Add(word, definition string) error {
	if _, ok := d[word]; ok {
		return ErrWordExists
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Update(word, newDefinition string) {
	d[word] = newDefinition
}
