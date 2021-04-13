package main

import (
	"fmt"
	"math"
)

type Coord struct {
	x, y int
}

func midpoint(p1, p2 Coord) Coord {
	x := (p2.x + p1.x) / 2
	y := (p2.y + p1.y) / 2
	return Coord{x, y}
}

func dist(p1, p2 Coord) float64 {
	return math.Sqrt(math.Pow(float64(p2.y-p1.y), 2) + math.Pow(float64(p2.x-p1.x), 2))
}

func main() {
	coord1 := Coord{10, 15}
	coord2 := Coord{-2, 22}
	fmt.Println(midpoint(coord1, coord2))
	fmt.Println(dist(coord1, coord2))
}
