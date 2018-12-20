package params

import (
	"github.com/deckarep/golang-set"
)

// Mapper returns the mapped key-value pair.
type Mapper func(v string) (key string, value interface{})

// Params holds the params' internal state.
type Params struct {
	permitted mapset.Set
	mappers   map[string]Mapper
}

// NewParams returns a new params.
func NewParams() *Params {
	return &Params{
		permitted: mapset.NewSet(),
		mappers:   map[string]Mapper{},
	}
}

// Permit a param.
func (p *Params) Permit(params ...string) {
	for _, param := range params {
		p.permitted.Add(param)
	}
}

// IsPermitted checks whether the param is permitted.
func (p *Params) IsPermitted(param string) bool {
	return p.permitted.Cardinality() == 0 || p.permitted.Contains(param)
}

// SetMapper sets a param mapper.
func (p *Params) SetMapper(key string, mapper Mapper) {
	p.mappers[key] = mapper
}

// Convert produces a new key-value pairs by corresponding mapper.
func (p *Params) Convert(key string, value string) (string, interface{}) {
	mapper, ok := p.mappers[key]
	if ok {
		return mapper(value)
	}
	return key, value
}
