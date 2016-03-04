package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}


type MyFloat float64
/*Here are two ways of writting the same method Abs().
Methods are simply functions with a receiver argument
in this case the receiver is v Vertex*/

func (v Vertex) Abs1() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


func Abs2(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}else { return float64(f)}
}

/*Methods with receivers as pointers, i.e. the receiver type is *T
Scale method of Vertex with pointer receiver. When the receiver is a pointer, then the method can change the value of the variable it points to. When the receiver is a value, the the method will act on a copy of the variable and not change the variable itself. Thus, is most common to use pointer receivers.*/

func (v *Vertex) Scale(f float64) { // function with a pointer receiver v *Vertex
	v.X = v.X*f
	v.Y = v.Y*f
}

//An Aletrnative way to write the Scale method
func Scale1(v *Vertex, f float64) {  // function with a pointer argument v *Vertex
	v.X = v.X*f
	v.Y = v.Y*f
}

func main() {
	v := Vertex{3, 4}
	w := Vertex{5, 2}
	f := MyFloat(5)
	scale := float64(10)
	fmt.Println(Abs2(v))
	fmt.Println(v.Abs1())
	fmt.Printf("Abs val of %f is %f\n", f, f.Abs())
	p := &v // p is a pointer referencing v
	// function with pointer receivers take either a value or a pointer as the receiver when they are called:
	v.Scale(scale)  //value receiver
	fmt.Printf("Scale with value receiver results in %v\n", v)
	p.Scale(5) // pointer receiver
	fmt.Printf("Scale with pointer receiver results in %v\n", p)
	Scale1(&w, scale) // function with pointer argument must take a pointer &W
	fmt.Printf("scaled by %f results in %v\n", scale, w)
	fmt.Printf("Abs value of w is %f or %f\n", w.Abs1(), Abs2(w))
}



/*Summary: Pointers or Value Receivers?
Prefer pointers for two reasons:
1.	to allow method to change values of the variable that receiver points to
2. value receivers make a copy of the original object, while pointers don't. Thus pointers are more efficient in resource usage.
In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both
*/