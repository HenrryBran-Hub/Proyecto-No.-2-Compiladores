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
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	gen.AddComment("Datos Compuestos Arreglo-Parametros")
	op := o.Op.Ejecutar(ast, gen)
	gen.AddBr()
	gen.MainCodeF()
	return op
}
