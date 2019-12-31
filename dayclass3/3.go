package main
import (
	"fmt"

)

func main() {
	slice1 := []int{1,2,3}
	fmt.Println(slice1) // [1 2 3]
	slice2 := append(slice1, 4,5)
	fmt.Println(slice2) // [1 2 3 4 5]
	slice3 := []int{1,2,3,4,5,6,7,8,9,10,11}
	fmt.Println(slice3) // [1 2 3 4 5 6 7 8 9 10 11]
	slice4 := make([]int, 5) 
	fmt.Println(slice4) // [0 0 0 0 0]
	fmt.Println(slice3,slice4) // [1 2 3 4 5 6 7 8 9 10 11] [0 0 0 0 0]
	c := copy(slice4,slice3) 
	fmt.Println(c) // 5
	fmt.Println(slice4) // [1 2 3 4 5]
}