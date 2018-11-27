package parser

import (
	"fmt"
	"strings"

	"github.com/gotopia/insight/token"
)

// ParsingError is an interface for parsing error.
type ParsingError interface {
	error
	Pos() int
}

type parsingError struct {
	code       int
	pos        int
	unexpected string
	expectings []token.Token
}

func (e *parsingError) Error() string {
	var expectingStrings []string
	for _, e := range e.expectings {
		expectingStrings = append(expectingStrings, e.String())
	}
	var expecting string
	if len(expectingStrings) > 0 {
		expecting = ", expecting " + strings.Join(expectingStrings, " or ")
	}
	return fmt.Sprintf("unexpected `%v`%v at position %v", e.unexpected, expecting, e.pos+1)
}

func (e *parsingError) Pos() int {
	return e.pos
}

func newParsingError(pos int, unexpected string, expectings ...token.Token) *parsingError {
	return &parsingError{
		pos:        pos,
		unexpected: unexpected,
		expectings: expectings,
	}
}
