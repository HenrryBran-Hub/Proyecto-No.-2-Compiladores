package datoscompuestos

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
)

type ArregloParametro struct {
	Op interfaces.Expression
}

func NewArregloParametro(op interfaces.Expression) ArregloParametro {
	exp := ArregloParametro{Op: op}
	return exp
}

func (o ArregloParametro) Ejecutar(ast *environment.AST, gen *generator.Generator) environment.Value {
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	gen.AddComment("Datos Compuestos Arreglo-Parametro")
	op := o.Op.Ejecutar(ast, gen)
	gen.AddBr()
	gen.MainCodeF()
	return op
}
