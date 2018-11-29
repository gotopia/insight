package generator

import (
	"github.com/gotopia/insight/ast"
)

// Generator holds the generator's internal state.
type Generator struct {
	e    ast.Expr
	args []string
}

// New returns a new generator.
func New(expr ast.Expr) *Generator {
	return &Generator{
		e: expr,
	}
}

// Generate generates a SQL query and arguments from ast.
func (g *Generator) Generate() (query string, args []string, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = er.(error)
		}
	}()
	query = g.generate(g.e)
	args = g.args
	return
}

func (g *Generator) generate(e ast.Expr) (query string) {
	switch te := e.(type) {
	case *ast.Ident:
		query = te.Name
	case *ast.Value:
		g.args = append(g.args, te.Value)
		query = "?"
	case *ast.BinaryExpr:
		query = buildSQL(g.generate(te.X), te.Op, g.generate(te.Y))
	default:
	}
	return
}
