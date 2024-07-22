package points

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) String() string {
	return fmt.Sprintf("(%.2f, %.2f)", p.X, p.Y)
}

func (p Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func Abs(point Point) float64 {
	return math.Sqrt(point.X*point.X + point.Y*point.Y)
}

func (p Point) CompareTo(other Point) Point {
	if p.Abs() > other.Abs() {
		return p
	}
	return other
}
