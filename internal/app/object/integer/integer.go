package integer

import (
	"pisang/internal/app/object"
	"pisang/internal/app/object/boolean"
)

type Integer struct {
	Value interface{}
	Type  string
}

func New(t string, val interface{}) object.IObject {
	return &Integer{
		Value: val,
		Type:  t,
	}
}

func (obj *Integer) Assign(newObj object.IObject) {

}

func (obj *Integer) Dispose() {
	obj = nil
}

func (obj *Integer) Add(newObj object.IObject) {
	obj.Value = obj.Value.(int) + newObj.GetValue().(int)
}

func (obj *Integer) Minus(newObj object.IObject) {
	obj.Value = obj.Value.(int) - newObj.GetValue().(int)
}

func (obj *Integer) Asterisk(newObj object.IObject) {
	obj.Value = obj.Value.(int) * newObj.GetValue().(int)
}

func (obj *Integer) Caret(newObj object.IObject) {
	obj.Value = obj.Value.(int) ^ newObj.GetValue().(int)
}

func (obj *Integer) Slash(newObj object.IObject) {
	obj.Value = obj.Value.(int) / newObj.GetValue().(int)
}

func (obj *Integer) And(newObj object.IObject) object.IObject {
	return nil
}

func (obj *Integer) Or(newObj object.IObject) object.IObject {
	return nil
}

func (obj *Integer) XOr(newObj object.IObject) object.IObject {
	return nil
}

func (obj *Integer) Mod(newObj object.IObject) object.IObject {
	return &Integer{
		Type:  "INTEGER",
		Value: obj.Value.(int) % newObj.GetValue().(int),
	}
}

func (obj *Integer) LShift(newObj object.IObject) object.IObject {
	return nil
}

func (obj *Integer) RShift(newObj object.IObject) object.IObject {
	return nil
}

func (obj *Integer) Gt(newObj object.IObject) object.IObject {
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

func (obj *Integer) Gte(newObj object.IObject) object.IObject {
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

func (obj *Integer) Lt(newObj object.IObject) object.IObject {
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

func (obj *Integer) Lte(newObj object.IObject) object.IObject {
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

func (obj *Integer) Eq(newObj object.IObject) object.IObject {
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

func (obj *Integer) In(newObj object.IObject) object.IObject {
	return nil
}

func (obj *Integer) Neq(newObj object.IObject) object.IObject {
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

func (obj *Integer) GetValue() interface{} {
	return obj.Value
}

func (obj *Integer) GetType() string {
	return "INTEGER"
}
