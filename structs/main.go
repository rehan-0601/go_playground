package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	// just contactInfo means conatactInfo: contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}
	//alex := person{firstName: "Alex", lastName: "Anderson"}
	// go assigns zero value below
	// string "", int 0, float 0, bool false
	// firstName and lastName will be empty strings
	var alex person
	fmt.Println(alex)
	// prints all fields and their values in alex
	fmt.Printf("%+v", alex)

	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000},
	}
	// every line needs to end in a comma ^ ^

	// jimPointer := &jim
	// jimPointer.updateName("jimmy")

	// below is go's shortcut. you can use jim OR jimPointer
	jim.updateName("jimmy")
	jim.print()
}

// this will not update the actual person. PASS BY VALUE
// need a pointer if you need to update the actual person
// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// Go pointer gotchas
// with a slice, assignment will modify the original slice
// array : primitive data structure, cannot be resized, rarely used directly
// slice : can grow and shrink, internally - ptr to array head, cap, length
// 		^^ slice is one data structure, which points to the arrray data structure.
//		local copy also contains the ptr pointing to same array.

// reference type: slice, map, channels, pointers, functions
// value types: int, float, string, bool, structs
