package main

import "fmt"

// iterating over data structure
func main() {
	nums := []int {6, 7, 8}

	// for i:=0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// sum
    sum := 0
	for _, num := range nums {
		// fmt.Println(num)
		sum = sum + num
	}
	fmt.Println(sum)

	// index
	for i, num := range nums {
		fmt.Println(num, i)
	}


	// with map
	m := map[string]string{"fname": "john", "lname": "doe"}
	for k, v := range m {
		fmt.Println(k, v)
	}

	for k := range m {      // key mean keys and v means value
		fmt.Println(k)
	}


    // unicode code point rune
	// staring byte of rune
	// 300 -> 1 byte, 2 byte
	for i, c := range "golang" {
		fmt.Println(i, c)
		fmt.Println(i, string(c))
	}

}