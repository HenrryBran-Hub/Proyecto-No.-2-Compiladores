package instructionsc3d

import (
	"Optimizar/environmentc3d"
)

type Stackop1 struct {
	Numero string
}

func NewStackop1(numero string) Stackop1 {
	return Stackop1{numero}
}

func (v Stackop1) Ejecutar(ast *environmentc3d.AST) interface{} {
	ast.FinalCode.PushBack("\tP = " + v.Numero + ";\n")
	return nil
}
