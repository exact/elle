package main

import (
	"time"

	"github.com/exact/elle/io"
	"github.com/exact/elle/random"
	"github.com/exact/elle/types"
)

// Create a global, synced pool that can be used anywhere
var pool = io.SyncPool(30)

func main() {
	// Measure how long this function takes to execute
	defer io.Timer("main")()

	// A zero-size, empty struct is included for efficiency
	test := make(chan types.None, 1)
	test <- types.None{}

	// Simulate a panicing function, it will recover automatically
	pool.Go(func() {
		time.Sleep(3 * time.Second)
		panic("something went horribly wrong!")
	})

	// Simulate a random amount of I/O work
	for i := range random.Number(25, 50) {
		pool.Go(func() {
			resp, err := io.Request("GET", "https://example.com/", nil, nil, true)
			if err != nil {
				io.Log.Warn("Request Failed", "thread", i)
			} else {
				io.Log.Info("Request Done", "thread", i, "resp", resp)
			}
		})
	}

	// Wait for all work to finish
	pool.Wait()
}
