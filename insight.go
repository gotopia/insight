package insight

import (
	"strings"

	"github.com/gotopia/insight/generator"
	"github.com/gotopia/insight/parser"
)

// Insight holds the insight's internal state.
type Insight struct {
	params []string
}

func new() *Insight {
	return &Insight{
		params: []string{},
	}
}

// Permit returns a new insight instance with permitted params.
func Permit(params ...string) *Insight {
	return new().Permit(params...)
}

// Permit returns a new insight instance with permitted params.
func (i *Insight) Permit(params ...string) *Insight {
	i.params = append(i.params, params...)
	return i
}

// Filter generates a SQL clause and arguments. If permitted params is not set, it means any parameters are permitted.
func Filter(filter string) (clause string, args []interface{}, err error) {
	return new().Filter(filter)
}

// Filter generates a SQL clause and arguments. If permitted params is not set, it means any parameters are permitted.
func (i *Insight) Filter(filter string) (clause string, args []interface{}, err error) {
	if strings.TrimSpace(filter) == "" {
		return
	}
	expr, err := parser.New([]byte(filter)).Parse()
	if err != nil {
		return
	}
	clause, args = generator.New(expr, i.params).Generate()
	return
}
