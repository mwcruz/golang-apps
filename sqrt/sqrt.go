//Implement the square root function using Newton's method.

//In this case, Newton's method is to approximate Sqrt(x) by picking a starting point z and then repeating:

// z = z - ((z*z - x) / (2.0 * x))

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	distance := 1.0
	const delta = 1e-10
	for distance > delta {
		current := z - ((z*z - x) / (2.0 * x))
		distance = math.Abs(z - current)
		z = current
	}
	return z

}

func main() {
	fmt.Println(Sqrt(24))
}
