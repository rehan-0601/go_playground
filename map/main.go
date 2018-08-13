package main

import (
	"fmt"
)

func main() {
	// in a map all keys should be of same type.
	//All values should be of same type

	// var colors map[string]string
	// colors := make(map[string]string)
	colors := map[string]string{
		"red":   "adcsd",
		"blue":  "vsdvas",
		"white": "#ffff",
	}

	// colors["bla"] = "acc"
	// delete(colors, bla)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, "is", hex)
	}
}

// struct vs map
// cannot iterate over all keys in the struct
// struct is value type. map is reference type
// all keys in struct should be known at compile time
// map - keys are indexed, can be iterated over
