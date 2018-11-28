package ast

// Value node represents a value node .
type Value struct {
	ValuePos int
	Value    string
}

// NewValue returns a new value node.
func NewValue(valuePos int, value string) *Value {
	return &Value{
		ValuePos: valuePos,
		Value:    value,
	}
}

// Pos implementations for value nodes.
func (e *Value) Pos() int { return e.ValuePos }

// End implementations for value nodes.
func (e *Value) End() int { return e.ValuePos + len(e.Value) }

func (*Value) exprNode() {}
