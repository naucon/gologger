# Go Standard library Logger Adapter

Implementation of the standard library logger for a [Logger Interface](../../README.md) in Go.

## Install

see documentation of the [Logger Interface](../../README.md).

## Usage

```go
package main

import (
  "github.com/naucon/gologger"
  "github.com/naucon/gologger/adapter/std"
  "log"
)

func main() {
  l := std.NewAdapter(log.Default())
  doSomething(l)
}

func doSomething(l logger.Logger) {
  l.Info("a meaningful std message")
}
```
