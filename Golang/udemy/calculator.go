package main

import "fmt"

type TypeCalculator interface {
	Penjumlahan() int
	Pembagian() int
	Perkalian() int
	Pengurangan() int
}

type InputOfTypeCalculator struct {
	x int
	y int
}

func (HandleInput InputOfTypeCalculator) Penjumlahan() int {
	return HandleInput.x + HandleInput.y
}

func (HandleInput InputOfTypeCalculator) Pengurangan() int {
	return HandleInput.x - HandleInput.y
}

func (HandleInput InputOfTypeCalculator) Perkalian() int {
	return HandleInput.x * HandleInput.y
}

func (HandleInput InputOfTypeCalculator) Pembagian() int {
	return HandleInput.x / HandleInput.y
}

func PrintPenjumlahan(HandleTypeCalculator TypeCalculator) {
	fmt.Println(HandleTypeCalculator.Penjumlahan())
}
