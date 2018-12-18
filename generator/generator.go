package generator

import (
	"sort"

	"github.com/gotopia/insight/ast"
	"github.com/gotopia/insight/token"
)

// Generator holds the generator's internal state.
type Generator struct {
	e      ast.Expr
	params []string
	args   []interface{}
}

// New returns a new generator.
func New(expr ast.Expr, params []string) *Generator {
	sort.Strings(params)
	return &Generator{
		e:      expr,
		params: params,
	}
}

// Generate a SQL clause and arguments from ast.
func (g *Generator) Generate() (string, []interface{}) {
	return g.generate(g.e), g.args
}

func (g *Generator) generate(e ast.Expr) string {
	switch te := e.(type) {
	case *ast.Ident:
		return te.Name
	case *ast.Value:
		g.args = append(g.args, te.Value)
		return "?"
	case *ast.BinaryExpr:
		if x, ok := te.X.(*ast.Ident); ok {
			if !g.isPermitted(x.Name) {
				return ""
			}
		}
		return g.build(g.generate(te.X), te.Op, g.generate(te.Y))
	default:
		return ""
	}
}

func (g *Generator) isPermitted(param string) bool {
	if len(g.params) == 0 {
		return true
	}
	i := sort.SearchStrings(g.params, param)
	return i < len(g.params) && g.params[i] == param
}

func (g *Generator) build(x string, op token.Token, y string) string {
	if x == "" || y == "" {
		return x + y
	}
	return buildSQL(x, op, y)
}
