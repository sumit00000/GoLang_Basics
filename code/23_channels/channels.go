// channels -- it similar as like a pipe in which where we send data from one side and received this data in second side
// main concept -- where multiples go routines are running  these are run concurrently
// where we want to send or received data one goruoutines to another goroutines so this whole process is done by channels

// deadlock --->  when multiple processes run concurrently these process hold the resource like memory and another process
// wait for resources  both are depention on anotehr process and stuck in infinity loop, her go will provide deadlock error

package main

import (
	"fmt"
	"time"
)

func processsNum(numChan chan int)  {
	fmt.Println("processing numv=ber", <- numChan)
}

func main()  {

	numChan := make(chan int)

	go processsNum(numChan)

	numChan <- 5

	time.Sleep(time.Second *  2)


	// messageChan := make(chan string)

	// messageChan <- "ping"  // blocking

	// msg := <- messageChan

	// fmt.Println(msg)
}