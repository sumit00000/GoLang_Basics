package main

import "fmt"


func printSlice(items []int) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// another way to write programme
func printSlice2[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// for scope
func printSlice3[T int | string | bool](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func printStringSlice(items []string) {
	for _, item := range items {
		fmt.Println(item)
	}
}


// struct example
type stack struct {
	elements []int
}
// by using generic
type stack2 [T any] struct {
	elements []T
}

func printSlice4[T comparable, V string](items []T, name V ) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}


func main()  {
	nums := []int{1,2,3}
	names := []string{"golang, typescript"}
	vals := []bool{true, false, true}

	printSlice(nums)
	printStringSlice(names)
    
	// now its take both or any type even string bool int
	// mow we do not need to add duplicate functions
	printSlice2(names)
	printSlice3(vals)

	myStack := stack{
		elements: []int{1,2,3},
	}
	myStack2 := stack2[string]{
		elements: []string{"golang"},
	}


	fmt.Println(myStack)
	fmt.Println(myStack2)
	fmt.Println(vals)

	printSlice4(vals, "singh")

}