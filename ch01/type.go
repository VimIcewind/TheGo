package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	var p Point
	fmt.Printf("p.X is %d, p.Y is %d\n", p.X, p.Y)
}
