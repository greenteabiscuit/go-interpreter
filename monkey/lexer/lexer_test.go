package lexer

import (
	"github.com/greenteabiscuit/go-interpreter/monkey/token"
	"testing"
)

func TestLexer_NextToken(t *testing.T) {
	input := "=+"
	tests := []struct {
		name            string
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{"equal", token.ASSIGN, "="},
		{"plus", token.PLUS, "+"},
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
