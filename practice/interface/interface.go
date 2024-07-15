package main

import "fmt"

type coordinates interface {
	xaxis() int
	yaxis() int
}

type point2D struct {
	X int
	Y int
}

func (s point2D) xaxis() int {
	return s.X
}

func (s point2D) yaxis() int {
	return s.Y
}

func findCoordinates(a coordinates) {
	fmt.Println("X:", a.xaxis(), "Y:", a.yaxis())
}

type xpoint int

func (s xpoint) xaxis() int {
	return int(s)
}

func (s xpoint) yaxis() int {
	return 0
}

func main() {
	x := point2D{X: -1, Y: 12}
	fmt.Println(x)
	findCoordinates(x)

	y := xpoint(10)
	findCoordinates(y)
}
