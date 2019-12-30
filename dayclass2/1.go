package main
import (
	"fmt"
	"reflect"
)
var title string = "Saha"

func main() {
	name :="Shuvo"
	fmt.Println(reflect.TypeOf(name)) //string
	fmt.Println("My name is",name + " " + title) // My name is Shuvo Saha
}
