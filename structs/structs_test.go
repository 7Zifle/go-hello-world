package structs

import "testing"

func TestParimeter(t *testing.T) {
	rectangle := Rectangle{2, 2}
	want := float64(8)
	got := Parimeter(rectangle)

	if want != got {
		t.Errorf("want: %.2f got: %.2f", want, got)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 2, Height: 2}, want: 4},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v want: %g got: %g", tt.shape, tt.want, got)
			}
		})
	}
}
