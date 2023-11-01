package expressionsc3d

import "Optimizar/environmentc3d"

type Operation struct {
	Lin int
}

func NewOperation(lin int) Operation {
	exp := Operation{Lin: lin}
	return exp
}

func (o Operation) Ejecutar(ast *environmentc3d.AST) environmentc3d.Symbol {

	return environmentc3d.Symbol{
		Valor: nil,
		Tipo:  environmentc3d.ID,
	}
}
