package object

type Object struct {
	Value interface{}
	Type  string
}

func New(name string, t string, val interface{}) *Object {
	if t == "" {
		t = "_object"
	}
	return &Object{
		Value: val,
		Type:  t,
	}
}

func (obj *Object) Assign(newObj *Object) {
	obj = newObj
}

func (obj *Object) Dispose() {
	obj = nil
}
