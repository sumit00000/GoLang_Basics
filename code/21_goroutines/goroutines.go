package main

import (
	"fmt"
	"time"
)

// func task(id int) {
// 	fmt.Println("doing task", id)
// }

// in this all functions execute concurrently
// in this how multithreading work
// go ---> we run this function another goroutines
func main()  {
	for i:= 0; i <= 10; i++ {
		// go task(i)

		func(i int) {
			fmt.Println(i)
		} (i)
	}

	time.Sleep(time.Second * 1)
}


