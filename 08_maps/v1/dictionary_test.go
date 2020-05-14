package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		actual, _ := dictionary.Search("test")
		expected := dictionary["test"]

		assertStrings(t, actual, expected)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, actual := dictionary.Search("unknown")
		expected := ErrNotFound

		assertErrors(t, actual, expected)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	t.Run("Add to dictionary", func(t *testing.T) {
		key := "test"
		value := "this is just a test"

		dictionary.Add(key, value)

		assertDefinition(t, dictionary, key, value)
	})
}

func assertStrings(t *testing.T, actual, expected string) {
	t.Helper()

	if actual != expected {
		t.Errorf("actual %q expected %q", actual, expected)
	}
}

func assertErrors(t *testing.T, actual, expected error) {
	t.Helper()

	if actual != expected {
		t.Errorf("actual %q expected error %q", actual, expected)
	}
}

func assertDefinition(t *testing.T, d Dictionary, key, value string) {
	t.Helper()

	actual, err := d.Search("test")
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	if actual != value {
		t.Errorf("actual %q expected %q", actual, value)
	}
}
