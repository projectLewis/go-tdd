package main

import (
	"testing"
)

func TestAdder(t *testing.T) {

	assert := func(t *testing.T, expected, actual int) {
		t.Helper()
		if expected != actual {
			t.Errorf("expected %v, actual %v", expected, actual)
		}
	}

	t.Run("Adds two numbers", func(t *testing.T) {
		expected := 4
		actual := Add(2, 2)

		assert(t, expected, actual)
	})

	t.Run("Adds three numbers", func(t *testing.T) {
		expected := 10
		actual := Add(2, 3, 5)

		assert(t, expected, actual)
	})
}
