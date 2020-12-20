package evaluator

import "pisang/internal/pkg/token"

type IEval interface {
	EvalPostfix() int
	Eat(token.Token)
}
