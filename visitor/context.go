package visitor

import (
	"sysy/ast/types"
)

type Context struct {
	Symbols []map[string]types.Type
}

func NewContext() *Context {
	return &Context{
		Symbols: []map[string]types.Type{
			make(map[string]types.Type),
		},
	}
}

func (c *Context) DeclSymbol(Identifier string, Type types.Type) bool {
	if _, ok := c.Symbols[len(c.Symbols)-1][Identifier]; ok {
		return false
	}
	c.Symbols[len(c.Symbols)-1][Identifier] = Type
	return true
}

func (c *Context) GetSymbol(Identifier string) types.Type {
	for i := len(c.Symbols) - 1; i >= 0; i-- {
		if Type, ok := c.Symbols[i][Identifier]; ok {
			return Type
		}
	}
	return nil
}

func (c *Context) PushSymbolTable() {
	c.Symbols = append(c.Symbols, make(map[string]types.Type))
}

func (c *Context) PopSymbolTable() {
	c.Symbols = c.Symbols[:len(c.Symbols)-1]
}
