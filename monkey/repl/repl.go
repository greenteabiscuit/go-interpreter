package repl

import (
	"bufio"
	"fmt"
	"github.com/greenteabiscuit/go-interpreter/monkey/compiler"
	"github.com/greenteabiscuit/go-interpreter/monkey/vm"
	"io"

	"github.com/greenteabiscuit/go-interpreter/monkey/lexer"
	"github.com/greenteabiscuit/go-interpreter/monkey/parser"
)

// PROMPT ...
const PROMPT = ">> "

// Start ...
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	// not used when using compiler
	// env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		// io.WriteString(out, program.String())
		// io.WriteString(out, "\n")
		/*
			// unused when using compiler
			evaluated := evaluator.Eval(program, env)
			if evaluated != nil {
				io.WriteString(out, evaluated.Inspect())
				io.WriteString(out, "\n")
			}
		*/
		comp := compiler.New()
		if err := comp.Compile(program); err != nil {
			fmt.Fprintf(out, "woops, error")
			continue
		}

		machine := vm.New(comp.Bytecode())
		if err := machine.Run(); err != nil {
			fmt.Fprintf(out, "woops, error")
			continue
		}

		stackTop := machine.StackTop()
		io.WriteString(out, stackTop.Inspect())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
