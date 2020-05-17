package main

import "fmt"

type shape interface {
	getArea() float64
}
type triangle struct{}
type square struct{}

func main() {
	t := triangle{}
	s := square{}

	printArea(t)
	printArea(s)
}

func (t triangle) getArea() float64 {
	return 0.4
}

func (s square) getArea() float64 {
	return 0.53
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
