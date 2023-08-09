package main

import (
	"github.com/cybozu-go/golang-custom-analyzer/pkg/eventuallycheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(eventuallycheck.Analyzer)
}
