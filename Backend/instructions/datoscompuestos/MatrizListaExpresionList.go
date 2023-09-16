package datoscompuestos

import (
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

/*
func (o MatrizListaExpresionList) Ejecutar(ast *environment.AST) interface{} {
	o.Op1.Ejecutar(ast)
	o.Op2.Ejecutar(ast)
	return nil
}
*/
