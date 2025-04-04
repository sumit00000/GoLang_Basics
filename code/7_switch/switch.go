package main

import (
	"fmt"
	// "time"
)

func main() {
	// simple switch 
	// in go we do not use to break keyword
	// i :=5

	// switch i {
	// case 1:
	// 	fmt.Println(("one"))
	// case 2:
	// 	fmt.Println(("two"))
	// case 3:
	// 	fmt.Println(("three"))
	// default:
	// 	fmt.Println("other")
	// }

	// multiple condition switch
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println(("it's weekend"))
	// default: fmt.Println("it's workday")
	// }


	// type swaitch
	whoAmI := func (i any)  {
		switch i.(type) {
		case int:
			fmt.Println("its an integer")
		case string:
			fmt.Println("its a string")
		case bool:
			fmt.Println("its a boolean")
		default:
			fmt.Println("other")
		}
	}
	whoAmI("string")
}