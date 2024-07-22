package points

import (
	"fmt"
	"math"
)

type PointXYZ struct {
	X, Y, Z float64
}

func (point PointXYZ) Abs() float64 {
	return math.Sqrt(point.X*point.X + point.Y*point.Y + point.Z*point.Z)
}

func (point PointXYZ) String() string {
	return fmt.Sprintf("(%.2f, %.2f, %.2f) -> %f", point.X, point.Y, point.Z, point.Abs())
}
