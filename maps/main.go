package main

import "fmt"

func main() {
	// Literal syntax
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#0000ff",
		"white": "#ffffff",
	}
	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color)
		fmt.Println(hex)
	}
}
