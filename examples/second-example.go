package examples

import (
	"fmt"
	"sync"
	"time"
)

// This is the function we'll run in every goroutine.
func run_worker(id int) error {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
	return nil
}

func start() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			run_worker(i)
		}()
	}

	wg.Wait()
}
