/*An interface type is defined as a set of method signatures
*/

package main

import (
	"fmt"
	"math"
	)

type Abser interface {
	Abs() float64
}

func main() {
	var a1, a2, a3 Abser
	f := MyFloat(-math.Sqrt(2))
	v := Vertex{3, 4}

	a1 = f
	//a2 = v
	a3 = &v

	fmt.Println(a1, a2, a3)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} else { return float64(f)}

}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return (v.X*v.X + v.Y*v.Y)
}
