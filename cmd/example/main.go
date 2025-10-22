package main

import (
	"time"

	"github.com/exact/elle/io"
	"github.com/exact/elle/random"
)

var p = io.SyncPool(30)

func main() {
	// Handle a misbehaving function automatically
	p.Go(func() {
		time.Sleep(3 * time.Second)
		panic("something went horribly wrong D:")
	})

	// Simulate a random amount of concurrent work
	for i := range random.Number(50, 150) {
		/*p.Go(func() {
			time.Sleep(500 * time.Millisecond)
			io.Print("done!", "thread", i)
		})*/
		p.Go(func() {
			resp, err := io.Request("GET", "https://example.com/", nil, nil, true)
			if err != nil {
				io.Log.Warn("Request Failed", "thread", i)
			} else {
				io.Log.Info("Request Done", "thread", i, "resp", resp)
			}
		})
	}

	// Wait for all work to finish
	p.Wait()
	io.Log.Info("goodbye!")
}
