package main

import "fmt"


func add(a int, b int) int {
	return a + b
}

func getLanguages() (string, string, string) {
	return "golang", "javascript", "c++"
}

func processIt(fn func(a int) int) {
	fn(1)
}

func processItt() func (a int) int  {
	return func (a int) int  {
		return 4
	}
	
}

func main()  {
	result := add(3, 5)
	fmt.Println(result)
	fmt.Println(getLanguages())

	lang1, lang2, lang3 := getLanguages()
	fmt.Println(lang1, lang3, lang2)


	fn := func(a int) int {
		return 2
	}
	processIt(fn)

	fnn := processItt()
	fnn(6)
}