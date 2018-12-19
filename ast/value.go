package ast

// Value node represents a value node .
type Value struct {
	ValuePos int
	Literal  string
	Value    interface{}
}

// NewValue returns a new value node.
func NewValue(valuePos int, literal string) *Value {
	return &Value{
		ValuePos: valuePos,
		Literal:  literal,
		Value:    literal,
	}
}

// Pos implementations for value nodes.
func (e *Value) Pos() int { return e.ValuePos }

// End implementations for value nodes.
func (e *Value) End() int { return e.ValuePos + len(e.Literal) }

func (*Value) exprNode() {}
