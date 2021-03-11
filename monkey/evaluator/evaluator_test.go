package evaluator

import (
	"testing"

	"github.com/greenteabiscuit/go-interpreter/monkey/lexer"
	"github.com/greenteabiscuit/go-interpreter/monkey/object"
	"github.com/greenteabiscuit/go-interpreter/monkey/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer")
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value")
		return false
	}
	return true
}