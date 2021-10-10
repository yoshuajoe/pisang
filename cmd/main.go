package main

import (
	"bufio"
	"fmt"
	"os"
	"pisang/internal/app/lexer"
	"pisang/internal/app/syntax"
	"pisang/internal/pkg/tokentype"
	"strings"
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
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	lexer, lexerErr := lexer.New(strings.Join(lines, "\n"))
	interceptPanic(lexerErr)

	syntax, syntaxErr := syntax.New(lexer, nil, nil, nil, "")
	interceptPanic(syntaxErr)

	_, exprErr := syntax.Program()
	interceptPanic(exprErr)
}
