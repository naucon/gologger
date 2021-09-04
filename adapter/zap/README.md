# ZAP Adapter

Implementation of the [zap](https://github.com/uber-go/zap) logger for a [Logger Interface](../../README.md) in Go.

## Install

see documentation of the [Logger Interface](../../README.md).

## Usage

```go
package main

import (
  "github.com/naucon/gologger"
  "github.com/naucon/gologger/adapter/zlog"
  "github.com/rs/zerolog"
  "os"
)

func main() {
  z := zerolog.New(os.Stdout)
  l := zlog.NewAdapter(&z)
  doSomething(l)
}

func doSomething(l logger.Logger) {
  l.Info("a meaningful std message")
}

```
