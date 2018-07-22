package main

func main() {}

//Search takes a dictionary and a key and returns the value
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

// Dictionary is a wrapper around map
type Dictionary map[string]string

// DictionaryErr is a template for various dictionary errors
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	// ErrNotFound is thrown, when a word is not found in the dictionary
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	// ErrWordExists is thrown, when adding an existing word to the dictionary
	ErrWordExists = DictionaryErr("cannot add word, because it already exists")
)

// Search takes a key and returns the corresponding value from the dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add adds a word and a definition to the dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
