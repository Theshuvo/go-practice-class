package main
import (
	"fmt"
	"reflect"
)

func main() {
	x := [5]float64{ 98, 93, 77, 82, 83 } 
	fmt.Println(x) // [98 93 77 82 83]
	fmt.Println(reflect.TypeOf(x))// [5]float64
	 y := [5]float64{ 
		98,
		93,
		77,
		82,
		// 83,
	}
	fmt.Println(y)  // 98 93 77 82 0]
	var a []float64
	fmt.Println(a) // []
	b := make([]float64, 5, 5)
	fmt.Println(b) //[0 0 0 0 0]
	b = append(b,6) 
	fmt.Println(b) // [0 0 0 0 0 6]
	b = []float64{1,2,3,4,5}
	fmt.Println(b) //[1 2 3 4 5]
	fmt.Println(b[1:5])// [2 3 4 5]
	fmt.Println(b[0:len(b)])// [1 2 3 4 5]
}