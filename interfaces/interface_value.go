package main

import ("fmt"; "math")

type I interface {
	M()
}

type T struct {
	S string
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func (t *T) M(){
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i , i)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	/*The code below will cause a panic run-time error because it will call
	M() with receiver e, where e is a nil value. Following error will be generated:
	panic: runtime error: invalid memory address or nil pointer dereference
	[signal 0xb code=0x1 addr=0x20 pc=0x4015ab]

	goroutine 1 [running]:
	main.main()
		/home/mwcruz/GOLANG/GoCode/src/github.com/mwcruz/golang-apps/interfaces/interface_value.go:48 +0x16b
	exit status 2*/

    var e I
    e.M()
}


/*
The null Interface
type N interface{}

An empty interface may hold values of any type. (Every type implements at least
zero methods.)

Empty interfaces are used by code that handles values of unknown type. For
example, fmt.Print takes any number of arguments of type interface{}.  */


