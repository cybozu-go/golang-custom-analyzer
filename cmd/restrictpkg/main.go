package main

import (
	"github.com/cybozu-go/golang-custom-analyzer/pkg/restrictpkg"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(restrictpkg.RestrictPackageAnalyzer)
}
