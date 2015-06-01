package ezconf

import "reflect"

// Parser provides mechanisms to parse and transform configuration structures.
type Parser struct {
	typeTransforms map[FilterType]TransformFunc
	keyTransforms  map[string]TransformFunc
}

// A Filter determines what parts of a configuration structure should be transformed.
type Filter struct {
	// Type indicates that all values matching a type should be parsed.
	Type FilterType
	// KeyFields indicates names of fields to be parsed.
	KeyFields []string
}

// FilterType matches a type of field to parse.
type FilterType int

// Various FilterTypes (just String currently).
const (
	_                 = iota
	String FilterType = iota + 1
)

// TransformFunc is the signature for a function to operate on an input and return an output.
type TransformFunc func(input interface{}) interface{}

// NewParser initializes a new Parser for use.
func NewParser() *Parser {
	return &Parser{keyTransforms: make(map[string]TransformFunc), typeTransforms: make(map[FilterType]TransformFunc)}
}

// RegisterTransformer registers a TransformFunc and a Filter to a Parser.
func (p *Parser) RegisterTransformer(f TransformFunc, filter Filter) {
	if filter.Type != 0 {
		p.typeTransforms[filter.Type] = f
		return
	}
	for i := range filter.KeyFields {
		p.keyTransforms[filter.KeyFields[i]] = f
	}
}

// Transform the input configuration and return the output.
// The output is a copy of the input.
func (p *Parser) Transform(conf interface{}) interface{} {
	original := reflect.ValueOf(conf)
	parsed := reflect.New(original.Type()).Elem()
	p.transformRecursive(parsed, original, nil)
	return parsed.Interface()
}

func (p *Parser) transformRecursive(parsed, original reflect.Value, transformer TransformFunc) {
	switch original.Kind() {
	case reflect.Ptr:
		originalValue := original.Elem()
		if !originalValue.IsValid() {
			return
		}
		parsed.Set(reflect.New(originalValue.Type()))
		p.transformRecursive(parsed.Elem(), originalValue, nil)
	case reflect.Struct:
		p.transformStruct(parsed, original, transformer)
	case reflect.String:
		p.transformString(parsed, original, transformer)
	default:
		if parsed.CanSet() {
			parsed.Set(original)
		}
	}
}

func (p *Parser) transformStruct(parsed, original reflect.Value, transformer TransformFunc) {
	for i := 0; i < original.NumField(); i++ {
		key := original.Type().Field(i).Name
		if v, ok := p.keyTransforms[key]; ok {
			p.transformRecursive(parsed.Field(i), original.Field(i), v)
		} else {
			p.transformRecursive(parsed.Field(i), original.Field(i), nil)
		}
	}
}

func (p *Parser) transformString(parsed, original reflect.Value, transformer TransformFunc) {
	result := original.Interface().(string)
	if transformer != nil {
		result = transformer(result).(string)
	} else if v, ok := p.typeTransforms[String]; ok {
		result = v(result).(string)
	}
	if parsed.CanSet() {
		parsed.SetString(result)
	}
}
