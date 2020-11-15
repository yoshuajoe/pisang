package tokentype

type ITokenType interface {
	Insert(string, interface{}) int
	Lookup(string) int
	PrintAll()
}
