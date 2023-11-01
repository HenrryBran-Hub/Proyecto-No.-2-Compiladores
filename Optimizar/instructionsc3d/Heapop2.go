package instructionsc3d

import (
	"Optimizar/environmentc3d"
)

type Heapop2 struct {
	Id string
}

func NewHeapop2(id string) Heapop2 {
	return Heapop2{id}
}

func (v Heapop2) Ejecutar(ast *environmentc3d.AST) interface{} {
	ast.FinalCode.PushBack("\t" + v.Id + " = H ;\n")
	return nil
}
