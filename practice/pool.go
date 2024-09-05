package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	id int
}

func Worker(id int, taskChan <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range taskChan {
		fmt.Println("Worker", id, "processing task", task.id)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	numWorkers := 3

	taskChan := make(chan Task)

	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go Worker(i, taskChan, &wg)
	}

	for i := 1; i <= 5; i++ {
		taskChan <- Task{id: i}
	}

	close(taskChan)

	wg.Wait()

	fmt.Println("All tasks processed")
}
