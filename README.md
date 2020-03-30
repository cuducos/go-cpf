# Go CPF ![Tests](https://github.com/cuducos/go-cpf/workflows/Tests/badge.svg)

A Go module to validate CPF numbers (Brazilian unique identifier for the Federal Revenue).

```go
package main

import "github.com/cuducos/go-cpf"


func main() {
	// these return true
	cpf.IsValid("23858488135")
	cpf.IsValid("238.584.881-35")

	// these return false
	cpf.IsValid("111.111.111-11")
	cpf.IsValid("123.456.769/01")
	cpf.IsValid("ABC.DEF.GHI-JK")
	cpf.IsValid("123")

	// this returns 11111111111
	cpf.Unmask("111.111.111-11")

	// this returns 111.111.111-11
	cpf.Mask("11111111111")
}
```

## A bit of story and thankfulness

I started to learn [Go](https://golang.org/) with [_Learn Go With Tests_](https://quii.gitbook.io/learn-go-with-tests/) and this CPF (Brazilian unique identifier for the Federal Revenue) validation script is actually **my very first lines in Go** (except the ones from the book). I'm sharing it here to get **feedback** ❤️
