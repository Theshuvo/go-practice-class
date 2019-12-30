package main
import (
	"fmt"
	"reflect"
)
func main() {
	var x string = "hello" 
	fmt.Println(reflect.TypeOf(x)) //string
	var y string 
	y= "world"
	fmt.Println(y) // world
	y = "Hello " + y
	fmt.Println(y) // Hello world
}