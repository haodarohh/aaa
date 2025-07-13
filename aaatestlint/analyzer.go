package aaatestlint

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "aaatest",
	Doc:  "checks unit test functions for // Arrange, // Act, and // Assert comments in correct order",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		for _, decl := range file.Decls {
			funcDecl, ok := decl.(*ast.FuncDecl)
			if !ok || !strings.HasPrefix(funcDecl.Name.Name, "Test") || funcDecl.Body == nil {
				continue
			}

			var found []string
			for _, commentGroup := range file.Comments {
				for _, comment := range commentGroup.List {
					if !positionInFunc(funcDecl, comment.Pos()) {
						continue
					}

					trimmed := strings.TrimSpace(strings.TrimPrefix(comment.Text, "//"))

					switch {
					case strings.HasPrefix(trimmed, "arrange"):
						found = append(found, "arrange")
					case strings.HasPrefix(trimmed, "act"):
						found = append(found, "act")
					case strings.HasPrefix(trimmed, "assert"):
						found = append(found, "assert")
					}
				}
			}

			// Check order
			order := map[string]int{"arrange": 0, "act": 1, "assert": 2}
			last := -1
			for _, keyword := range found {
				if order[keyword] < last {
					pass.Reportf(funcDecl.Pos(), "invalid AAA pattern order: must follow arrange -> act -> assert\n")
					break
				}
				last = order[keyword]
			}

			// Must contain at least Act and Assert in order
			hasAct := false
			hasAssert := false
			for _, k := range found {
				if k == "act" {
					hasAct = true
				}
				if k == "assert" {
					hasAssert = true
				}
			}
			if !(hasAct && hasAssert) {
				pass.Reportf(funcDecl.Pos(), "missing required keywords: need at least act and assert\n")
			}
		}
	}
	return nil, nil
}

func positionInFunc(fn *ast.FuncDecl, pos token.Pos) bool {
	return fn.Body != nil && pos >= fn.Body.Pos() && pos <= fn.Body.End()
}
