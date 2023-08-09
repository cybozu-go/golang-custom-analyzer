[![GitHub release](https://img.shields.io/github/release/cybozu-go/golang-custom-analyzer.svg?maxAge=60)][releases]
[![CI](https://github.com/cybozu-go/golang-custom-analyzer/actions/workflows/ci.yaml/badge.svg)](https://github.com/cybozu-go/golang-custom-analyzer/actions/workflows/ci.yaml)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/cybozu-go/golang-custom-analyzer?tab=overview)](https://pkg.go.dev/github.com/cybozu-go/golang-custom-analyzer?tab=overview)
[![Go Report Card](https://goreportcard.com/badge/github.com/cybozu-go/golang-custom-analyzer)](https://goreportcard.com/report/github.com/cybozu-go/golang-custom-analyzer)

golang custom analyzer
======================

This repository contains custom analysers for Go.

## custom-checker

`custom-checker` is the program to run the following analysers at once:

- `eventuallycheck`
- `restrictpkg`

### Usage

```console
$ custom-checker -restrictpkg.packages=html/template,log targetfile.go
```

## `eventuallycheck`

`eventuallycheck` is a static analysis tool to detect [`gomega.Eventually`](https://godoc.org/github.com/onsi/gomega#Eventually) without [`Should`](https://godoc.org/github.com/onsi/gomega#Should) or [`ShouldNot`](https://godoc.org/github.com/onsi/gomega#ShouldNot).

### Usage

```console
$ eventuallycheck [FILES]
```

### Target functions

- [`Consistently`](https://godoc.org/github.com/onsi/gomega#Consistently)
- [`ConsistentlyWithOffset`](https://godoc.org/github.com/onsi/gomega#ConsistentlyWithOffset)
- [`Eventually`](https://godoc.org/github.com/onsi/gomega#Eventually)
- [`EventuallyWithOffset`](https://godoc.org/github.com/onsi/gomega#EventuallyWithOffset)

## restrictpkg

`restrictpkg` is a static analysis tool to restrict which package to use.

### Usage

Let `PACKAGE1` and `PACKAGE2` be the names of the packages you want to restrict.

```console
$ restrictpkg -packages PACKAGE1,PACKAGE2,...
```
