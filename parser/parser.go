package parser

import (
	"github.com/gotopia/insight/ast"
	"github.com/gotopia/insight/scanner"
	"github.com/gotopia/insight/token"
)

// Parser holds the parser's internal state.
type Parser struct {
	scanner *scanner.Scanner
	pos     int
	tok     token.Token
	lit     string
}

// New returns a new parser.
func New(src []byte) *Parser {
	scanner := scanner.New(src)
	p := &Parser{
		scanner: scanner,
	}
	p.next()
	return p
}

// Parse the source code in parser.
func (p *Parser) Parse() (e ast.Expr, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = er.(error)
		}
	}()
	e = p.parseBinaryExpr(token.LowestPrec + 1)
	return
}

func (p *Parser) parseBinaryExpr(prec int) ast.Expr {
	x := p.parsePrimaryExpr()
	for {
		pos := p.pos
		op := p.tok
		oprec := op.Precedence()
		if oprec < prec {
			return x
		}
		p.next()
		var y ast.Expr
		y = p.parseBinaryExpr(oprec + 1)
		x = newBinaryExpr(x, pos, op, y)
	}
}

func (p *Parser) parsePrimaryExpr() ast.Expr {
	switch p.tok {
	case token.IDENT:
		x := ast.NewIdent(p.pos, p.lit)
		p.next()
		return x
	case token.VALUE:
		x := ast.NewValue(p.pos, p.lit)
		p.next()
		return x
	case token.AND, token.OR, token.EOF:
		return ast.NewValue(p.pos, "")
	default:
		err := newParsingError(p.pos, p.tok.String(), token.IDENT)
		panic(err)
	}
}

func newBinaryExpr(x ast.Expr, pos int, op token.Token, y ast.Expr) *ast.BinaryExpr {
	if op == token.AND || op == token.OR {
		_, xok := x.(*ast.BinaryExpr)
		_, yok := y.(*ast.BinaryExpr)
		if !(xok && yok) {
			err := newParsingError(pos, op.String())
			panic(err)
		}
	}
	return ast.NewBinaryExpr(x, pos, op, y)
}

func (p *Parser) next() {
	p.pos, p.tok, p.lit = p.scanner.Scan()
	if p.tok == token.ILLEGAL {
		err := newParsingError(p.pos, p.lit)
		panic(err)
	}
}
