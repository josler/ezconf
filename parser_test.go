package ezconf

import (
	"fmt"
	"testing"
)

func TestRegisterTransformerType(t *testing.T) {
	p := NewParser()
	filter := Filter{Type: String}
	p.RegisterTransformer(f, filter)
	result := p.Transform("boo")
	if result != "hi" {
		t.Errorf("result was not transformed was: %s", result)
	}
}

func TestRegisterTransformerKeyFields(t *testing.T) {
	p := NewParser()
	filter := Filter{KeyFields: []string{"Foo"}}
	p.RegisterTransformer(f, filter)
	result := p.Transform(simpleStruct{"boo", "bar"})
	fmt.Println(result)
	r := result.(simpleStruct)
	if r.Foo != "hi" {
		t.Errorf("result was not transformed was: %s", r.Foo)
	}
}

func TestRegisterTransformerTypeAndKey(t *testing.T) {
	p := NewParser()
	filter := Filter{Type: String, KeyFields: []string{"Foo"}}
	p.RegisterTransformer(f, filter)
	if p.keyTransforms["Foo"] != nil {
		t.Errorf("set a key transform when a type passed")
	}
	result := p.Transform(simpleStruct{"Foo", "bar"})
	if result.(simpleStruct).Foo != "hi" {
		t.Errorf("result was not transformed was: %s", result)
	}
}

func TestTransformNested(t *testing.T) {
	p := NewParser()
	filter := Filter{Type: String}
	p.RegisterTransformer(f, filter)
	result := p.Transform(nestedStruct{"Baz", simpleStruct{"Foo", "Bar"}})
	if result.(nestedStruct).Simple.Bar != "hi" {
		t.Errorf("result was not transformed")
	}
}

func TestTransformNestedKey(t *testing.T) {
	p := NewParser()
	filter := Filter{KeyFields: []string{"Foo"}}
	p.RegisterTransformer(f, filter)
	result := p.Transform(nestedStruct{"Baz", simpleStruct{"Foo", "Bar"}})
	if result.(nestedStruct).Simple.Foo != "hi" {
		t.Errorf("result was not transformed")
	}
}

type nestedStruct struct {
	Baz    string
	Simple simpleStruct
}

type simpleStruct struct {
	Foo string
	Bar string
}

func f(in interface{}) interface{} {
	return "hi"
}
