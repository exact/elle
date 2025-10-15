package main

import (
	"sync"

	"github.com/exact/elle/io"
	"github.com/exact/elle/secure"
)

func main() {
	io.Puts(io.Get("https://api-cloudfront.life360.com", nil, true))
}

func main2() {
	var wg sync.WaitGroup
	pool := io.Pool(25)

	for range secure.Number(1000, 2500) {
		pool.Add(&wg, func() {
			//io.Puts("starting...")
			//io.Sleep(1000)
			io.Puts("new:", secure.NewUserAgent())
		})
	}

	wg.Wait()
}
