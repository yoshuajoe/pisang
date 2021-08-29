package lexer

import "pisang/internal/pkg/token"

type ILexer interface {
	GetLine() string
	GetPosition() int
	GetNextToken() (token.Token, error)
	IsReservedKeyword(string) bool
}
