package main
import "fmt"

type Plants struct {
    Name,
    OtherNames,
    CountriesOfOrigin,
    Light,
    Watering,
    Fertilizing,
    soil,  string
    

func main() {
x := 5
y := 25
swap(&x, &y)
fmt.Println(x,y) // x is still 5
}
