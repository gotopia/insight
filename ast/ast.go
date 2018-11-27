package ast

// Node is an interface for node.
type Node interface {
	Pos() int // position of first character belonging to the node
	End() int // position of first character immediately after the node
}

// Expr is an interface for expression.
type Expr interface {
	Node
	exprNode()
}
