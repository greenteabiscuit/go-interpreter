package lexer

import (
	"github.com/greenteabiscuit/go-interpreter/monkey/token"
	"testing"
)

func TestLexer_NextToken(t *testing.T) {
	input := `let x = 1; let y = 2;
	let add = fn(w, z) {
		w + z;
	}
	!*/<>`
	tests := []struct {
		name            string
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{"let", token.LET, "let"},
		{"x", token.IDENT, "x"},
		{"=", token.ASSIGN, "="},
		{"1", token.INT, "1"},
		{"semicolon", token.SEMICOLON, ";"},
		{"let", token.LET, "let"},
		{"y", token.IDENT, "y"},
		{"=", token.ASSIGN, "="},
		{"2", token.INT, "2"},
		{"semicolon", token.SEMICOLON, ";"},
		{"let", token.LET, "let"},
		{"add", token.IDENT, "add"},
		{"=", token.ASSIGN, "="},
		{"fn", token.FUNCTION, "fn"},
		{"(", token.LPAREN, "("},
		{"w", token.IDENT, "w"},
		{",", token.COMMA, ","},
		{"z", token.IDENT, "z"},
		{")", token.RPAREN, ")"},
		{"{", token.LBRACE, "{"},
		{"w", token.IDENT, "w"},
		{"+", token.PLUS, "+"},
		{"z", token.IDENT, "z"},
		{"semicolon", token.SEMICOLON, ";"},
		{"}", token.RBRACE, "}"},
		// !*/<>
		{"bang", token.BANG, "!"},
		{"asterisk", token.ASTERISK, "*"},
		{"slash", token.SLASH, "/"},
		{"less than", token.LT, "<"},
		{"greater than", token.GT, ">"},
	}

	l := New(input)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := l.NextToken()
			if got.Type != tt.expectedType {
				t.Errorf("NextToken() = %v, want %v", got.Type, tt.expectedType)
			}
			if got.Literal != tt.expectedLiteral {
				t.Errorf("NextToken() = %v, want %v", got.Literal, tt.expectedLiteral)
			}
		})
	}
}
