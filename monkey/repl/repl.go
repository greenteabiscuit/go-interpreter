package repl

import (
	"bufio"
	"fmt"
	"github.com/greenteabiscuit/go-interpreter/monkey/lexer"
	"github.com/greenteabiscuit/go-interpreter/monkey/token"
	"io"
)

const PROMPT = ">> "

func Start(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)

	for {
		fmt.Fprintf(w, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(w, "%+v\n", tok)
		}
	}
}
