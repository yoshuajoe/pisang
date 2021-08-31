package lexer

import (
	"fmt"
	"pisang/internal/pkg/token"
	"strconv"
	"unicode"
)

type Lexer struct {
	line             string
	position         int
	currChar         byte
	reservedKeywords map[string]token.Token
}

func New(line string) (ILexer, error) {
	reservedKeywords := map[string]token.Token{
		"PROGRAM": token.Token{
			Type:  "_id",
			Value: "PROGRAM",
		},
		"BEGIN": token.Token{
			Type:  "_id",
			Value: "BEGIN",
		},
		"END": token.Token{
			Type:  "_id",
			Value: "END",
		},
		"ASSERT": token.Token{
			Type:  "_id",
			Value: "ASSERT",
		},
	}

	return &Lexer{
		line:             line,
		position:         0,
		currChar:         line[0],
		reservedKeywords: reservedKeywords,
	}, nil
}

func (lexer *Lexer) IsReservedKeyword(key string) bool {
	if _, ok := lexer.reservedKeywords[key]; ok {
		return true
	}
	return false
}

func (lexer *Lexer) skipWhiteSpace() {
	for lexer.currChar == '\n' || lexer.currChar == ' ' || lexer.currChar == '\t' {
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

func (lexer *Lexer) string() (string, error) {
	var result string

	stringBound := 1
	lexer.advance()
	for {
		if stringBound < 2 {
			if lexer.currChar == '"' {
				stringBound += 1
			} else {
				result = fmt.Sprintf("%s%s", result, string(lexer.currChar))
			}
			lexer.advance()
		} else {
			break
		}
	}
	return result, nil
}

func (lexer *Lexer) _id() (token.Token, error) {
	var result string
	for unicode.IsLetter(rune(lexer.currChar)) || lexer.currChar == '_' {
		result = fmt.Sprintf("%s%s", result, string(lexer.currChar))
		lexer.advance()
	}

	if val, ok := lexer.reservedKeywords[result]; ok {
		return val, nil
	}

	return token.Token{
		Type:  "_id",
		Value: result,
	}, nil
}

func (lexer *Lexer) isSign() bool {
	if lexer.currChar == '+' || lexer.currChar == '-' {
		return true
	}
	return false
}

func (lexer *Lexer) peek() *string {
	peek_pos := lexer.position + 1
	if peek_pos > len(lexer.line)-1 {
		return nil
	} else {
		result := string(lexer.line[peek_pos])
		return &result
	}
}

func (lexer *Lexer) GetNextToken() (token.Token, error) {
	if lexer.currChar == '.' {
		return token.Token{
			Type:  "EOF",
			Value: '.',
		}, nil
	}
	if lexer.currChar == ' ' || lexer.currChar == '\n' || lexer.currChar == '\t' {
		lexer.skipWhiteSpace()
	}
	if lexer.currChar == '"' {
		s, e := lexer.string()
		return token.Token{
			Type:  "STRING",
			Value: s,
		}, e
	}
	if unicode.IsLetter(rune(lexer.currChar)) {
		return lexer._id()
	}
	if unicode.IsDigit(rune(lexer.currChar)) {
		i, e := lexer.integer()
		return token.Token{
			Type:  "INTEGER",
			Value: i,
		}, e
	}
	if lexer.currChar == '+' {
		lexer.advance()
		return token.Token{
			Type:  "PLUS",
			Value: "+",
		}, nil
	}
	if lexer.currChar == '-' {
		lexer.advance()
		return token.Token{
			Type:  "MINUS",
			Value: "-",
		}, nil
	}
	if lexer.currChar == '*' {
		lexer.advance()
		return token.Token{
			Type:  "MULTIPLY",
			Value: "*",
		}, nil
	}
	if lexer.currChar == '/' {
		lexer.advance()
		return token.Token{
			Type:  "DIVIDE",
			Value: "/",
		}, nil
	}
	if lexer.currChar == '(' {
		lexer.advance()
		return token.Token{
			Type:  "LPAREN",
			Value: "(",
		}, nil
	}
	if lexer.currChar == ')' {
		lexer.advance()
		return token.Token{
			Type:  "RPAREN",
			Value: ")",
		}, nil
	}
	if lexer.currChar == ':' && *lexer.peek() == "=" {
		lexer.advance()
		lexer.advance()
		return token.Token{
			Type:  "ASSIGNMENT",
			Value: ":=",
		}, nil
	}
	if lexer.currChar == '>' && *lexer.peek() == "=" {
		lexer.advance()
		lexer.advance()
		return token.Token{
			Type:  "GTE",
			Value: ">=",
		}, nil
	}
	if lexer.currChar == '<' && *lexer.peek() == "=" {
		lexer.advance()
		lexer.advance()
		return token.Token{
			Type:  "LTE",
			Value: "<=",
		}, nil
	}

	if lexer.currChar == ';' {
		lexer.advance()
		return token.Token{
			Type:  "SEMICOLON",
			Value: ";",
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
