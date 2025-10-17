package main

import (
	"github.com/exact/elle/io"
)

func main() {
	p := io.SyncPool(1)
	p.Go(func() {
		panic("101010101010101011")
	})
	p.Wait()
	io.Puts("still good!!")
}

/*func main3() {
	io.Puts(io.Get("https://api-cloudfront.life360.com", nil, true))
}

func main() {
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
