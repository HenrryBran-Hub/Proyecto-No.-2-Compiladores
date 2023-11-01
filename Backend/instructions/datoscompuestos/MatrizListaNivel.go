package datoscompuestos

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
)

type MatrizListaNivel struct {
	Op interfaces.Instruction
}

func NewMatrizListaNivel(op1 interfaces.Instruction) MatrizListaNivel {
	exp := MatrizListaNivel{op1}
	return exp
}

func (o MatrizListaNivel) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	gen.AddComment("Datos Compuestos Matriz-Lista-Nivel")
	ast.AumentarNivel()
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	o.Op.Ejecutar(ast, gen)
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	gen.AddBr()
	gen.MainCodeF()
	return nil
}
