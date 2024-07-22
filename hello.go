package main

import (
	"fmt"
	"math"
	"sort"
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

type Points []Point

func (points Points) Len() int {
	return len(points)
}

func (points Points) Less(i, j int) bool {
	return points[i].Abs() < points[j].Abs()
}

func (points Points) Swap(i, j int) {
	points[i], points[j] = points[j], points[i]
}

func (points Points) String() string {
	str := ""
	for _, point := range points {
		str += fmt.Sprintln(point.String(), "\t", fmt.Sprintf("%.2f", point.Abs()))
		// str += point.String() + "\t" + fmt.Sprintf("%f", point.Abs()) + "\n"
	}
	return str
}

func main() {
	// point := Point{10.5, 7.5}
	// point2 := Point{10.5, 8}
	// points := []Point{point, point2}
	// fmt.Println(points[0].X)
	// point2 := Point{30, 40}
	// fmt.Println(point.Abs())
	// fmt.Println(point.CompareTo(point2))
	// fmt.Println(point == point2)

	// point := Point{3, 4}
	// fmt.Println(point)
	points := Points{
		{3, 4},
		{5, 10},
		{1, 2},
	}
	sort.Sort(points)
	fmt.Println(points)
}
