package main
import "fmt"

func main()  {
	// age := 16

	// if age >= 18 {
	// 	fmt.Println("person is an adult") 
	// } else {
	// 	fmt.Println("person is not an adult") 
	// }

	// age := 16
	// if age >= 18 {
	// 	fmt.Println("person is an adult") 
	// } else if age >= 12 {
	// 	fmt.Println("person is teenager") 
	// } else {
	// 	fmt.Println("person is a kid")
	// }

    var role = "admin"
	var hasPermissions = false

	// if role == "admin" || hasPermissions {
	// fmt.Println("yes")
	// }

	if role == "admin" && hasPermissions {
		fmt.Println("yes")
	} else {
		fmt.Println("not meet with conditions...")
	}

	// we can declare a variable inside if construct
	// if age :=15; age >= 18 {
	// 	fmt.Println("peron is an aduly", age)
	// } else if age >= 12 {
	// 	fmt.Println("person is an teenager", age)
	// }


	// go does not have ternary, we have to use normal if else

}