package main

import "fmt"

func main() {
	const name = "golang"

	// name = "javascript"
	// we never change const value

	const (
		port = 5000
		host = "localhost"
	)

	// we never change this const we define many const in a single

	fmt.Println(port, host)

}