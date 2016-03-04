//Implement the square root function using Newton's method.
//In this case, Newton's method is to approximate Sqrt(x) 
//by picking a starting point z and then repeating:
// z = z - ((z*z - x) / (2.0 * x))
// return a non-nil error value when given a negative
//number, as it doesn't support complex numbers.
package main

import (
	"fmt"
	"math"
)
// Create a new type
//type ErrNegativeSqrt float64
type ErrNegativeSqrt float64

//and make it an error by giving it a
//func (e ErrNegativeSqrt) Error() string
//method such that ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative number: -2". 
func (e ErrNegativeSqrt) Error() string {
// implements error interface method Error()
	return fmt.Sprintf("cannot Sqrt negative number: %#v", e) 
}

//Sqrt should return a non-nil error value when given a negative
//number, as it doesn't support complex numbers.
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0.0, ErrNegativeSqrt(x)
	}
	z := 1.0
	distance := 1.0
	const delta = 1e-10
	for distance > delta {
		current := z - ((z*z - x) / (2.0 * x))
		distance = math.Abs(z - current)
		z = current
	}
	return z, nil

}

func main() {
	fmt.Println(Sqrt(-4))
	fmt.Println(Sqrt(4))
}
