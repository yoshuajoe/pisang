package syntax

import (
	"errors"
	"fmt"
	"pisang/internal/app/evaluator"
	"pisang/internal/app/lexer"
	"pisang/internal/app/object"
	"pisang/internal/app/symbol"
	"pisang/internal/pkg/token"
)

type Syntax struct {
	lexer       lexer.ILexer
	currToken   token.Token
	program     string
	symbolTable symbol.Symbol
	evaluator   evaluator.IEval
}

func New(lexer lexer.ILexer) (*Syntax, error) {
	token, err := lexer.GetNextToken()
	intercept(err)

	symbolTable := symbol.New()
	evaluator, _ := evaluator.New()
	return &Syntax{
		lexer:       lexer,
		currToken:   token,
		symbolTable: symbolTable,
		evaluator:   evaluator,
	}, nil
}

func (syntax *Syntax) InitProgram() (interface{}, error) {
	if syntax.currToken.Value == "PROGRAM" {
		syntax.shouldBe("_id", syntax.currToken.Type)
		syntax.Fetch()

		syntax.shouldBe("_id", syntax.currToken.Type)
		syntax.program = syntax.currToken.Value.(string)
		syntax.Fetch()

		syntax.shouldBe("SEMICOLON", syntax.currToken.Type)
		syntax.Fetch()
		return nil, nil
	}
	return nil, nil
}

func (syntax *Syntax) Program() (interface{}, error) {
	syntax.InitProgram()
	syntax.CompoundStatement()

	return nil, nil
}

func (syntax *Syntax) Assert() (interface{}, error) {
	syntax.shouldBe("_id", syntax.currToken.Type)
	syntax.Fetch()
	_var := syntax.Expression()
	vval := fmt.Sprintf("<type:%s> %v", _var.Type, _var.Value)
	fmt.Println(vval)
	syntax.Fetch()
	return nil, nil
}

func (syntax *Syntax) CompoundStatement() (interface{}, error) {
	/*
		compound_statement: BEGIN statement_list END
	*/
	if syntax.currToken.Value == "BEGIN" {
		syntax.shouldBe("_id", syntax.currToken.Type)
		syntax.Fetch()

		syntax.StatementList()

		if syntax.currToken.Value == "END" {
			syntax.shouldBe("_id", syntax.currToken.Type)
			syntax.Fetch()
		}
		return nil, nil
	}

	panic(errors.New("Syntax Error"))
	return nil, nil
}

func (syntax *Syntax) StatementList() (interface{}, error) {
	/*
	   statement_list : statement
	                      | statement SEMI statement_list
	*/

	syntax.Statement()
	for {
		if syntax.currToken.Type == "SEMICOLON" {
			syntax.shouldBe("SEMICOLON", syntax.currToken.Type)
			syntax.Fetch()
			syntax.Statement()
		} else {
			break
		}
	}
	return nil, nil
}

func (syntax *Syntax) Statement() (interface{}, error) {
	/*
		statement : compound_statement
				| assignment_statement
				| empty
	*/
	if syntax.currToken.Type == "_id" && !syntax.lexer.IsReservedKeyword(syntax.currToken.Value.(string)) {
		return syntax.AssignmetStatement()
	} else if syntax.currToken.Type == "_id" && syntax.currToken.Value == "ASSERT" {
		return syntax.Assert()
	}
	return nil, nil
}

func (syntax *Syntax) AssignmetStatement() (interface{}, error) {
	var _symName string
	syntax.shouldBe("_id", syntax.currToken.Type)
	_symName = syntax.currToken.Value.(string)
	syntax.Fetch()
	syntax.shouldBe("ASSIGNMENT", syntax.currToken.Type)
	syntax.Fetch()
	_pseudoval := syntax.Expression()

	_obj := object.New(_symName, _pseudoval.Type, _pseudoval.Value)
	syntax.symbolTable.Push(_symName, _obj)

	return nil, nil
}

func (syntax *Syntax) Variable() *object.Object {
	/*
		variable : ID
	*/
	syntax.shouldBe("_id", syntax.currToken.Type)
	_var := syntax.symbolTable.Get(syntax.currToken.Value.(string))
	syntax.Fetch()
	return _var
}

func (syntax *Syntax) Expression() *object.Object {
	result := syntax.Term()
	for {
		if syntax.currToken.Type == "PLUS" || syntax.currToken.Type == "MINUS" {
			if syntax.currToken.Type == "PLUS" {
				syntax.Fetch()
				result = syntax.evaluator.Eat(result, syntax.Term(), "+")
			} else if syntax.currToken.Type == "MINUS" {
				syntax.Fetch()
				result = syntax.evaluator.Eat(result, syntax.Term(), "-")
			} else if syntax.currToken.Type == "WHITESPACE" {
				syntax.Fetch()
			}
		} else {
			break
		}
	}
	return result
}

func (syntax *Syntax) Term() *object.Object {
	result := syntax.Factor()
	for {
		if syntax.currToken.Type == "MULTIPLY" || syntax.currToken.Type == "DIVIDE" {
			if syntax.currToken.Type == "MULTIPLY" {
				syntax.Fetch()
				result = syntax.evaluator.Eat(result, syntax.Factor(), "*")
			} else if syntax.currToken.Type == "DIVIDE" {
				syntax.Fetch()
				result = syntax.evaluator.Eat(result, syntax.Factor(), "/")
			} else if syntax.currToken.Type == "WHITESPACE" {
				syntax.Fetch()
			}
		} else {
			break
		}
	}
	return result
}

func (syntax *Syntax) Factor() *object.Object {
	if syntax.currToken.Type == "INTEGER" {
		syntax.shouldBe("INTEGER", syntax.currToken.Type)
		i, _ := syntax.currToken.Value.(int)
		syntax.Fetch()
		return &object.Object{
			Value: i,
			Type:  "INTEGER",
		}
	} else if syntax.currToken.Type == "LPAREN" {
		syntax.shouldBe("LPAREN", syntax.currToken.Type)
		syntax.Fetch()
		i := syntax.Expression()
		syntax.shouldBe("RPAREN", syntax.currToken.Type)
		syntax.Fetch()
		return i
	} else if syntax.currToken.Type == "STRING" {
		_string := syntax.currToken.Value
		syntax.shouldBe("STRING", syntax.currToken.Type)
		syntax.Fetch()
		return &object.Object{
			Value: _string,
			Type:  "STRING",
		}
	} else {
		_var := syntax.Variable()
		return &object.Object{
			Value: _var.Value,
			Type:  _var.Type,
		}
	}
	return nil
}

func (syntax *Syntax) Fetch() {
	token, err := syntax.lexer.GetNextToken()
	intercept(err)
	syntax.currToken = token
}

func (syntax *Syntax) shouldBe(should string, given string) {
	if should != given {
		panic(fmt.Sprintf("Syntax error: should be %v yet given %v at position %v", should, given, syntax.lexer.GetPosition()))
	}
}

func intercept(p error) {
	if p != nil {
		panic(p)
	}
}
