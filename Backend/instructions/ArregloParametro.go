package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
)

type ArregloParametro struct {
	Op interfaces.Expression
}

func NewArregloParametro(op interfaces.Expression) ArregloParametro {
	exp := ArregloParametro{Op: op}
	return exp
}

func (o ArregloParametro) Ejecutar(ast *environment.AST) environment.Symbol {
	var op environment.Symbol
	op = o.Op.Ejecutar(ast)
	return environment.Symbol{Lin: op.Lin, Col: op.Col, Tipo: op.Tipo, Valor: op.Valor}
}
