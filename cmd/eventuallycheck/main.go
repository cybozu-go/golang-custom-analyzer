package main

import (
	"github.com/cybozu/neco-containers/golang/analyzer/pkg/eventuallycheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(eventuallycheck.Analyzer)
}
