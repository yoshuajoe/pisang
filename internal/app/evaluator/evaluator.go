package evaluator

import "pisang/internal/pkg/token"

type BasicEval struct {
	tableToken []token.Token
}

func New() IEval {
	return &BasicEval{
		tableToken: []token.Token{},
	}
}

func (basic *BasicEval) EvalPostfix() int {
	return 0
}

func (basic *BasicEval) Eat(t token.Token) {
	basic.tableToken = append(basic.tableToken, t)
}

func precedence(op string) int {
	switch op {
	case "Exponential":
		return 3
	case "Multiplication":
	case "Division":
		return 2
	case "Addition":
	case "Subtraction":
		return 1
	}
	return -1
}
