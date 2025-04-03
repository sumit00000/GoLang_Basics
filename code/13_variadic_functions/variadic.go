package main

import "fmt"


func sum(nums ...int) int  {
	total := 0

	for _, num:= range nums {
		total = total + num
	}

	return total
}

func sums(nums ...any) int  {
	total := 0

	for _, num:= range nums {

		if n, ok := num.(int); ok {
			total = total + n
		} else {
			fmt.Println("Skipping non-integer value:", num)
		}
	}

	return total
}

func main()  {
	result := sum(3, 4, 5, 78, 90)
	fmt.Println(result)


	result2 := sums(3, 4, 6, "golang")
	fmt.Println(result2)

	nums := []int{3, 4, 6, 7, 9}
	result3 := sum(nums...)
	fmt.Println(result3)
}