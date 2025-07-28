package aaatestlint

import (
	"go/ast"
	"go/token"
	"strings"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("aaa", New)
}

type Settings struct {
	One   string `json:"one"`
	Two   int    `json:"two"`
	Three bool   `json:"three"`
}

type Plugin struct {
	settings Settings
}

func New(settings any) (register.LinterPlugin, error) {
	// The configuration type will be map[string]any or []interface, it depends on your configuration.
	// You can use https://github.com/go-viper/mapstructure to convert map to struct.

	s, err := register.DecodeSettings[Settings](settings)
	if err != nil {
		return nil, err
	}

	return &Plugin{settings: s}, nil
}

func (p *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{Analyzer}, nil
}

func (p *Plugin) GetLoadMode() string {
	// NOTE: the mode can be `register.LoadModeSyntax` or `register.LoadModeTypesInfo`.
	// - `register.LoadModeSyntax`: if the linter doesn't use types information.
	// - `register.LoadModeTypesInfo`: if the linter uses types information.

	return register.LoadModeSyntax
}

var Analyzer = &analysis.Analyzer{
	Name: "aaa",
	Doc:  "checks unit test functions for // Arrange, // Act, and // Assert comments in correct order",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
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
