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

	v := syntax.Expression()

	if actual == v.GetValue().(int) {
		t.Fatalf("Actual %v got %v", actual, v.GetValue())
	}

}

func TestStatement(t *testing.T) {
	code := `b:=1;`

	lexer, lexerErr := lexer.New(code)
	interceptPanic(lexerErr)

	syntax, syntaxErr := syntax.New(lexer)
	interceptPanic(syntaxErr)

	syntax.Program()
}

func TestList(t *testing.T) {
	code := `b:=[1];a:=b[0];assert a;`

	lexer, lexerErr := lexer.New(code)
	interceptPanic(lexerErr)

	syntax, syntaxErr := syntax.New(lexer)
	interceptPanic(syntaxErr)

	syntax.Program()
}

func TestIf(t *testing.T) {
	code := `
		siswaDalamKelas := ["Aldo", "Aldi", "Andi", "Ali", "Ando"];
		index := 1;
		if siswaDalamKelas[index] == "Aldi"{
			assert "Halo " + siswaDalamKelas[index] + ", salam kenal!";
		}
	`

	lexer, lexerErr := lexer.New(code)
	interceptPanic(lexerErr)

	syntax, syntaxErr := syntax.New(lexer)
	interceptPanic(syntaxErr)

	syntax.Program()
}
