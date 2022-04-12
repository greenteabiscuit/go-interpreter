package lexer

import (
	"github.com/greenteabiscuit/go-interpreter/monkey/token"
	"testing"
)

func TestLexer_NextToken(t *testing.T) {
	input := "let x = 1; let y = 2;"
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
