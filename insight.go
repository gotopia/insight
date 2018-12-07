package insight

import (
	"strings"

	"github.com/gotopia/insight/generator"
	"github.com/gotopia/insight/parser"
)

// Filter generates a SQL query and arguments.
func Filter(filter string) (query string, args []string, err error) {
	if strings.TrimSpace(filter) == "" {
		return
	}
	expr, err := parser.New([]byte(filter)).Parse()
	if err != nil {
		return
	}
	return generator.New(expr).Generate()
}
