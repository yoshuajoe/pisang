package syntax

import (
	"fmt"
	"pisang/internal/app/lexer"
	"pisang/internal/pkg/token"
)

type Syntax struct {
	lexer     lexer.ILexer
	currToken token.Token
}

func New(lexer lexer.ILexer) (*Syntax, error) {
	return &Syntax{
		lexer: lexer,
	}, nil
}

func (syntax *Syntax) Expression() interface{} {
	t, err := syntax.lexer.GetNextToken()
	syntax.currToken = t
	if err != nil {
		return err
	}

	left := syntax.currToken
	syntax.verify("INTEGER")
	op := syntax.currToken
	syntax.verify("PLUS")
	right := syntax.currToken
	syntax.verify("INTEGER")

	result := 0
	if op.Type == "PLUS" {
		result = left.Value.(int) + right.Value.(int)
	}
	return result
}

func (syntax *Syntax) verify(tokenType string) error {
	fmt.Println(syntax.currToken)
	if syntax.currToken.Type == tokenType {
		t, err := syntax.lexer.GetNextToken()
		syntax.currToken = t
		if err != nil {
			return err
		}
	}
	return nil
}
