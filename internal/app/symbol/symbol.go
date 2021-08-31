package symbol

import "pisang/internal/app/object"

type Symbol struct {
	list map[string]*object.Object
}

func New() Symbol {
	l := make(map[string]*object.Object)
	return Symbol{
		list: l,
	}
}

func (sym Symbol) Push(key string, val *object.Object) {
	sym.list[key] = val
}

func (sym Symbol) Get(key string) *object.Object {
	if val, ok := sym.list[key]; ok {
		return val
	}
	return nil
}
