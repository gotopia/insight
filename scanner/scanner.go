package scanner

import (
	"unicode"
	"unicode/utf8"

	"github.com/gotopia/insight/token"
)

// A Scanner holds the scanner's internal state while processing
// a given text. It can be allocated as part of another data
// structure but must be initialized via Init before use.
//
type Scanner struct {
	src      []byte
	ch       rune
	offset   int
	rdOffset int
	lhs      bool
}

// New returns a new scanner.
func New(src []byte) *Scanner {
	s := &Scanner{
		src: src,
		lhs: true,
	}
	return s
}

// Scan scans the next token and returns the token position, the token,
// and its literal string if applicable. The source end is indicated by
// token.EOF.
//
func (s *Scanner) Scan() (pos int, tok token.Token, lit string) {
	s.skipWhitespace()
	pos = s.offset
	ch := s.ch
	s.next()
	switch ch {
	case -1:
		tok = token.EOF
	case '=':
		switch s.ch {
		case '=':
			s.next()
			tok = token.EQL
		case '@':
			s.next()
			tok = token.CONTAIN
		case '~':
			s.next()
			tok = token.MATCH
		default:
			lit = string(ch)
		}
	case '!':
		switch s.ch {
		case '=':
			s.next()
			tok = token.NEQ
		case '@':
			s.next()
			tok = token.NOTCONTAIN
		case '~':
			s.next()
			tok = token.NOTMATCH
		default:
			lit = string(ch)
		}
	case '>':
		switch s.ch {
		case '=':
			s.next()
			tok = token.GEQ
		default:
			tok = token.GTR
		}
	case '<':
		switch s.ch {
		case '=':
			s.next()
			tok = token.LEQ
		default:
			tok = token.LSS
		}
	case ';':
		tok = token.AND
		s.lhs = true
	case ',':
		tok = token.OR
		s.lhs = true
	default:
		if s.lhs {
			lit = s.scanIdentifier()
			tok = token.IDENT
			s.lhs = false
		} else {
			lit = s.scanValue()
			tok = token.VALUE
			s.lhs = true
		}
	}
	return
}

func (s *Scanner) scanIdentifier() string {
	offs := s.offset
	for isLetter(s.ch) || isDigit(s.ch) {
		s.next()
	}
	return string(s.src[offs-1 : s.offset])
}

func (s *Scanner) scanValue() string {
	offs := s.offset
	for s.ch != ';' && s.ch != ',' && s.ch != -1 {
		s.next()
	}
	return string(s.src[offs-1 : s.offset])
}

func (s *Scanner) next() {
	if s.rdOffset < len(s.src) {
		s.offset = s.rdOffset
		r, w := rune(s.src[s.offset]), 1
		if r >= utf8.RuneSelf {
			r, w = utf8.DecodeRune(s.src[s.rdOffset:])
		}
		s.rdOffset += w
		s.ch = r
	} else {
		s.offset = len(s.src)
		s.ch = -1
	}
}

func (s *Scanner) skipWhitespace() {
	for s.ch == rune(0) || s.ch == ' ' || s.ch == '\t' || s.ch == '\n' || s.ch == '\r' {
		s.next()
	}
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch >= utf8.RuneSelf && unicode.IsLetter(ch)
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9' || ch >= utf8.RuneSelf && unicode.IsDigit(ch)
}
