package main

import (
	"github.com/haodarohh/aaa/aaatestlint"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	l, _ := aaatestlint.New(nil)
	a, _ := l.BuildAnalyzers()
	singlechecker.Main(a[0])
}
