package structs

import "testing"

func TestPerimeter(t *testing.T) {

	rect := Rectangle{10.0, 10.0}
	got := Perimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}

}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rect := Rectangle{10.0, 10.0}
		want := 100.0

		checkArea(t, rect, want)
	})

	t.Run("circles", func(t *testing.T) {
		circ := Circle{10.0}
		want := 314.1592653589793

		checkArea(t, circ, want)
	})

}

// Table based tests to make your assertions clearer
// and your suites easier to extend & maintain
func TestAreaTableDriven(t *testing.T) {

	type TestData struct {
		shape   Shape
		hasArea float64
	}

	// array of struct
	areaTests := []TestData{
		TestData{shape: Rectangle{10.0, 10.0}, hasArea: 100.0},
		TestData{shape: Circle{10.0}, hasArea: 314.1592653589793},
		TestData{shape: Triangle{12, 6}, hasArea: 36.0},
	}

	// alt syntax, array of anonymous struct
	/*
		areaTests := []struct {
			shape Shape
			hasArea  float64
		}{
			{ Rectangle{10.0, 10.0}, 100.0 },
			{Circle{10.0}, 314.1592653589793 },
		}
	*/

	for _, test := range areaTests {

		got := test.shape.Area()

		if got != test.hasArea {
			// print out our struct with the values in its fields
			t.Errorf("%#v got %.2f want %.2f", test, got, test.hasArea)
		}
	}

}
