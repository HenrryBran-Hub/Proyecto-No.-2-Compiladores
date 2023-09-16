package datoscompuestos

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
)

type ArregloParametros struct {
	Lin int
	Col int
	Op  interfaces.Expression
}

func NewArregloParametros(lin int, col int, op interfaces.Expression) ArregloParametros {
	exp := ArregloParametros{Lin: lin, Col: col, Op: op}
	return exp
}

func (o ArregloParametros) Ejecutar(ast *environment.AST, gen *generator.Generator) environment.Value {
	var op environment.Value
	op = o.Op.Ejecutar(ast, gen)
	return op
}
