package main

import (
	"testing"
)

func TestGreeting(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, expected, actual string) {
		t.Helper()
		if expected != actual {
			t.Errorf("expected %q, actual %q", expected, actual)
		}
	}

	t.Run("Hello, World with no input", func(t *testing.T) {
		expected := "Hello, World"
		actual := Greeting()
		assertCorrectMessage(t, expected, actual)
	})

	t.Run("Hello, World with an empty string name", func(t *testing.T) {
		expected := "Hello, World"
		actual := Greeting("")
		assertCorrectMessage(t, expected, actual)
	})

	t.Run("Hello, World with a provided name", func(t *testing.T) {
		expected := "Hello, Seth"
		actual := Greeting("Seth")
		assertCorrectMessage(t, expected, actual)
	})

}
