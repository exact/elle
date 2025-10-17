package main

import (
	"time"

	"github.com/exact/elle/io"
	"github.com/exact/elle/random"
)

func main() {
	// Make a pool of synced goroutines
	p := io.SyncPool(20)

	// Handle a misbehaving function automatically
	p.Go(func() {
		time.Sleep(3 * time.Second)
		panic("something went horribly wrong D:")
	})

	// Simulate a random amount of concurrent work
	for i := range random.Number(50, 150) {
		p.Go(func() {
			time.Sleep(500 * time.Millisecond)
			io.Log(io.S("[go %d] done!", i))
		})
	}

	// Wait for all work to finish
	p.Await()

	// Finish while showing execution is fine
	io.Log("goodbye!")
}
