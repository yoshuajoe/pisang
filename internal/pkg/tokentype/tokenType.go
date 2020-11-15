package tokentype

import (
	"fmt"
)

type TokenType struct {
	Key    string
	Actual interface{}
}

type TableTokenType struct {
	Table []TokenType
	mode  string
}

func New(mode string) (ITokenType, error) {
	return &TableTokenType{
		mode: mode,
	}, nil
}

func (table *TableTokenType) Insert(key string, actual interface{}) int {
	table.Table = append(table.Table, TokenType{
		Key:    key,
		Actual: actual,
	})
	return len(table.Table) - 1
}

func (table *TableTokenType) Lookup(lookup string) int {
	for i, v := range table.Table {
		if v.Key == lookup {
			return i
		}
	}
	return -1
}

func (table *TableTokenType) PrintAll() {
	if table.mode == "verbose" {
		for i, v := range table.Table {
			fmt.Println(
				fmt.Sprintf("Index %d: Type({Key:%s, Actual:%v)", i, v.Key, v.Actual),
			)
		}
	}
}
