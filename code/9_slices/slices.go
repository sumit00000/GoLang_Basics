package main

import (
	"fmt"
	"slices"
)

// slice  -> dynamic array
// most used construct in go
// + useful methods
// if we do not know in advance tat how many elements are placed in this array so there we use slices
func main()  {
	// uninitialized slice is nil(null)
	var num []int

	fmt.Println(num == nil)
    fmt.Println(len(num))

	var nums = make([]int, 0, 5)
	// capacity -> maximum number of elements can fit
	fmt.Println(cap(nums))
	fmt.Println(nums)
	fmt.Println(nums==nil)


	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	fmt.Println(nums)
	fmt.Println(cap(nums))
	fmt.Println(len(nums))

    
	// numbers := []int{}
	var numbers = make([]int, 2, 5)
	numbers[0] = 3
	numbers[1] = 5
	// numbers = append(numbers, 1)
	// numbers = append(numbers, 2)
	fmt.Println(numbers)
	fmt.Println(cap(numbers))
	fmt.Println(len(numbers))




	// copy function
	var numb = make([]int, 0, 5)
	numb = append(numb, 2)

	var numb2 = make([]int, len(numb))

	// copy
	copy(numb2, numb)
	fmt.Println(numb, numb2)




	// slice operator
	var num3 = []int{1,2,3, 4, 5}
	fmt.Println(num3[0:2])     // exclude than 2 mean 1 2
	fmt.Println(num3[1:])



	// slice package
	var num4 = []int{1,2}
	var num5 = []int{1,2}

	fmt.Println(slices.Equal(num4, num5))


	// 2d
	var num6 = [][]int{{1,2,3},{5,6,7}}
	fmt.Println(num6)
} 