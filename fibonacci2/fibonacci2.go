//Implement a fibonacci function that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
// Implementation: fibonacci(n) will return a slice of integers containing the sequence of all consecutive "n" fibonacci numbers. "n int" is the argument of the function. Once n is bigger than 92, the fibinacci value will overflow the samll integer buffer and the results are no longer valid. Thus the need to implement it with the package "math/big". It includes a big.Int which we will use for the new fibonacci sequence.

package main

import "fmt"
import "math/big"

// Working with big.Int to compute the fibonacci(n int), one must pay attention to the methods from "math/big" package:
// var x Int      =====>  &x is an *Int of value 0
// var r = &Rat{} =====>  r is a *Rat of value 0
// y := new(Float)=====>  y is a *Float of value 0
// big.Int variables are pointers of big.Int type (*Int). So the fibonacci function should return a pointer of type big.Int. In our case the slice content will be pointers of type big.Int

func fibonacci(n int) []*big.Int {
	// Set the two first values of the sequence as pointers of type big.Int
	f := big.NewInt(0)
	g := big.NewInt(1)
	// Create a nil slice to hold the *Int values named fib_sequence.
	fib_sequence := []*big.Int{}
	for i := 0; i <= n; i++ {
		f.Add(f, g) // big.Int method for adding big.Int, since + is not defined by big.Int
		fib_sequence = append(fib_sequence, f)
		f, g = g, f
	}
	return fib_sequence
}

func main() {
	f := fibonacci(100)
	for i, f := range f {
		fmt.Printf("Fibonacci of %d is %d\n", i, f)
	}
}
