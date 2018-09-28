penticons.go
============
> An implementation of Penticons in [go](http://golang.org/) programming language.

Identicons similar to GitHub's contribution activity calendar.


[![GoDoc](https://godoc.org/github.com/penticons/penticons.go?status.svg)](http://godoc.org/github.com/penticons/penticons.go)

```go
package main

import (
	"fmt"
	"github.com/penticons/penticons-go"
)

func main() {
	arg := "Penticons"
	p := penticons.Uri_image(arg)
	fmt.Println(p)
}
```

#### Installation

⚠ Use the alias “penticons-go” to use Penticons in your Go code:

```go
go get github.com/penticons/penticons-go
```

#### API

```go
func Generate(arg string) string
```

Returns a SVG string representation of the penticon for the string argument *arg*

```go
func Base64_string(arg string) string
```

Returns Base64 encoded string representation of the penticon for the string argument *arg*

```go
func Uri_image(arg string) string
```

Returns URI image representatino of the penticon for the string argument *arg*

#### LICENSE

Penticons implementations are released under the MIT license.
