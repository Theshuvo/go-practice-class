package main
import (
	"fmt"
	"reflect"
)
func main() {
	var x string = "hello"
	fmt.Println(reflect.TypeOf(x))
	var y string 
	y= "world"
	fmt.Println(y)
	y = "Hello " + y
	fmt.Println(y)
}