package lexer

import "pisang/internal/pkg/token"

type ILexer interface {
	GetNextToken() (token.Token, error)
}
