package main

import (
	"pisang/internal/app/lexer"
	"pisang/internal/app/syntax"
	"testing"
)

func TestParentheses(t *testing.T) {
	code := "(2+3)\n"
	actual := 5

	lexer, lexerErr := lexer.New(code)
	interceptPanic(lexerErr)

	syntax, syntaxErr := syntax.New(lexer)
	interceptPanic(syntaxErr)

	v, exprErr := syntax.Expression()
	interceptPanic(exprErr)

	if actual != v.(int) {
		t.Errorf("Expected %v but got %v", v, actual)
	}
}

func TestStatement(t *testing.T) {
	code := "PROGRAM helloworld;"

	lexer, lexerErr := lexer.New(code)
	interceptPanic(lexerErr)

	syntax, syntaxErr := syntax.New(lexer)
	interceptPanic(syntaxErr)

	syntax.Program()
}
