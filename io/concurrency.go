package io

import (
	"sync"
	"time"
)

type none struct{}

type sema chan none

func Pool(n int) sema {
	return make(chan none, n)
}

func (s sema) Acquire() { s <- none{} }
func (s sema) Release() { <-s }

func (s sema) Go(f func()) {
	go func() {
		s.Acquire()
		defer s.Release()
		f()
	}()
}

func (s sema) Add(wg *sync.WaitGroup, f func()) {
	wg.Add(1)
	go func(g *sync.WaitGroup) {
		s.Acquire()
		defer s.Release()
		defer g.Done()
		f()
	}(wg)
}

func Async(f func()) {
	go f()
}

func Sleep(n int) {
	time.Sleep(time.Duration(n) * time.Millisecond)
}
