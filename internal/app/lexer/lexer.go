package lexer

import (
	"errors"
	"fmt"
	"pisang/internal/pkg/token"
	"strconv"
	"unicode"
)

type Lexer struct {
	line     string
	position int
}

func New(line string) (ILexer, error) {
	return &Lexer{
		line:     line,
		position: 0,
	}, nil
}

func (lexer *Lexer) GetNextToken() (token.Token, error) {
	currInput := lexer.line[lexer.position]

	if currInput == '\n' {
		lexer.position += 1
		return token.Token{
			Type:  "EOF",
			Value: '\n',
		}, nil
	} else if unicode.IsDigit(rune(currInput)) {
		lexer.position += 1
		converted, convertedErr := strconv.Atoi(string(currInput))
		return token.Token{
			Type:  "INTEGER",
			Value: converted,
		}, convertedErr
	} else if currInput == '+' {
		lexer.position += 1
		return token.Token{
			Type:  "PLUS",
			Value: "+",
		}, nil
	}
	return token.Token{}, errors.New(fmt.Sprintf("Lexer Error: Invalid input at position: %v", lexer.position))
}
