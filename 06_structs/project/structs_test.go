package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {

	t.Run("Perimiter of a Rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		expected := 40.0
		actual := rectangle.Perimeter()

		if actual != expected {
			t.Errorf("actual %g expected %g", actual, expected)
		}
	})
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		actual := shape.Area()
		if actual != expected {
			t.Errorf("actual %g expected %g", actual, expected)
		}
	}

	t.Run("Area of a Rectangle", func(t *testing.T) {
		rectangle := Rectangle{8.0, 4.0}
		expected := 32.0
		checkArea(t, rectangle, expected)
	})

	t.Run("Area of a Circle", func(t *testing.T) {
		circle := Circle{10}
		expected := 314.1592653589793
		checkArea(t, circle, expected)
	})

	areaTests := []struct {
		shape    Shape
		expected float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		actual := tt.shape.Area()
		if actual != tt.expected {
			t.Errorf("%#v: actual %.2f expected %.2f", tt.shape, actual, tt.expected)
		}
	}
}
