package main

import "fmt"

// pointers -> it is basically whatever variable or any type of datastructure we store
// in our computer memory and pointers are the address of that stored variable memory location.


// by value 
func changeNum(num int) {
	num = 5
	fmt.Println("In changeNum by value", num)
}

// by reference
// * is basically pointer or address of that varibale directly
func changeNumber(num *int)  {
	*num = 5
	fmt.Println("In changeNum by reference", num)


}
func main() {
	num := 1
	changeNum(num)
	fmt.Println("After changeNum in main by value", num)

	fmt.Println("Memory address", &num)
	changeNumber(&num)
	fmt.Println("After changeNum in main by reference", num)

	
}