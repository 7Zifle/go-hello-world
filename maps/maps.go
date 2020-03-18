package maps

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrUpdateWordNotFound = DictionaryErr("Unable to find word in update")
	ErrWordNotFound       = DictionaryErr("Unknown word")
	ErrWordExists         = DictionaryErr("Word exists")
)

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrWordNotFound:
		return ErrUpdateWordNotFound
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Search(key string) (string, error) {
	value, okay := d[key]
	if !okay {
		return "", ErrWordNotFound
	}
	return value, nil
}
