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

func (syntax *Syntax) Expression() (interface{}, error) {
	syntax.Fetch()
	result := syntax.Term()
	for {
		if syntax.currToken.Type != "EOF" {
			syntax.Fetch()
			if syntax.currToken.Type == "PLUS" {
				syntax.Fetch()
				result += syntax.Term()
			} else if syntax.currToken.Type == "MINUS" {
				syntax.Fetch()
				result -= syntax.Term()
			}
		} else {
			break
		}
	}
	return result, nil
}

func (syntax *Syntax) Term() int {
	syntax.shouldBe("INTEGER", syntax.currToken.Type)
	i, _ := syntax.currToken.Value.(int)
	return i
}

func (syntax *Syntax) Fetch() {
	token, err := syntax.lexer.GetNextToken()
	intercept(err)
	syntax.currToken = token
}

func (syntax *Syntax) shouldBe(should string, given string) {
	if should != given {
		panic(fmt.Sprintf("Sytax error: should be %v yet given %v at position %v", should, given, syntax.lexer.GetPosition()))
	}
}

func intercept(p error) {
	if p != nil {
		panic(p)
	}
}
