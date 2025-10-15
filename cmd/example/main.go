package main

import (
	"sync"

	"github.com/exact/elle/io"
	"github.com/exact/elle/secure"
)

func main() {
	var wg sync.WaitGroup

	// Make a pool of 25 goroutines
	pool := io.Pool(25)

	// Do some work out of pool
	io.Async(func() {
		for range 5 {
			io.Sleep(secure.Number(100, 500))
			io.Puts("bg completed!")
		}
	})

	// Simulate random, concurrent work
	for range secure.Number(50, 150) {
		pool.Add(&wg, func() {
			io.Puts("working...")
			io.Sleep(secure.Number(1000, 2500))
			io.Puts("worked:", secure.NewUserAgent())
		})
	}

	// Wait for work to finish
	wg.Wait()
}
