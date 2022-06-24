package vm

import (
	"fmt"
	"github.com/greenteabiscuit/go-interpreter/monkey/ast"
	"github.com/greenteabiscuit/go-interpreter/monkey/compiler"
	"github.com/greenteabiscuit/go-interpreter/monkey/lexer"
	"github.com/greenteabiscuit/go-interpreter/monkey/object"
	"github.com/greenteabiscuit/go-interpreter/monkey/parser"
	"testing"
)

// integrated testing
type testVMCases struct {
	input    string
	expected interface{}
}

func parse(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}

func runVMTests(t *testing.T, tests []testVMCases) {
	t.Helper()
	for _, tt := range tests {
		program := parse(tt.input)

		comp := compiler.New()
		if err := comp.Compile(program); err != nil {
			t.Fatalf("error compiling")
		}

		vm := New(comp.Bytecode())
		if err := vm.Run(); err != nil {
			t.Fatalf("error reading bytecode in vm")
		}

		//testExpectedObject(t, tt.expected, vm.StackTop())

		stackElem := vm.LastPoppedStackElem()
		testExpectedObject(t, tt.expected, stackElem)
	}
}

func testExpectedObject(t *testing.T, expected interface{}, actual object.Object) {
	t.Helper()

	switch expected := expected.(type) {
	case int:
		if err := testIntegerObject(int64(expected), actual); err != nil {
			t.Errorf("testIntegerObject failed: %s", err)
		}
	}
}

func testIntegerObject(expected int64, actual object.Object) error {
	val, ok := actual.(*object.Integer)
	if !ok {
		return fmt.Errorf("not integer object")
	}
	if val.Value != expected {
		return fmt.Errorf("values do not match: %d and %d", val.Value, expected)
	}
	return nil
}

func TestIntegerCases(t *testing.T) {
	tests := []testVMCases{
		{
			input:    "1 + 2",
			expected: 3,
		},
		{
			input:    "1 - 2",
			expected: -1,
		},
		{
			input:    "6 * 2",
			expected: 12,
		},
		{
			input:    "10 / 2",
			expected: 5,
		},
	}
	runVMTests(t, tests)
}
