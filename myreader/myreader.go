//Implement a Reader type that emits an infinite
// stream of the ASCII character 'A'
package main

import ("golang.org/x/tour/reader"
		"strings"
		"io"
		)

//type MyReader struct with one element of type io.Reader
type MyReader struct{
	R io.Reader

}

//Add a Read([]byte) (int, error) method to MyReader.
func (r *MyReader) Read(b []byte) (int, error) {
		r.R = strings.NewReader("A") //assign to r.R a stream of strings "A"
		n, err := r.R.Read(b)	// assign to n, err the result of calling Read on the MyReader pointer receiver
		if err != nil {
			return n, err
			}
		return n, err
}


func main() {
	reader.Validate(&MyReader{})
}



