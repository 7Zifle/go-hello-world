package maps

import "testing"

func assertStrings(t *testing.T, want, got string) {
	t.Helper()
	if want != got {
		t.Errorf("Want: %q, got: %q", want, got)
	}
}

func assertError(t *testing.T, want, got error) {
	t.Helper()
	if want != got {
		t.Errorf("Want: %q, got: %q", want.Error(), got.Error())
	}
	if got == nil {
		if want == nil {
			return
		}
		t.Fatal("expected to get an error.")
	}
}

func assertDefinition(
	t *testing.T,
	dict Dictionary,
	word string,
	definition string) {

	t.Helper()
	got, err := dict.Search(word)

	if err != nil {
		t.Fatalf("Failed finding word err: %q", err.Error())
	}
	assertStrings(t, definition, got)
}

func TestAdd(t *testing.T) {
	t.Run("New word", func(t *testing.T) {
		var dictionary = Dictionary{}
		definition := "this is a test"
		word := "test"
		err := dictionary.Add(word, definition)
		assertError(t, nil, err)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("Existing word", func(t *testing.T) {
		definition := "this is a test"
		word := "test"
		var dictionary = Dictionary{word: definition}
		err := dictionary.Add(word, definition)
		assertError(t, ErrWordExists, err)
		assertDefinition(t, dictionary, word, definition)
	})
}
func TestDelete(t *testing.T) {
	word := "test"
	definition := "this is a test"
	dictionary := Dictionary{word: definition}
	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrWordNotFound {
		t.Errorf("%q was not removed", word)
	}
}

func TestUpdate(t *testing.T) {
	t.Run("Update existing word", func(t *testing.T) {
		word := "Test"
		dictionary := Dictionary{word: "this is a test"}
		newDefinition := "Just another test"

		err := dictionary.Update(word, newDefinition)
		assertError(t, nil, err)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("Update new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "Test"
		definition := "Just another test"

		err := dictionary.Update(word, definition)
		assertError(t, err, ErrUpdateWordNotFound)
	})
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}
	t.Run("Known word", func(t *testing.T) {
		want := "this is a test"
		got, _ := dictionary.Search("test")

		assertStrings(t, want, got)
	})
	t.Run("Unknown word", func(t *testing.T) {
		_, err := dictionary.Search("broken")

		assertError(t, ErrWordNotFound, err)
	})
}
