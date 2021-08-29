package syntax

import (
	"errors"
	"fmt"
	"pisang/internal/app/lexer"
	"pisang/internal/pkg/token"
)

type Symbol struct {
	Value interface{}
	Name  string
	Type  string
}

type Syntax struct {
	lexer       lexer.ILexer
	currToken   token.Token
	program     string
	symbolTable map[string]Symbol
}

func New(lexer lexer.ILexer) (*Syntax, error) {
	token, err := lexer.GetNextToken()
	intercept(err)

	symbolTable := make(map[string]Symbol)
	return &Syntax{
		lexer:       lexer,
		currToken:   token,
		symbolTable: symbolTable,
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
	val, _ := syntax.Variable()
	vval := fmt.Sprintf("%v", val)
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
	var _symVal int

	syntax.shouldBe("_id", syntax.currToken.Type)
	_symName = syntax.currToken.Value.(string)
	syntax.Fetch()
	syntax.shouldBe("ASSIGNMENT", syntax.currToken.Type)
	syntax.Fetch()

	_pseudoval, _ := syntax.Expression()
	_symVal = _pseudoval.(int)
	syntax.symbolTable[_symName] = Symbol{
		Value: _symVal,
		Name:  _symName,
		Type:  "int",
	}
	return nil, nil
}

func (syntax *Syntax) Variable() (interface{}, error) {
	/*
		variable : ID
	*/
	syntax.shouldBe("_id", syntax.currToken.Type)
	if val, ok := syntax.symbolTable[syntax.currToken.Value.(string)]; ok {
		syntax.Fetch()
		return val.Value, nil
	}
	return nil, nil
}

func (syntax *Syntax) Expression() (interface{}, error) {
	result := syntax.Term()
	for {
		if syntax.currToken.Type == "PLUS" || syntax.currToken.Type == "MINUS" {
			if syntax.currToken.Type == "PLUS" {
				syntax.Fetch()
				result += syntax.Term()
			} else if syntax.currToken.Type == "MINUS" {
				syntax.Fetch()
				result -= syntax.Term()
			} else if syntax.currToken.Type == "WHITESPACE" {
				syntax.Fetch()
			}
		} else {
			break
		}
	}
	return result, nil
}

func (syntax *Syntax) Term() int {
	result := syntax.Factor()
	for {
		if syntax.currToken.Type == "MULTIPLY" || syntax.currToken.Type == "DIVIDE" {
			if syntax.currToken.Type == "MULTIPLY" {
				syntax.Fetch()
				result *= syntax.Factor()
			} else if syntax.currToken.Type == "DIVIDE" {
				syntax.Fetch()
				result /= syntax.Factor()
			} else if syntax.currToken.Type == "WHITESPACE" {
				syntax.Fetch()
			}
		} else {
			break
		}
	}
	return result
}

func (syntax *Syntax) Factor() int {
	if syntax.currToken.Type == "INTEGER" {
		syntax.shouldBe("INTEGER", syntax.currToken.Type)
		i, _ := syntax.currToken.Value.(int)
		syntax.Fetch()
		return i
	} else if syntax.currToken.Type == "LPAREN" {
		syntax.shouldBe("LPAREN", syntax.currToken.Type)
		syntax.Fetch()
		i, syntaxErr := syntax.Expression()
		if syntaxErr != nil {
			panic(syntaxErr)
		}
		syntax.shouldBe("RPAREN", syntax.currToken.Type)
		syntax.Fetch()
		return i.(int)
	} else {
		variable, _ := syntax.Variable()
		return variable.(int)
	}
	return -100
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
