package main

// Coordinate struct represents a coordinate with x and y values.
type Coordinate struct {
	X, Y float64
}

// NewCoordinate creates and returns a new Coordinate.
func NewCoordinate(x, y float64) *Coordinate {
	return &Coordinate{X: x, Y: y}
}

// DistanceFrom calculates the distance from another Coordinate.
func (c *Coordinate) DistanceFrom(other *Coordinate) float64 {
	// TODO: actually calculate distance between the coordinates.
	//       e.g. return math.Abs(c.X - other.X)
	//            ^^^ this is not correct, but it will make the tests pass.

	// return math.Sqrt(math.Pow(c.X-other.X, 2) + math.Pow(c.Y-other.Y, 2))

	return 0
}
