package datoscompuestos

import (
	"Backend/interfaces"
)

type MatrizDimension struct {
	Lin  int
	Col  int
	Type interfaces.Expression
}

func NewMatrizDimension(lin, col int, tipo interfaces.Expression) MatrizDimension {
	exp := MatrizDimension{Lin: lin, Col: col, Type: tipo}
	return exp
}

/*
func (o MatrizDimension) Ejecutar(ast *environment.AST) environment.Symbol {
	valor := o.Type.Ejecutar(ast)
	result := valor.Valor.(int) + 1
	return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: valor.Tipo, Valor: result}
}
*/
