package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i < 10; i++ {
		wg.Add(1)
		doWork(i, &wg)
	}
	wg.Wait()
}

func doWork(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(i)
}

