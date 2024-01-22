package ast

import (
	"fmt"
	"sysy/ast/types"
)

func BaseTypeDefault(p types.Type) (Expr, error) {
	switch p := p.(type) {
	case *types.Base:
		switch p.Type {
		case types.Int:
			return &ConstExpression{
				Value: 0,
			}, nil
		case types.Float:
			return &ConstExpression{
				Value: 0.0,
			}, nil
		case types.Void:
			return &ConstExpression{
				Value: nil,
			}, nil
		}
	}
	return nil, fmt.Errorf("not base type")
}
