package instructionsc3d

import (
	"Optimizar/environmentc3d"
)

type Stackop2 struct {
	Id string
}

func NewStackop2(id string) Stackop2 {
	return Stackop2{id}
}

func (v Stackop2) Ejecutar(ast *environmentc3d.AST) interface{} {
	ast.FinalCode.PushBack("\t" + v.Id + " = P ;\n")
	return nil
}
