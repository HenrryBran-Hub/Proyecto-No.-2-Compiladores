package instructionsc3d

import (
	"Optimizar/environmentc3d"
)

type Stackop4 struct {
	Op     string
	Numero string
}

func NewStackop4(op, numero string) Stackop4 {
	return Stackop4{op, numero}
}

func (v Stackop4) Ejecutar(ast *environmentc3d.AST) interface{} {
	ast.FinalCode.PushBack("\tP = P " + v.Op + " " + v.Numero + ";\n")
	return nil
}
