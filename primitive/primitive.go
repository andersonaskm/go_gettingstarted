package main

import "fmt"

/*
Declaration and primitives
Pointers
Constants
*/
func main() {

	fmt.Println("------------------------------")
	fmt.Println("Primitive Data Types")
	fmt.Println("------------------------------")

	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Anderson"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	// complex numbers
	c := complex(3, 4)
	fmt.Println(c)

	real, imaginaria := real(c), imag(c)
	fmt.Println(real, imaginaria)

	pointers()
	constants()
}

/*
pointers
*/
func pointers() {

	fmt.Println("------------------------------")
	fmt.Println("Pointers")
	fmt.Println("------------------------------")

	var firstName *string = new(string)

	*firstName = "Anderson Szalai"

	fmt.Println("Pointer address: ", firstName)
	fmt.Println("Pointer value: ", *firstName)

	lastName := "Szalai"
	fmt.Println(lastName)

	ptr := &lastName // pointer to address of lastName
	fmt.Println(ptr, *ptr)

	lastName = "Szalai Motta"
	fmt.Println(ptr, *ptr)
}

const (
	male   = 1
	female = 2
)

const (
	inactive = iota
	active   = iota
)

const (
	first = iota + 5
	second
)

const (
	third = iota
)

/*
Constants
*/
func constants() {
	fmt.Println("------------------------------")
	fmt.Println("Constants")
	fmt.Println("------------------------------")

	const pi = 3.1415
	fmt.Println(pi)

	fmt.Println("Constants: ", male, female)

	// iota
	fmt.Println("Iota: ", inactive, active)
	fmt.Println("New Iota: ", third)

	// constants expressions
	fmt.Println("Iota Constant Expressions: ", first, second)

}
