package generator

import (
	"fmt"

	"github.com/gotopia/insight/token"
)

func buildSQL(x string, op token.Token, y string) string {
	switch op {
	case token.CONTAIN, token.NOTCONTAIN:
		return fmt.Sprintf("%v(%v, %v)", operators[op], x, y)
	case token.OR:
		return fmt.Sprintf("(%v %v %v)", x, operators[op], y)
	default:
		return fmt.Sprintf("%v %v %v", x, operators[op], y)
	}
}

type formatter func(x string, y string) string

var operators = map[token.Token]string{
	token.EQL:        "=",
	token.NEQ:        "<>",
	token.GTR:        ">",
	token.LSS:        "<",
	token.GEQ:        ">=",
	token.LEQ:        "<=",
	token.CONTAIN:    "INSTR",
	token.NOTCONTAIN: "!INSTR",
	token.MATCH:      "RLIKE",
	token.NOTMATCH:   "NOT RLIKE",
	token.AND:        "AND",
	token.OR:         "OR",
}
