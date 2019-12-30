package main
import "fmt"


func main() {
	for x := 1; x <= 10; x++ {
		fmt.Println(x)
	}
	x := 10
	if x % 2 == 0 { //true
		fmt.Println("divisible by 2") //divisible by 2
		x = 14 // update x value
	}
	if x % 5 == 0 { // false 
		fmt.Println("divisible by 5")
	}
	switch x {
	case 1:fmt.Println("one")
	case 2:fmt.Println("two")
	case 3:fmt.Println("tree")
	default:fmt.Println("Defualt Value") //true
		
	}
}

