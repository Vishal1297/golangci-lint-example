package examples

import (
	"log"
	"sync"
	"time"
)

// This is the function we'll run in every goroutine.
func worker(id int) {
	log.Printf("Worker %d starting", id)
	time.Sleep(time.Second)
	log.Printf("Worker %d done", id)
}

// Make some changes to produce errors for linter.
// Start starts the worker.
func Start() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
}
