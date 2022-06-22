package vm

import (
	"fmt"
	"github.com/greenteabiscuit/go-interpreter/monkey/code"
	"github.com/greenteabiscuit/go-interpreter/monkey/compiler"
	"github.com/greenteabiscuit/go-interpreter/monkey/object"
)

const StackSize = 2048

type VM struct {
	constants    []object.Object
	instructions code.Instructions

	stack []object.Object
	sp    int
}

func New(bytecode *compiler.Bytecode) *VM {
	return &VM{
		instructions: bytecode.Instructions,
		constants:    bytecode.Constants,

		stack: make([]object.Object, StackSize),
		sp:    0,
	}
}

func (v *VM) StackTop() object.Object {
	if v.sp == 0 {
		return nil
	}
	return v.stack[v.sp-1]
}

func (v *VM) Run() error {
	for ip := 0; ip < len(v.instructions); ip++ {
		op := code.Opcode(v.instructions[ip])
		switch op {
		case code.OpConstant:
			constIndex := code.ReadUint16(v.instructions[ip+1:])
			ip += 2

			if err := v.push(v.constants[constIndex]); err != nil {
				return err
			}
		case code.OpAdd:
			right := v.pop()
			left := v.pop()
			leftValue, rightValue := left.(*object.Integer).Value, right.(*object.Integer).Value
			result := leftValue + rightValue
			v.push(&object.Integer{Value: result})
		case code.OpSub:
			right := v.pop()
			left := v.pop()
			leftValue, rightValue := left.(*object.Integer).Value, right.(*object.Integer).Value
			result := leftValue - rightValue
			v.push(&object.Integer{Value: result})
		case code.OpMul:
			right := v.pop()
			left := v.pop()
			leftValue, rightValue := left.(*object.Integer).Value, right.(*object.Integer).Value
			result := leftValue * rightValue
			v.push(&object.Integer{Value: result})
		case code.OpDiv:
			right := v.pop()
			left := v.pop()
			leftValue, rightValue := left.(*object.Integer).Value, right.(*object.Integer).Value
			result := leftValue / rightValue
			v.push(&object.Integer{Value: result})
		case code.OpPop:
			v.pop()
		}
	}
	return nil
}

func (v *VM) push(obj object.Object) error {
	if v.sp >= StackSize {
		return fmt.Errorf("stack overflow")
	}
	v.stack[v.sp] = obj
	v.sp++
	return nil
}

func (v *VM) pop() object.Object {
	o := v.stack[v.sp-1]
	v.sp--
	return o
}

func (v *VM) LastPoppedStackElem() object.Object {
	return v.stack[v.sp]
}
