package io

import (
	"fmt"
	"runtime/debug"
	"sync"

	"github.com/exact/elle/types"
)

type sema chan types.None
type sync_sema struct {
	c sema
	g *sync.WaitGroup
}

// Recover from any panics and output stack trace
func recoverPanic() {
	if v := recover(); v != nil {
		Log.Error(fmt.Sprintf("panic: %q\n%s", v, debug.Stack()))
	}
}

// Acuiring and releasing semaphore helpers
func (s sema) acquire()      { s <- types.None{} }
func (s sema) release()      { <-s }
func (s sync_sema) acquire() { s.c <- types.None{} }
func (s sync_sema) release() { <-s.c }

// OpenPool acts as lightweight semaphore and is used to limit the amount of active goroutines actually doing work at once.
func OpenPool(n int) sema {
	return make(sema, n)
}

// Go submits a function to be executed immediately.
//
// It will wait for an available goroutine before actually doing work.
func (s sema) Go(f func()) {
	go func() {
		s.acquire()
		defer s.release()
		defer recoverPanic()
		f()
	}()
}

// SyncPool acts as a lightweight semaphore, similarly to Pool, but includes a sync.WaitGroup which allows further coordination in work.
func SyncPool(n int) sync_sema {
	var wg sync.WaitGroup
	return sync_sema{c: make(sema, n), g: &wg}
}

// Go submits a function to be executed immediately.
//
// It will wait for an available goroutine before actually doing work.
//
// As well as incrementing the included sync.WaitGroup
func (s sync_sema) Go(f func()) {
	s.g.Add(1)
	go func() {
		s.acquire()
		defer s.release()
		defer s.g.Done()
		defer recoverPanic()
		f()
	}()
}

// Wait simply waits for the included sync.WaitGroup to complete.
func (s sync_sema) Wait() {
	s.g.Wait()
}
