package main

import "testing"

func TestDelete(t *testing.T) {
	t.Run("Deletes existing word", func(t *testing.T) {
		dictionary := Dictionary{}
		_ = dictionary.Add("key", "value")
		err := dictionary.Delete("key")
		assertNoError(t, err)
		_, err = dictionary.Search("value")
		if err == nil {
			t.Errorf("Found word when it should've been deleted")
		}

	})
	t.Run("Tries to delete word that does not exist", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Delete("key")
		assertError(t, err, ErrKeyDoesNotExist)
	})
}

// update should only work if a value already exists
func TestUpdate(t *testing.T) {
	t.Run("Updates existing word", func(t *testing.T) {
		dictionary := Dictionary{}
		_ = dictionary.Add("key", "value")
		err := dictionary.Update("key", "anothervalue")
		assertNoError(t, err)
	})

	t.Run("Tries to update non-existing word", func(t *testing.T) {})
	dictionary := Dictionary{}
	err := dictionary.Update("key", "value")
	assertError(t, err, ErrKeyDoesNotExist)
}

// Add should not modify existing values*/
func TestAdd(t *testing.T) {
	t.Run("Adds new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("key", "value")
		assertNoError(t, err)
		got, _ := dictionary.Search("key")
		assertStrings(t, got, "value")
	})

	t.Run("Tries to add existing word", func(t *testing.T) {
		dictionary := Dictionary{}
		_ = dictionary.Add("key", "value")
		err := dictionary.Add("key", "value")
		assertError(t, err, ErrKeyAlreadyExists)
	})

}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("word that isn't there")
		assertError(t, err, ErrNotFound)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
