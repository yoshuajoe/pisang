package main

import (
	"bufio"
	"fmt"
	"os"
	"pisang/internal/app/lexer"
	"pisang/internal/app/syntax"
	"pisang/internal/pkg/tokentype"
)

func initialization() {
	tokenTypeTable, err := tokentype.New("silent")
	if err != nil {
		panic("Panic")
	}
	tokenTypeTable.Insert("INTEGER", 0)
	tokenTypeTable.Insert("PLUS", fmt.Sprintf("%v", '+'))
	tokenTypeTable.Insert("MINUS", fmt.Sprintf("%v", '-'))
	tokenTypeTable.Insert("EOF", fmt.Sprintf("%v", '\n'))
	tokenTypeTable.PrintAll()
}

func interceptPanic(p interface{}) {
	cvt, ok := p.(error)
	if ok {
		if p != nil {
			panic(cvt.Error)
		}
	}
}

func main() {
	initialization()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("pisang> ")
	text, inputErr := reader.ReadString('\n')
	interceptPanic(inputErr)

	lexer, lexerErr := lexer.New(text)
	interceptPanic(lexerErr)

	syntax, syntaxErr := syntax.New(lexer)
	interceptPanic(syntaxErr)

	v, exprErr := syntax.Expression()
	interceptPanic(exprErr)
	fmt.Println(v)
}
