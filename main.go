package main

import (
	"github.com/haodarohh/aaa/aaatestlint"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(aaatestlint.Analyzer)
}

func New(conf any) ([]*analysis.Analyzer, error) {
	// fmt.Printf("My configuration (%[1]T): %#[1]v\n", conf)

	// The configuration type will be map[string]any or []interface, it depends on your configuration.
	// You can use https://github.com/go-viper/mapstructure to convert map to struct.

	return []*analysis.Analyzer{aaatestlint.Analyzer}, nil
}
