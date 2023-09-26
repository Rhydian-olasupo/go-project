package main

import "testing"

///func TestPerimeter(t *testing.T) {
//rectangle := Rectangle{10.0, 10.0}
//got := Perimeter(rectangle)
//want := 40.0

//if got != want {
//t.Errorf("got %.2f want %.2f", got, want)
//}
//}

/*func TestArea(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		got := rectangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}

	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := 314.2

		if got != want {
			t.Errorf("got %.2f want %.2f ", got, want)
		}
	})

}*/

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	checkPerimter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %g want%g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
		checkPerimter(t, rectangle, 36)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
		checkPerimter(t, circle, 62.83185307179586)
	})

}
