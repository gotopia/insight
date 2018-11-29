package token

import "strconv"

// Token is the set of lexical tokens of the insight.
type Token int

// The list of tokens.
const (
	ILLEGAL Token = iota
	EOF

	literalBeg
	IDENT
	VALUE
	literalEnd

	operatorBeg
	EQL // ==
	NEQ // !=

	GTR // >
	LSS // <
	GEQ // >=
	LEQ // <=

	CONTAIN    // =@
	NOTCONTAIN // !@
	MATCH      // =~
	NOTMATCH   // !~

	AND // ;
	OR  // ,
	operatorEnd
)

// A set of constants for precedence-based expression parsing.
// Non-operators have lowest precedence, followed by operators
// starting with precedence 1 up to unary operators. The highest
// precedence serves as "catch-all" precedence for selector,
// indexing, and other operator and delimiter tokens.
//
const (
	LowestPrec  = 0
	HighestPrec = 4
)

// Precedence returns the operator precedence of the binary
// operator op. If op is not a binary operator, the result
// is LowestPrecedence.
//
func (tok Token) Precedence() int {
	switch tok {
	case AND:
		return 1
	case OR:
		return 2
	case EQL, NEQ, GTR, LSS, GEQ, LEQ, CONTAIN, NOTCONTAIN, MATCH, NOTMATCH:
		return 3
	}
	return LowestPrec
}

var tokens = []string{
	ILLEGAL:    "ILLEGAL",
	EOF:        "EOF",
	IDENT:      "IDENT",
	VALUE:      "VALUE",
	EQL:        "==",
	NEQ:        "!=",
	GTR:        ">",
	LSS:        "<",
	GEQ:        ">=",
	LEQ:        "<=",
	CONTAIN:    "=@",
	NOTCONTAIN: "!@",
	MATCH:      "=~",
	NOTMATCH:   "!~",
	AND:        ";",
	OR:         ",",
}

// String returns the string corresponding to the token tok.
func (tok Token) String() string {
	s := ""
	if 0 <= tok && tok < Token(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}

// IsLiteral returns true for tokens corresponding to identifiers
// and basic type literals; it returns false otherwise.
//
func (tok Token) IsLiteral() bool { return literalBeg < tok && tok < literalEnd }

// IsOperator returns true for tokens corresponding to operators and
// delimiters; it returns false otherwise.
//
func (tok Token) IsOperator() bool { return operatorBeg < tok && tok < operatorEnd }
