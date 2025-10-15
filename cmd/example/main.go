package main

import (
	"sync"

	"github.com/exact/elle/io"
)

func main() {
	outsideVariable := "wow"

	var wg sync.WaitGroup
	pool := io.Pool(25)
	for range 100 {
		pool.Add(&wg, func() {
			io.Puts("starting...")
			io.Sleep(1000)
			io.Puts("captured:", outsideVariable)
		})
	}

	wg.Wait()
}
