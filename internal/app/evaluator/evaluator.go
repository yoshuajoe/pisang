package evaluator

import (
	"pisang/internal/app/object"
)

type Eval struct {
}

func New() (IEval, error) {
	return &Eval{}, nil
}

func (eval *Eval) Eat(a, b object.IObject, op string) object.IObject {
	if op == "+" {
		a.Add(b)
	} else if op == "-" {
		a.Minus(b)
	} else if op == "*" {
		a.Asterisk(b)
	} else if op == "/" {
		a.Slash(b)
	}
	return a
}
