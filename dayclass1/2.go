package main 
import (
	"fmt"
	"reflect"
)

// this is a comment

func main(){
	x:=1.0+1.0 // 1
	fmt.Println(reflect.TypeOf(x)) // float64
	fmt.Println(x) // 1
	fmt.Println("1 + 1 =",1+1) // 1 + 1 = 2
	fmt.Println("1.0 + 1.0 =",1.0+1.0) // 1.0 + 1.0 = 2
	fmt.Println(len("Hello word")) // 10
	fmt.Println("Hello World"[1]) // 101
	fmt.Println("Hello "+"world") // Hello World
	fmt.Println(true && true) // true
	fmt.Println(true && false) // false
	fmt.Println(true || true) // true
	fmt.Println(true || false) // true
	fmt.Println(!true) // flase
}
