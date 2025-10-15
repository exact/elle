package main

import (
	"github.com/exact/elle/io"
)

func main() {
	io.Puts(io.Get("https://api-cloudfront.life360.com", nil, true))
}

/*func main() {
	var wg sync.WaitGroup
	pool := io.Pool(25)

	for range 100 {
		pool.Add(&wg, func() {
			//io.Puts("starting...")
			//io.Sleep(1000)
			io.Puts("new:", secure.NewUserAgent())
		})
	}

	wg.Wait()
}*/
