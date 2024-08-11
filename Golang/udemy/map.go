package main

import "fmt"

func Mapping() {
	// var colors map[string]string
	colors := map[string]string{
		"white": "#ffffff",
		"black": "#000000",
		"red":   "#4bf745",
	}
	// colors := make(map[string]string)
	// colors["white"] = "#ffffff"
	// colors := make(map[int]string)
	// colors[10] = "#ffffff"
	delete(colors, "red")
	PrintMapColor(colors)
}

//iteration of mapping
func PrintMapColor(color map[string]string) {
	for color, hex := range color {
		fmt.Println("color is:", color, "hex is:", hex)
	}
}
