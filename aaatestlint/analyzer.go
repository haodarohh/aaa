package aaatestlint

import (
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "aaatestlint",
	Doc:  "checks that test functions follow Act and Assert comments, with optional Arrange",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Recv != nil || !isTestFunc(fn) {
				continue
			}

			positions := map[string]int{}
			order := []string{}

			for _, group := range file.Comments {
				for _, comment := range group.List {
					if !commentInFunc(pass.Fset, comment, fn) {
						continue
					}

					text := strings.TrimSpace(strings.TrimPrefix(comment.Text, "//"))
					switch text {
					case "Arrange", "Act", "Assert":
						positions[text] = pass.Fset.Position(comment.Pos()).Line
						order = append(order, text)
					}
				}
			}

			if _, ok := positions["Act"]; !ok {
				pass.Reportf(fn.Pos(), "missing '// Act' section in test")
				continue
			}
			if _, ok := positions["Assert"]; !ok {
				pass.Reportf(fn.Pos(), "missing '// Assert' section in test")
				continue
			}

			// Enforce order
			actIndex := indexOf(order, "Act")
			assertIndex := indexOf(order, "Assert")
			if actIndex == -1 || assertIndex == -1 || actIndex > assertIndex {
				pass.Reportf(fn.Pos(), "// Act must appear before // Assert")
			}

			if arrangeIndex := indexOf(order, "Arrange"); arrangeIndex != -1 && arrangeIndex > actIndex {
				pass.Reportf(fn.Pos(), "// Arrange must appear before // Act")
			}
		}
	}

	return nil, nil
}

func isTestFunc(fn *ast.FuncDecl) bool {
	return fn.Name != nil &&
		strings.HasPrefix(fn.Name.Name, "Test") &&
		fn.Type.Results == nil &&
		len(fn.Type.Params.List) == 1
}

func indexOf(slice []string, target string) int {
	for i, s := range slice {
		if s == target {
			return i
		}
	}
	return -1
}

func commentInFunc(fset *token.FileSet, comment *ast.Comment, fn *ast.FuncDecl) bool {
	return comment.Pos() >= fn.Body.Lbrace && comment.End() <= fn.Body.Rbrace
}
