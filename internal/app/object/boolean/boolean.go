package boolean

import "pisang/internal/app/object"

type Boolean struct {
	Value interface{}
	Type  string
}

func New(t string, val interface{}) object.IObject {
	return &Boolean{
		Value: val,
		Type:  t,
	}
}

func (obj *Boolean) Assign(newObj object.IObject) {

}

func (obj *Boolean) Dispose() {
	obj = nil
}

func (obj *Boolean) Add(newObj object.IObject) {
	obj.Value = obj.Value.(int) + newObj.GetValue().(int)
}

func (obj *Boolean) Minus(newObj object.IObject) {
	obj.Value = obj.Value.(int) - newObj.GetValue().(int)
}

func (obj *Boolean) Asterisk(newObj object.IObject) {
	obj.Value = obj.Value.(int) * newObj.GetValue().(int)
}

func (obj *Boolean) Caret(newObj object.IObject) {
	obj.Value = obj.Value.(int) ^ newObj.GetValue().(int)
}

func (obj *Boolean) Slash(newObj object.IObject) {
	obj.Value = obj.Value.(int) / newObj.GetValue().(int)
}

func (obj *Boolean) And(newObj object.IObject) object.IObject {
	return nil
}

func (obj *Boolean) Or(newObj object.IObject) object.IObject {
	return nil
}

func (obj *Boolean) XOr(newObj object.IObject) object.IObject {
	return nil
}

func (obj *Boolean) Mod(newObj object.IObject) object.IObject {
	return nil
}

func (obj *Boolean) LShift(newObj object.IObject) object.IObject {
	return nil
}

func (obj *Boolean) RShift(newObj object.IObject) object.IObject {
	return nil
}

func (obj *Boolean) Gt(newObj object.IObject) object.IObject {
	return nil
}

func (obj *Boolean) Gte(newObj object.IObject) object.IObject {
	return nil
}

func (obj *Boolean) Lt(newObj object.IObject) object.IObject {
	return nil
}

func (obj *Boolean) Lte(newObj object.IObject) object.IObject {
	return nil
}

func (obj *Boolean) Eq(newObj object.IObject) object.IObject {
	if obj.Value == newObj.GetValue() {
		return &Boolean{
			Type:  "BOOL",
			Value: true,
		}
	}
	return &Boolean{
		Type:  "BOOL",
		Value: false,
	}
}

func (obj *Boolean) In(newObj object.IObject) object.IObject {
	return nil
}

func (obj *Boolean) Neq(newObj object.IObject) object.IObject {
	if obj.Value != newObj.GetValue() {
		return &Boolean{
			Type:  "BOOL",
			Value: true,
		}
	}
	return &Boolean{
		Type:  "BOOL",
		Value: false,
	}
}

func (obj *Boolean) GetValue() interface{} {
	return obj.Value
}

func (obj *Boolean) GetType() string {
	return "BOOL"
}
