package maps

type Dictionary map[string]string

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

const (
	ErrNotFound          = DictionaryError("could not find the word you were looking for")
	ErrWordExists        = DictionaryError("cannot add word because it already exists")
	ErrWordDoesNotExists = DictionaryError("cannot update word because it doesn't exists")
)

func (d Dictionary) Search(word string) (string, error) {
	record, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return record, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
		return nil
	case nil:
		return ErrWordExists
	default:
		return err
	}
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		d[word] = newDefinition
		return nil
	default:
		return err
	}
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
