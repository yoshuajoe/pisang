package token

import (
	"fmt"
)

type Token struct {
	Value interface{}
	Type  string
}

func New(value interface{}, type_ string) (IToken, error) {
	return &Token{
		Value: value,
		Type:  type_,
	}, nil
}

func (tok *Token) ToString() string {
	return fmt.Sprintf("Token(%v,%v)", tok.Value, tok.Type)
}
