package datoscompuestos

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
)

type MatrizListaExpresionList struct {
	Op1 interfaces.Instruction
	Op2 interfaces.Instruction
}

func NewMatrizListaExpresionList(op1 interfaces.Instruction, op2 interfaces.Instruction) MatrizListaExpresionList {
	exp := MatrizListaExpresionList{op1, op2}
	return exp
}

func (o MatrizListaExpresionList) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	gen.AddComment("Datos Compuestos Matriz-Lista-Expresion-List")
	o.Op1.Ejecutar(ast, gen)
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	o.Op2.Ejecutar(ast, gen)
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}

	gen.AddBr()
	gen.MainCodeF()
	return nil
}
