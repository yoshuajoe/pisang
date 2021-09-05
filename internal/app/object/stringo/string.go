package stringo

import (
	"pisang/internal/app/object"
	"pisang/internal/app/object/boolean"
)

type String struct {
	Value interface{}
	Type  string
}

func New(t string, val interface{}) object.IObject {
	return &String{
		Value: val,
		Type:  t,
	}
}

func (obj *String) Assign(newObj object.IObject) {

}

func (obj *String) Dispose() {
	obj = nil
}

func (obj *String) Add(newObj object.IObject) {
	obj.Value = obj.Value.(string) + newObj.GetValue().(string)
}

func (obj *String) Minus(newObj object.IObject) {
}

func (obj *String) Asterisk(newObj object.IObject) {
	i := 0
	_obj := obj.Value.(string)
	for i < newObj.GetValue().(int) {
		_obj += _obj
	}
}

func (obj *String) Caret(newObj object.IObject) {
}

func (obj *String) Slash(newObj object.IObject) {
}

func (obj *String) And(newObj object.IObject) object.IObject {
	return nil
}

func (obj *String) Or(newObj object.IObject) object.IObject {
	return nil
}

func (obj *String) XOr(newObj object.IObject) object.IObject {
	return nil
}

func (obj *String) Mod(newObj object.IObject) object.IObject {
	return &String{
		Type:  "String",
		Value: obj.Value.(int) % newObj.GetValue().(int),
	}
}

func (obj *String) LShift(newObj object.IObject) object.IObject {
	return nil
}

func (obj *String) RShift(newObj object.IObject) object.IObject {
	return nil
}

func (obj *String) Gt(newObj object.IObject) object.IObject {
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

func (obj *String) Gte(newObj object.IObject) object.IObject {
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

func (obj *String) Lt(newObj object.IObject) object.IObject {
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

func (obj *String) Lte(newObj object.IObject) object.IObject {
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

func (obj *String) Eq(newObj object.IObject) object.IObject {
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

func (obj *String) Neq(newObj object.IObject) object.IObject {
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

func (obj *String) GetValue() interface{} {
	return obj.Value
}

func (obj *String) GetType() string {
	return "String"
}
