# Logger Interface for Go

[![Build](https://github.com/naucon/gologger/actions/workflows/go-ci.yml/badge.svg)](https://github.com/naucon/gologger/actions/workflows/go-ci.yml)
[![Coverage](https://codecov.io/gh/naucon/gologger/branch/main/graph/badge.svg?token=DNODQMR5RZ)](https://codecov.io/gh/naucon/gologger)

This package is a logger interface for go projects. The goal is to use a common logging interface without coupling the project to concrete logging implementations.

## Adapter

It comes with the following adapters to common logger packages:

* [Standard library](adapter/std/README.md)
* [Zap](adapter/zap/README.md)
* [Zerolog](adapter/zlog/README.md)

To implement another logger is straight forward. Write your own adapter that implements following functions:

```go
type Logger interface {
  Error(msg string)
  Errorf(format string, v ...interface{})
  ErrorErr(err error)
  Warn(msg string)
  Warnf(format string, v ...interface{})
  WarnErr(err error)
  Info(msg string)
  Infof(format string, v ...interface{})
  Debug(msg string)
  Debugf(format string, v ...interface{})
}
```

## Installation

install the latest version via go get

```
go get -u github.com/naucon/gologger
```

## Import package

```
import (
  "github.com/naucon/gologger"
)
```

## Usage

First create an instance of your logger. In this example we're using zerolog.

```go
z := zerolog.New(os.Stdout)
```

Import the adapter for your concrete logger.

```
import (
  adapter "github.com/naucon/gologger/adapter/zlog"
)
```

Next create an adapter instance with `adapter.NewAdapter()` and pass in the logger instance.

```go
l := zlog.NewAdapter(&z)
```

Finally, you can decouple the logger implementation by using the `Logger` interface.

```go
func doSomething(l logger.Logger) {
  err := errors.New("some error")
  l.Error("a meaningful std message")
  l.Errorf("a meaningful std message: %v", err)
  l.ErrorErr(err)
  l.Warn("a meaningful std message")
  l.Warnf("a meaningful std message: %v", err)
  l.WarnErr(err)
  l.Info("a meaningful std message")
  l.Infof("a meaningful std message: %v", err)
  l.Debug("a meaningful std message")
  l.Debugf("a meaningful std message: %v", err)
}
```

## Logger Wrapper

The package comes with a `loggerWrapper` to wrap another adapter.

```go
  z := zerolog.New(os.Stdout)
  l := zlog.NewAdapter(&z)
  w := logger.NewWrapper(l)
  w.Info("a meaningful std message")
```

## Example

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
