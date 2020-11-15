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
	tokenTypeTable.Insert("EOF", fmt.Sprintf("%v", '\n'))
	tokenTypeTable.PrintAll()
}

func main() {
	initialization()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("pisang> ")
	text, inputErr := reader.ReadString('\n')
	if inputErr != nil {
		panic(inputErr)
	}

	lexer, lexerErr := lexer.New(text)
	if lexerErr != nil {
		panic(lexerErr)
	}

	syntax, syntaxErr := syntax.New(lexer)
	if syntaxErr != nil {
		panic(syntaxErr)
	}

	fmt.Println(syntax.Expression())
}
