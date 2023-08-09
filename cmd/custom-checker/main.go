package main

import (
	"github.com/cybozu-go/golang-custom-analyzer/pkg/eventuallycheck"
	"github.com/cybozu-go/golang-custom-analyzer/pkg/restrictpkg"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		restrictpkg.RestrictPackageAnalyzer,
		eventuallycheck.Analyzer,
	)
}
