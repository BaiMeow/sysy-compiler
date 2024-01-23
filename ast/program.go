package ast

type Program struct {
	// *FuncDef || *Decl
	Children []any `draw:"unfold"`
}

func (p *Program) Draw() DrawUnit {
	return draw(p)
}
