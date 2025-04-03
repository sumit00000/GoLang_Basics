package main

import (
	"fmt"
	"sync"
)

// *sync.WaitGroup this is a pointer
func task(id int, w *sync.WaitGroup) {
	defer w.Done()                        // deffer is run at last or after function is executed defer part is run
	fmt.Println("doing task", id)
}

// instead of use time sleep use wait groups
func main()  {

	var wg sync.WaitGroup
	for i:= 0; i <= 10; i++ {
		wg.Add(1)        // 1 beacus 1 goroutines is started
		go task(i, &wg)
	}

	wg.Wait()
}


