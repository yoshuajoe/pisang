package evaluator

import (
	"fmt"
	"pisang/internal/app/object"
)

type Eval struct {
}

func New() (IEval, error) {
	return &Eval{}, nil
}

func (eval *Eval) Eat(a, b *object.Object, op string) *object.Object {
	// integer
	if a.Type == "INTEGER" && b.Type == "INTEGER" {
		if op == "+" {
			result := a.Value.(int) + b.Value.(int)
			return &object.Object{
				Value: result,
				Type:  "INTEGER",
			}
		} else if op == "-" {
			result := a.Value.(int) - b.Value.(int)
			return &object.Object{
				Value: result,
				Type:  "INTEGER",
			}
		} else if op == "*" {
			result := a.Value.(int) * b.Value.(int)
			return &object.Object{
				Value: result,
				Type:  "INTEGER",
			}
		} else if op == "/" {
			result := a.Value.(int) / b.Value.(int)
			return &object.Object{
				Value: result,
				Type:  "INTEGER",
			}
		}
	} else if a.Type == "STRING" && b.Type == "STRING" {
		if op == "+" {
			result := fmt.Sprintf("%s%s", a.Value.(string), b.Value.(string))
			return &object.Object{
				Value: result,
				Type:  "STRING",
			}
		} else if op == "-" {
			result := a.Value.(int) - b.Value.(int)
			return &object.Object{
				Value: result,
				Type:  "INTEGER",
			}
		} else if op == "*" {
			result := a.Value.(int) * b.Value.(int)
			return &object.Object{
				Value: result,
				Type:  "INTEGER",
			}
		} else if op == "/" {
			result := a.Value.(int) / b.Value.(int)
			return &object.Object{
				Value: result,
				Type:  "INTEGER",
			}
		}
	}
	return nil
}
