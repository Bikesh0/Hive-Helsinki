package sprint

type Coords struct {
	X  int
	Y  int
}
type Rectangle struct {
	Width    int
	Height   int
	Area     int
	Perimeter int
}

func GetRectangle(min, max Coords) Rectangle {

	b := max.X - min.X
	l := max.Y - min.Y

	return Rectangle {

		Width: b,
		Height: l,
		Area: l * b,
		Perimeter: 2 * ( l + b),
	}

}