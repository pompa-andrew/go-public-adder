// Package public_adder it's a package comment!
package public_adder

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Float | constraints.Integer
}

// Add
//
// # Attention!
// Function... adds to number to each other!
//
// - A first element
// - A second element
//
// Jon said I have to add this [link]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
