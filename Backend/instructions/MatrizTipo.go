package instructions

import (
	"Backend/environment"
)

type MatrizTipo struct {
	Lin int
	Col int
	Op  environment.TipoExpresion
}

func NewMatrizTipo(lin, col int, op environment.TipoExpresion) MatrizTipo {
	exp := MatrizTipo{Lin: lin, Col: col, Op: op}
	return exp
}

func (o MatrizTipo) Ejecutar(ast *environment.AST) environment.Symbol {
	return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: o.Op, Valor: 1}
}
