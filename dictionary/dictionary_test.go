package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "test string"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "test string"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("foo")
		want := ErrNotFound

		if err == nil {
			t.Fatal("expected failure")
		}

		assertError(t, err, want)
	})
	
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		value := "test string"
		err := dictionary.Add("test", value)

		assertError(t, err, nil)
		assertInDictionary(t, dictionary, "test", value)
	})

	t.Run("existing word", func(t * testing.T) {
		word := "test"
		value := "test string"
		dictionary := Dictionary{word: value}

		err := dictionary.Add(word, "new test string")
		assertError(t, err, ErrWordExists)
		assertInDictionary(t, dictionary, word, value)
	})

}

func assertInDictionary(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}