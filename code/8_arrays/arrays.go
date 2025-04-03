package main

import "fmt"

// numbered sequence of specific length
func main()  {
    // int -> 0, string -> "", bool -> false
	// zero values
	var nums [4]int

	nums[0] = 7

	fmt.Println(nums)
	// array length
	// fmt.Println((len(nums)))


	var vals [4]bool
	vals[2] = true
	fmt.Println(vals)


	var names [3]string
	names[2] = "golang"
	fmt.Println(names)

    // to declare it in single line 
	numbers := [3]int{1, 2, 3}
	fmt.Println(numbers)


	// 2d array
	number := [2][2]int{{3, 4},{5, 6}}
	fmt.Println(number)

	// - fixed size, that is predictable
	// - memory optimization
	// - constant time access
}