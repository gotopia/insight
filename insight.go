package insight

import (
	"strings"

	"github.com/gotopia/insight/generator"
	"github.com/gotopia/insight/params"
	"github.com/gotopia/insight/parser"
)

// Insight holds the insight's internal state.
type Insight struct {
	params *params.Params
}

func new() *Insight {
	return &Insight{
		params: params.NewParams(),
	}
}

// Permit returns a new insight instance with permitted params.
func Permit(params ...string) *Insight {
	return new().Permit(params...)
}

// Permit returns a new insight instance with permitted params.
func (i *Insight) Permit(params ...string) *Insight {
	i.params.Permit(params...)
	return i
}

// Map returns a new insight instance with params mappers.
func Map(key string, mapper params.Mapper) *Insight {
	return new().Map(key, mapper)
}

// Map returns a new insight instance with params mappers.
func (i *Insight) Map(key string, mapper params.Mapper) *Insight {
	i.params.SetMapper(key, mapper)
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

// OrderBy generates a SQL clause and arguments. If permitted params is not set, it means any parameters are permitted.
func (i *Insight) OrderBy(orderBy string) string {
	subs := strings.Split(orderBy, ",")
	clauses := []string{}
	for idx := 0; idx < len(subs); idx++ {
		clause := strings.TrimSpace(subs[idx])
		pairs := strings.SplitN(clause, " ", 2)
		old := pairs[0]
		if i.params.IsPermitted(old) {
			new, _ := i.params.Convert(old, "")
			clause = strings.Replace(clause, old, new, 1)
			clauses = append(clauses, clause)
		}
	}
	return strings.TrimSpace(strings.Join(clauses, ", "))
}
