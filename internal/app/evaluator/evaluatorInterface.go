package evaluator

import (
	"pisang/internal/app/object"
)

type IEval interface {
	Eat(a, b *object.Object, op string) *object.Object
}
