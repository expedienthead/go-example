package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
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

func main() {
	point := Point{10.5, 7.5}
	// point2 := Point{10.5, 8}
	// points := []Point{point, point2}
	// fmt.Println(points[0].X)
	point2 := Point{30, 40}
	// fmt.Println(point.Abs())
	fmt.Println(point.CompareTo(point2))
	fmt.Println(point == point2)
}
