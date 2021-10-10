package object

type Object struct {
	Value interface{}
	Type  string
}

type IObject interface {
	Add(IObject)
	Minus(IObject)
	Asterisk(IObject)
	Caret(IObject)
	Slash(IObject)
	And(IObject) IObject
	Or(IObject) IObject
	XOr(IObject) IObject
	Mod(IObject) IObject
	LShift(IObject) IObject
	RShift(IObject) IObject
	Assign(IObject)
	Dispose()

	Gt(IObject) IObject
	Gte(IObject) IObject
	Lt(IObject) IObject
	Lte(IObject) IObject
	Eq(IObject) IObject
	In(IObject) IObject
	Neq(IObject) IObject
	GetValue() interface{}
	GetType() string
}
