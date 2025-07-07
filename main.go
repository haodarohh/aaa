package main

import (
	"github.com/haodarohh/aaa/aaatestlint"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(aaatestlint.Analyzer)
}
