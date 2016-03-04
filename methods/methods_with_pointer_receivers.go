//  In this example, both Scale and Abs are with receiver type *Vertex, even
//though the Abs method needn't modify its receiver.
//In general, all methods on a given type should have either value or pointer
//receivers, but not a mixture of both 

package main

import (
		"fmt"
		"math"
		)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X*f
	v.Y = v.Y*f
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}