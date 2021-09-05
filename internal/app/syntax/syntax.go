package syntax

import (
	"fmt"
	"go/ast"
	"pisang/internal/app/evaluator"
	"pisang/internal/app/lexer"
	"pisang/internal/app/object"
	"pisang/internal/app/object/integer"
	"pisang/internal/app/object/listo"
	"pisang/internal/app/object/stringo"
	"pisang/internal/app/symbol"
	"pisang/internal/pkg/token"
)

type Syntax struct {
	lexer       lexer.ILexer
	currToken   token.Token
	ast         *ast.Node
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

func (syntax *Syntax) Program() (interface{}, error) {
	syntax.CompoundStatement()
	return nil, nil
}

func (syntax *Syntax) Assert() (interface{}, error) {
	syntax.shouldBe("_id", syntax.currToken.Type)
	syntax.Fetch()
	_var := syntax.Expression()
	vval := fmt.Sprintf("<type:%s> %v", _var.GetType(), _var.GetValue())
	fmt.Println(vval)
	return nil, nil
}

func (syntax *Syntax) ConditionStatement() (interface{}, error) {
	syntax.CompoundStatement()
	return nil, nil
}

func (syntax *Syntax) ConditionExpression() object.IObject {
	var returned object.IObject
	a := syntax.Expression()
	for {
		if syntax.currToken.Type == "EQ" {
			syntax.shouldBe("EQ", syntax.currToken.Type)
			syntax.Fetch()
			b := syntax.Expression()
			returned = a.Eq(b)
		} else if syntax.currToken.Type == "NEQ" {
			syntax.shouldBe("NEQ", syntax.currToken.Type)
			syntax.Fetch()
			b := syntax.Expression()
			returned = a.Eq(b)
		} else if syntax.currToken.Type == "GT" {
			syntax.shouldBe("GT", syntax.currToken.Type)
			syntax.Fetch()
			b := syntax.Expression()
			returned = a.Gt(b)
		} else if syntax.currToken.Type == "GTE" {
			syntax.shouldBe("GTE", syntax.currToken.Type)
			syntax.Fetch()
			b := syntax.Expression()
			returned = a.Gte(b)
		} else if syntax.currToken.Type == "LT" {
			syntax.shouldBe("LT", syntax.currToken.Type)
			syntax.Fetch()
			b := syntax.Expression()
			returned = a.Lt(b)
		} else if syntax.currToken.Type == "LTE" {
			syntax.shouldBe("LTE", syntax.currToken.Type)
			syntax.Fetch()
			b := syntax.Expression()
			returned = a.Lte(b)
		} else {
			break
		}
	}
	return returned
}

func (syntax *Syntax) IfStatement() (interface{}, error) {
	syntax.shouldBe("_id", syntax.currToken.Type)
	syntax.Fetch()
	condMustTrue := syntax.ConditionExpression()
	syntax.shouldBe("LCURLYBRACKET", syntax.currToken.Type)
	syntax.Fetch()
	if condMustTrue.GetValue().(bool) {
		syntax.ConditionStatement()
	}
	syntax.shouldBe("RCURLYBRACKET", syntax.currToken.Type)
	return nil, nil
}

func (syntax *Syntax) CompoundStatement() (interface{}, error) {
	syntax.StatementList()
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
	} else if syntax.currToken.Type == "_id" && syntax.currToken.Value == "assert" {
		return syntax.Assert()
	} else if syntax.currToken.Type == "_id" && syntax.currToken.Value == "if" {
		return syntax.IfStatement()
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

	syntax.symbolTable.Push(_symName, _pseudoval)

	return nil, nil
}

func (syntax *Syntax) Variable() object.IObject {
	/*
		variable : ID
	*/
	syntax.shouldBe("_id", syntax.currToken.Type)
	_id := syntax.currToken.Value.(string)
	_var := syntax.symbolTable.Get(_id)
	syntax.Fetch()
	for {
		if syntax.currToken.Type == "LBRACKET" {
			syntax.shouldBe("LBRACKET", syntax.currToken.Type)
			syntax.Fetch()
			syntax.shouldBe("INTEGER", syntax.currToken.Type)
			_index := syntax.currToken.Value
			__obj := _var.GetValue().([]object.IObject)
			_var = __obj[_index.(int)]
			syntax.Fetch()
			syntax.shouldBe("RBRACKET", syntax.currToken.Type)
			syntax.Fetch()
		} else {
			break
		}
	}
	return _var
}

func (syntax *Syntax) List() []object.IObject {
	/*
		variable : ID
	*/
	lst := []object.IObject{}
	lst = append(lst, syntax.Expression())
	for {
		if syntax.currToken.Type == "COMMA" {
			syntax.shouldBe("COMMA", syntax.currToken.Type)
			syntax.Fetch()
			lst = append(lst, syntax.Expression())
		} else {
			break
		}
	}

	return lst
}

func (syntax *Syntax) Expression() object.IObject {
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

func (syntax *Syntax) Term() object.IObject {
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

func (syntax *Syntax) Factor() object.IObject {
	if syntax.currToken.Type == "INTEGER" {
		syntax.shouldBe("INTEGER", syntax.currToken.Type)
		i, _ := syntax.currToken.Value.(int)
		syntax.Fetch()
		return integer.New("INTEGER", i)
	} else if syntax.currToken.Type == "LPAREN" {
		syntax.shouldBe("LPAREN", syntax.currToken.Type)
		syntax.Fetch()
		i := syntax.Expression()
		syntax.shouldBe("RPAREN", syntax.currToken.Type)
		syntax.Fetch()
		return i
	} else if syntax.currToken.Type == "LBRACKET" {
		syntax.shouldBe("LBRACKET", syntax.currToken.Type)
		syntax.Fetch()
		lst := syntax.List()
		syntax.shouldBe("RBRACKET", syntax.currToken.Type)
		syntax.Fetch()
		return listo.New("LIST", lst)
	} else if syntax.currToken.Type == "STRING" {
		_string := syntax.currToken.Value
		syntax.shouldBe("STRING", syntax.currToken.Type)
		syntax.Fetch()
		return stringo.New("STRING", _string)
	} else {
		_var := syntax.Variable()
		return _var
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
