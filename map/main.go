package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	fmt.Println(colors)
	colors["white"] = "#ffff"
	fmt.Println(colors)
	delete(colors, "white")
	fmt.Println(colors)
	for key, value := range colors {
		fmt.Println("key = ", key, " value = ", value)
	}
}
