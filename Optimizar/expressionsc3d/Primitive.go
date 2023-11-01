package expressionsc3d

import (
	"Optimizar/environmentc3d"
)

type Primitive struct {
	Valor interface{}
	Tipo  environmentc3d.TipoExpresion
}

func NewPrimitive(valor interface{}, tipo environmentc3d.TipoExpresion) Primitive {
	exp := Primitive{valor, tipo}
	return exp
}

func (p Primitive) Ejecutar(ast *environmentc3d.AST) environmentc3d.Symbol {
	return environmentc3d.Symbol{
		Valor: p.Valor,
		Tipo:  p.Tipo,
	}
}
