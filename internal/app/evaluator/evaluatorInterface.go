package evaluator

import (
	"pisang/internal/app/object"
)

type IEval interface {
	Eat(a, b object.IObject, op string) object.IObject
}
