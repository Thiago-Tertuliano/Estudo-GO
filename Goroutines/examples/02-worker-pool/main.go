package main

import (
	"fmt"
	"sync"
	"time"
)

type job struct {
	ID int
}

func worker(id int, jobs <-chan job, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("worker %d: processando job %d\n", id, j.ID)
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Printf("worker %d: encerrou\n", id)
}

func main() {
	const numWorkers = 3
	const numJobs = 10

	jobs := make(chan job, numWorkers) //backpressure leve
	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	for i := 1; i <= numJobs; i++ {
		jobs <- job{ID: i}
	}
	close(jobs)

	wg.Wait()
	fmt.Println("main: todos os jobs processados")
}