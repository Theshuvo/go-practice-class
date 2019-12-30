package main
import (
	"fmt"
	"time"
)
func main() {
	fmt.Println("Enter your age:") 
	var age float64
	fmt.Scanf("%f",&age) // 10
	currentTime	:= time.Now()
	DateOfYear := currentTime.Year() - int(age)
	fmt.Println("You're date of birth year", DateOfYear) // You're date of birth year 2009
	// Your mother is 20 years older than you
	age += 20 
	var MotherDateOfYear = currentTime.Year() - int(age)
	fmt.Println("Your mother date of birth year", MotherDateOfYear) // Your mother date of birth year 1989
}