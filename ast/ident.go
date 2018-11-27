package ast

// Ident node represents a identifier.
type Ident struct {
	NamePos int
	Name    string
}

// NewIdent returns a new ident node.
func NewIdent(namePos int, name string) *Ident {
	return &Ident{
		NamePos: namePos,
		Name:    name,
	}
}

// Pos implementations for expression nodes.
func (e *Ident) Pos() int { return e.NamePos }

// End implementations for expression nodes.
func (e *Ident) End() int { return e.NamePos + len(e.Name) }

func (*Ident) exprNode() {}
