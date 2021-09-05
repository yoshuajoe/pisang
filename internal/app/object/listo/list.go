package listo

import (
	"pisang/internal/app/object"
	"pisang/internal/app/object/boolean"
)

type List struct {
	Value interface{}
	Type  string
}

func New(t string, val interface{}) object.IObject {
	return &List{
		Value: val,
		Type:  t,
	}
}

func (obj *List) Assign(newObj object.IObject) {

}

func (obj *List) Dispose() {
	obj = nil
}

func (obj *List) Add(newObj object.IObject) {
	obj.Value = obj.Value.(int) + newObj.GetValue().(int)
}

func (obj *List) Minus(newObj object.IObject) {
	obj.Value = obj.Value.(int) - newObj.GetValue().(int)
}

func (obj *List) Asterisk(newObj object.IObject) {
	obj.Value = obj.Value.(int) * newObj.GetValue().(int)
}

func (obj *List) Caret(newObj object.IObject) {
	obj.Value = obj.Value.(int) ^ newObj.GetValue().(int)
}

func (obj *List) Slash(newObj object.IObject) {
	obj.Value = obj.Value.(int) / newObj.GetValue().(int)
}

func (obj *List) And(newObj object.IObject) object.IObject {
	return nil
}

func (obj *List) Or(newObj object.IObject) object.IObject {
	return nil
}

func (obj *List) XOr(newObj object.IObject) object.IObject {
	return nil
}

func (obj *List) Mod(newObj object.IObject) object.IObject {
	return &List{
		Type:  "List",
		Value: obj.Value.(int) % newObj.GetValue().(int),
	}
}

func (obj *List) LShift(newObj object.IObject) object.IObject {
	return nil
}

func (obj *List) RShift(newObj object.IObject) object.IObject {
	return nil
}

func (obj *List) Gt(newObj object.IObject) object.IObject {
	if obj.Value.(int) > newObj.GetValue().(int) {
		return &boolean.Boolean{
			Type:  "BOOL",
			Value: true,
		}
	}
	return &boolean.Boolean{
		Type:  "BOOL",
		Value: false,
	}
}

func (obj *List) Gte(newObj object.IObject) object.IObject {
	if obj.Value.(int) >= newObj.GetValue().(int) {

		return &boolean.Boolean{
			Type:  "BOOL",
			Value: true,
		}
	}
	return &boolean.Boolean{
		Type:  "BOOL",
		Value: false,
	}
}

func (obj *List) Lt(newObj object.IObject) object.IObject {
	if obj.Value.(int) < newObj.GetValue().(int) {
		return &boolean.Boolean{
			Type:  "BOOL",
			Value: true,
		}
	}
	return &boolean.Boolean{
		Type:  "BOOL",
		Value: false,
	}
}

func (obj *List) Lte(newObj object.IObject) object.IObject {
	if obj.Value.(int) <= newObj.GetValue().(int) {
		return &boolean.Boolean{
			Type:  "BOOL",
			Value: true,
		}
	}
	return &boolean.Boolean{
		Type:  "BOOL",
		Value: false,
	}
}

func (obj *List) Eq(newObj object.IObject) object.IObject {
	if obj.Value == newObj.GetValue() {
		return &boolean.Boolean{
			Type:  "BOOL",
			Value: true,
		}
	}
	return &boolean.Boolean{
		Type:  "BOOL",
		Value: false,
	}
}

func (obj *List) Neq(newObj object.IObject) object.IObject {
	if obj.Value != newObj.GetValue() {
		return &boolean.Boolean{
			Type:  "BOOL",
			Value: true,
		}
	}
	return &boolean.Boolean{
		Type:  "BOOL",
		Value: false,
	}
}

func (obj *List) GetValue() interface{} {
	return obj.Value
}

func (obj *List) GetType() string {
	return "List"
}
