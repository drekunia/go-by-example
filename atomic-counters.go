package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var ops atomic.Uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				time.Sleep(10 * time.Millisecond)
				ops.Add(1)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops:", ops.Load())
}
