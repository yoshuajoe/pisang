package symbol

import "pisang/internal/app/object"

type Symbol struct {
	list map[string]object.IObject
}

func New() Symbol {
	l := make(map[string]object.IObject)
	return Symbol{
		list: l,
	}
}

func (sym Symbol) Push(key string, val object.IObject) {
	sym.list[key] = val
}

func (sym Symbol) Get(key string) object.IObject {
	if val, ok := sym.list[key]; ok {
		return val
	}
	return nil
}
