<div align="center">  
  <p>
    <strong>Elegant Low-Level Elements</strong>
  </p>
  
  <p>
    <a href="https://github.com/exact/elle/stargazers"><img src="https://img.shields.io/github/stars/exact/elle" alt="Stars Badge"/></a>
    <a href="https://github.com/exact/elle/network/members"><img src="https://img.shields.io/github/forks/exact/elle" alt="Forks Badge"/></a>
    <a href="https://github.com/exact/elle/pulls"><img src="https://img.shields.io/github/issues-pr/exact/elle" alt="Pull Requests Badge"/></a>
    <a href="https://github.com/exact/elle/issues"><img src="https://img.shields.io/github/issues/exact/elle" alt="Issues Badge"/></a>
    <a href="https://github.com/exact/elle/graphs/contributors"><img alt="GitHub contributors" src="https://img.shields.io/github/contributors/exact/elle?color=2b9348"></a>
    <a href="https://github.com/exact/elle/blob/main/LICENSE"><img src="https://img.shields.io/github/license/exact/elle?color=2b9348" alt="License Badge"/></a>
  </p>
</div>

## 🧐 What is this?

elle (Elegant Low-Level Elements) is a lightweight Go library designed to simplify coding.

It provides a set of utilities and building blocks that make Go development faster and more efficient, especially for complex operations.

Inspired by [Elixir](https://github.com/elixir-lang/elixir)'s syntax

**Note: This project is intended for personal use cases, if anyone finds it useful I'm happy but this serves mostly for peer review and improvement.**

## ✨ Features

- 🧼 **Simple** - Stupidly simple syntax that makes coding fun
- 🚀 **Fast** - Zero-overhead model to abstractions & any syntax sugar
- 🛡️ **Secure** - Built from the ground up to be cryptographically secure

## 📝 Example

```go
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
			io.Sleep(500)
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
```

---

<div align="center">
  Made with ❤️ by <a href="https://github.com/exact">@exact</a>
</div>

