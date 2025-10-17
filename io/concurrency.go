package io

import (
	"runtime/debug"
	"sync"

	"github.com/exact/elle/types"
)

type sema chan types.Nothing
type sync_sema struct {
	c sema
	g *sync.WaitGroup
}

// Recover from any panics and output stack trace
func savePanic() {
	if v := recover(); v != nil {
		Warn(S("panic: %q\n%s", v, debug.Stack()))
	}
}

// Acuiring and releasing semaphore helpers
func (s sema) lock()        { s <- types.Nothing{} }
func (s sema) unlock()      { <-s }
func (s sync_sema) lock()   { s.c <- types.Nothing{} }
func (s sync_sema) unlock() { <-s.c }

// Pool acts as lightweight semaphore and is used to limit the amount of active goroutines actually doing work at once.
func Pool(n int) sema {
	return make(chan types.Nothing, n)
}

// Go submits a function to be executed immediately.
//
// It will wait for an available goroutine before actually doing work.
func (s sema) Go(f func()) {
	go func() {
		s.lock()
		defer s.unlock()
		defer savePanic()
		f()
	}()
}

// SyncPool acts as a lightweight semaphore, similarly to Pool, but includes a sync.WaitGroup which allows further coordination in work.
func SyncPool(n int) sync_sema {
	var wg sync.WaitGroup
	return sync_sema{
		c: make(chan types.Nothing, n),
		g: &wg,
	}
}

// Go submits a function to be executed immediately.
//
// It will wait for an available goroutine before actually doing work.
//
// As well as incrementing the included sync.WaitGroup
func (s sync_sema) Go(f func()) {
	s.g.Add(1)
	go func() {
		s.lock()
		defer s.unlock()
		defer s.g.Done()
		defer savePanic()
		f()
	}()
}

// Wait simply waits for the included sync.WaitGroup to complete.
func (s sync_sema) Await() {
	s.g.Wait()
}
