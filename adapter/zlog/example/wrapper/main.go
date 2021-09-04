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
	w := logger.NewWrapper(l)
	doSomething(w)
}

func doSomething(l logger.Logger) {
	l.Info("a meaningful std message")
}
