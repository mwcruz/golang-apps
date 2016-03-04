//Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
// Implementation: The main() function will implement a loop to recursively call fibonacci() for consecutive integer values. Fibonacci() is implemented as a function that returns as its value another function. The return from fibonacci will be assigned to a variable f int, which will operate as a closure and can be passed as parameter to other funtions.

package main


import "fmt"

func fibonacci() func() int {
	fib1, fib2 := 0, 1
	return func() int {
		fib1, fib2 = fib2, fib1+fib2
		return fib1
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
