package ast

import "github.com/gotopia/insight/token"

// BinaryExpr node represents a binary expr node.
type BinaryExpr struct {
	X     Expr        // left operand
	OpPos int         // position of Op
	Op    token.Token // operator
	Y     Expr        // right operand
}

// NewBinaryExpr returns a new binary expr node.
func NewBinaryExpr(x Expr, opPos int, op token.Token, y Expr) *BinaryExpr {
	return &BinaryExpr{
		X:     x,
		OpPos: opPos,
		Op:    op,
		Y:     y,
	}
}

// Pos implementations for expression nodes.
func (e *BinaryExpr) Pos() int { return e.X.Pos() }

// End implementations for expression nodes.
func (e *BinaryExpr) End() int { return e.Y.End() }

func (*BinaryExpr) exprNode() {}
