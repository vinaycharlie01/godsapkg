# godsapkg

> A small, easy-to-use Go utility package with handy helpers everyone can use.

**Status:** WIP — community friendly

## Overview

`godsapkg` is a collection of small, well-tested utility functions and helpers for Go projects. The goal is to keep APIs tiny, readable, and easy to use so any developer (beginners included) can pick them up quickly.

This repository is meant to be used as a utility library — copy the functions you need, or import the package.

## Features

* Small, focused helpers (strings, slices, math, io helpers)
* Clear and simple APIs
* Unit tests for each helper
* Easy-to-read examples

## Installation

Use `go get` to add the package to your project:

```bash
go get github.com/vinaycharlie01/godsapkg
```

Then import it:

```go
import "github.com/vinaycharlie01/godsapkg"
```

If you prefer to use specific subpackages (recommended):

```go
import "github.com/vinaycharlie01/godsapkg/strings"
```

## Quick start

Example: reverse a string (from `godsapkg/strings`):

```go
package main

import (
    "fmt"
    "github.com/vinaycharlie01/godsapkg/strings"
)

func main() {
    s := "hello"
    r := strings.ReverseString([]byte(s))
    fmt.Println(string(r)) // "olleh"
}
```

## Package layout (suggested)

```
/godsapkg
  /strings    # byte/string helpers
  /slice      # slice helpers
  /mathx      # small math helpers
  /io         # io helpers, helpers that wrap common io tasks
  /test       # test helpers
  go.mod
  README.md
```

## API examples

### strings

* `ReverseString([]byte) []byte` — reverse bytes in-place and return result
* `IsPalindrome(string) bool` — check palindrome

### slice

* `Chunk[T any](slice []T, size int) [][]T` — split into chunks
* `Contains[T comparable](slice []T, v T) bool` — check presence

(Keep functions small and well documented.)

## Tests

Write tests for each helper. Run tests with:

```bash
go test ./...
```

Make sure new functions include table-driven tests.

## Contributing

Contributions are welcome! A simple guide:

1. Fork the repo
2. Create a feature branch (`git checkout -b feat/your-feature`)
3. Write code and tests
4. Run `go test ./...`
5. Open a PR with a clear title and description

Please follow Go formatting (`gofmt`) and keep APIs simple. Add unit tests for bug fixes and new features.

## Coding style

* Prefer small, well-named functions
* Use simple English for comments and docs
* Keep public API minimal and stable
* Use generics where it simplifies APIs (Go 1.18+)

## License

This project is licensed under the MIT License — see `LICENSE` for details.

## Maintainers

* Your Name (@yourusername)

## Contact

If you have questions, open an issue or email: [your.email@example.com](mailto:your.email@example.com)

---
