package sprint

type Point struct {
	X  float32
	Y  float32
	Text string

}

func PointDiff(p1, p2 Point) Point {
	if p1.X + p1.Y > p2.X + p2.Y {
		return p1
	}
	return p2


}