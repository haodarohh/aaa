package aaatestlint_test

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/haodarohh/aaa/aaatestlint"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAAACommentCheck(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	testdata := filepath.Join(filepath.Dir(filename), "..", "testdata")
	analysistest.Run(t, testdata, aaatestlint.Analyzer, "a")
}

func TestAAACommentCheckFromLinter(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	testdata := filepath.Join(filepath.Dir(filename), "..", "testdata")

	l, _ := aaatestlint.New(nil)
	a, _ := l.BuildAnalyzers()

	analysistest.Run(t, testdata, a[0], "a")
}
