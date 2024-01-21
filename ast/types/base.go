package types

type Base struct {
	Type  string
	Const bool
}

const (
	Int   = "int"
	Float = "float"
	Void  = "void"
)

func (t *Base) Equal(p Type) bool {
	pp, ok := p.(*Base)
	if !ok {
		return false
	}
	return *pp == *t
}

func (t *Base) Assign(p Type) bool {
	if t.Const {
		return false
	}
	pp, ok := p.(*Base)
	if !ok {
		return false
	}
	return t.Type == pp.Type
}

func (t *Base) String() string {
	if t.Const {
		return "const " + t.Type
	}
	return t.Type
}
