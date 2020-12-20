package lexer

import (
	"fmt"
	"pisang/internal/pkg/token"
	"strconv"
	"unicode"
)

type Lexer struct {
	line     string
	position int
	currChar byte
}

func New(line string) (ILexer, error) {
	return &Lexer{
		line:     line,
		position: 0,
		currChar: line[0],
	}, nil
}

func (lexer *Lexer) skipWhiteSpace() {
	for lexer.currChar != '\n' && lexer.currChar == ' ' {
		lexer.advance()
	}
}

func (lexer *Lexer) advance() {
	lexer.position += 1
	lexer.currChar = lexer.line[lexer.position]
}

func (lexer *Lexer) integer() (int, error) {
	var result string
	for unicode.IsDigit(rune(lexer.currChar)) {
		result = fmt.Sprintf("%s%s", result, string(lexer.currChar))
		lexer.advance()
	}
	i, e := strconv.Atoi(result)
	return i, e
}

func (lexer *Lexer) isSign() bool {
	if lexer.currChar == '+' || lexer.currChar == '-' {
		return true
	}
	return false
}

func (lexer *Lexer) GetNextToken() (token.Token, error) {
	if lexer.currChar == '\n' {
		return token.Token{
			Type:  "EOF",
			Value: '\n',
		}, nil
	} else if lexer.currChar == ' ' {
		lexer.skipWhiteSpace()
	} else if unicode.IsDigit(rune(lexer.currChar)) {
		i, e := lexer.integer()
		return token.Token{
			Type:  "INTEGER",
			Value: i,
		}, e
	} else if lexer.currChar == '+' {
		lexer.advance()
		return token.Token{
			Type:  "PLUS",
			Value: "+",
		}, nil
	} else if lexer.currChar == '-' {
		lexer.advance()
		return token.Token{
			Type:  "MINUS",
			Value: "-",
		}, nil
	}
	return token.Token{}, fmt.Errorf(fmt.Sprintf("Lexer Error: Invalid input at position: %v", lexer.position))
}

func (lexer *Lexer) GetLine() string {
	return lexer.line
}

func (lexer *Lexer) GetPosition() int {
	return lexer.position
}
