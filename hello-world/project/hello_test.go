package main

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	expected := "Hello, World"
	actual := Greeting()

	if expected != actual {
		t.Errorf("expected %q, actual %q", expected, actual)
	}
}
