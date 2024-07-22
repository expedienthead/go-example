package main

import (
	"fmt"
	"sort"

	"github.com/expedienthead/go-points"
)

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
	points := points.Magnitudes{
		points.Point{X: 3, Y: 4},
		points.Point{X: 5, Y: 10},
		points.Point{X: 1, Y: 2},
		points.PointXYZ{X: 1, Y: 2, Z: 3},
	}

	sort.Sort(points)

	// sort.Sort(points)
	// fmt.Println(points)
	for _, point := range points {
		fmt.Println(point)
		fmt.Println(point.Abs())
	}
}
