package main

import "fmt"

type Shape interface {
	GetArea() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Square struct {
	SideLength float64
}

func (s Square) GetArea() float64 {
	return s.SideLength * s.SideLength
}

func (t Triangle) GetArea() float64 {
	return 0.5 * t.Base * t.Height
}

func PrintArea(s Shape) {
	fmt.Println(s.GetArea())
}
