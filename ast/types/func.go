package types

import (
	"fmt"
	"strings"
)

type Func struct {
	Return     Type
	ParamsList []Type
}

func (t *Func) Equal(p Type) bool {
	pp, ok := p.(*Func)
	if !ok {
		return false
	}
	if !t.Return.Equal(pp.Return) {
		return false
	}
	if len(t.ParamsList) != len(pp.ParamsList) {
		return false
	}
	for i := range t.ParamsList {
		if !t.ParamsList[i].Equal(pp.ParamsList[i]) {
			return false
		}
	}
	return true
}

func (t *Func) Assign(p Type) bool {
	return false
}

func (t *Func) String() string {
	params := make([]string, len(t.ParamsList))
	for i := range t.ParamsList {
		params[i] = t.ParamsList[i].String()
	}
	return fmt.Sprintf("func(%v) %v", strings.Join(params, ","), t.Return)
}
