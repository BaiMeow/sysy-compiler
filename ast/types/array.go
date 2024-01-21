package types

import (
	"slices"
	"strconv"
	"strings"
)

type Array struct {
	ElementType Type
	Shape       []int
}

func (t *Array) Equal(p Type) bool {
	pp, ok := p.(*Array)
	if !ok {
		return false
	}
	return pp.ElementType.Equal(t.ElementType) && slices.Equal(pp.Shape, t.Shape)
}

func (t *Array) Assign(p Type) bool {
	pp, ok := p.(*Array)
	if !ok {
		return false
	}
	return t.ElementType.Assign(pp.ElementType)
}

func (t *Array) String() string {
	var sb strings.Builder
	sb.WriteString(t.ElementType.String())
	for i := 0; i < len(t.Shape); i++ {
		sb.WriteByte('[')
		sb.WriteString(strconv.Itoa(t.Shape[i]))
		sb.WriteByte(']')
	}
	return sb.String()
}
