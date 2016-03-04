package main

import "fmt"
import "math"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt((v.X*v.X + v.Y*v.Y))
}

//check_type tests whether the interface value i holds a value of
//type string, int or Vertex (user defined type)
func check_type(i interface{}) {
	switch s:= i.(type) { // same syntax as a type assertion i.(T), but the specific type T is replaced with the keyword type.
		case int:
			fmt.Printf("%d is of type integer\n", s)
			fmt.Printf("The square root of %d is %f\n", s, math.Sqrt(float64(s)))

		case string:
			fmt.Printf("'%s' is of type string\n", s)
			fmt.Printf("The string content is: '%s'\n", s)

		case Vertex:
			fmt.Printf("'%v' is a struct of type Vertex\n", s)
			fmt.Printf("The absolute value of %T%v is %f\n", s,s, s.Abs())

		default:
			fmt.Printf("The type '%T' is unknown to me\n", s)
			fmt.Printf("It has the value %v\n", i)

	}
}

/* This switch statement tests whether the interface value i holds a value of
type string or int. In each of the string and int cases, the variable s will be of type string or
int respectively and hold the value held by i. In the default case (where there
is no match), the variable s is of the same interface type and value as i.  */



func main() {
	check_type(42)
	check_type("hello")
	check_type(true)
	check_type(Vertex{3, 4})
}


