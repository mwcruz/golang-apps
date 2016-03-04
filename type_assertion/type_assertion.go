package main

import "fmt"

func main() {
	var i interface{} = "Ola"
	fmt.Println(i.(string))

	s := i.(string) // asserts that the interface value i holds the concrete type string 
	fmt.Println(s)  // and assigns the underlying string value to the variable s. 

	s, ok := i.(string) //Test whether interface i holds value of type string. in this case it does so we
	fmt.Println(s, ok)  // get "Ola" for the value and true for ok.

	f,ok := i.(float64) // Test whether interface i hold value of type float64. in this case it does not so we
	fmt.Println(f, ok)  // get 0 for the value and false for ok. No panic will generated, however

	f = i.(float64)  // This however, will cause panic: interface conversion: interface string in not float64
	fmt.Println(f)



}