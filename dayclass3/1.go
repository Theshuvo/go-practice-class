package main
import (
	"fmt"
	"reflect"
)

func main() {
	var x [5]int
	x[0] = 100
	fmt.Println(x) // [100 0 0 0 0]
	fmt.Println(reflect.TypeOf(x)) // [5]int
	var y[5]float64
	y[0] = 100
	fmt.Println(y) // [100 0 0 0 0]
	fmt.Println(reflect.TypeOf(y)) // [5]float64
	y[0] = 98
	y[1] = 93
	y[2] = 77
	y[3] = 82
	y[4] = 83
	var total float64 = 0
	for i := 0; i < len(y); i++ {
		total += y[i]
	}
	fmt.Println(total) // 433
	fmt.Println(int(total) / len(y)) // 433 / 5 = 86  Wrong 
	fmt.Println(total / float64(len(y))) // 433 / 5 = 86.6
	total = 0
	for _, value := range y { 
		total += float64(value)
	  }
	  fmt.Println(total / float64(len(x))) // 433 / 5 = 86.6
}