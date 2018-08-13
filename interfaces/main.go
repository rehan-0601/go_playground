package main

import (
	"fmt"
)

type englishBot struct{}
type spanishBot struct{}

// cant create a value from this type.
// cant initialize a var of this type
type bot interface {
	getGreeting() string
}

// englishBot and spanishBot implement getGreeting, therefore
// can be treated as type bot

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	// printGreeting(eb)
	// printGreeting(sb)
}

// very custom implementation
// **can omit eb if not used inside function
func (eb englishBot) getGreeting() string {
	return "Hi there"
}

// very custom implementation
func (sb spanishBot) getGreeting() string {
	return "Hola"
}

// printGreeting looks like it can be re-used
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// no function overloading in GO. cannot have the same func name
// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
//}

// interface
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
