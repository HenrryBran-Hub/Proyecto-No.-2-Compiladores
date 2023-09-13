package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
)

type MatrizListaNivel struct {
	Op interfaces.Instruction
}

func NewMatrizListaNivel(op1 interfaces.Instruction) MatrizListaNivel {
	exp := MatrizListaNivel{op1}
	return exp
}

func (o MatrizListaNivel) Ejecutar(ast *environment.AST) interface{} {
	ast.AumentarNivel()
	o.Op.Ejecutar(ast)
	return nil
}
