package main

import (
	"errors"
	"github.com/naucon/gologger"
	"github.com/naucon/gologger/adapter/std"
	"log"
)

func main() {
	l := std.NewAdapter(log.Default())
	doSomething(l)
}

func doSomething(l logger.Logger) {
	err := errors.New("some error")
	fields := map[string]interface{}{
		"key": "value",
	}

	l.Error("a meaningful std message")
	l.ErrorWithFields(fields, "a meaningful std message")
	l.Errorf("a meaningful std message: %v", err)
	l.ErrorfWithFields(fields, "a meaningful std message: %v", err)
	l.ErrorErr(err)
	l.Warn("a meaningful std message")
	l.WarnWithFields(fields, "a meaningful std message")
	l.Warnf("a meaningful std message: %v", err)
	l.WarnfWithFields(fields, "a meaningful std message: %v", err)
	l.WarnErr(err)
	l.Info("a meaningful std message")
	l.InfoWithFields(fields, "a meaningful std message")
	l.Infof("a meaningful std message: %v", err)
	l.InfofWithFields(fields, "a meaningful std message: %v", err)
	l.Debug("a meaningful std message")
	l.DebugWithFields(fields, "a meaningful std message")
	l.Debugf("a meaningful std message: %v", err)
	l.DebugfWithFields(fields, "a meaningful std message: %v", err)
}
