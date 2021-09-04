package main

import (
	"errors"
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
