package instructionsc3d

import (
	"Optimizar/environmentc3d"
)

type Heapop1 struct {
	Numero string
}

func NewHeapop1(numero string) Heapop1 {
	return Heapop1{numero}
}

func (v Heapop1) Ejecutar(ast *environmentc3d.AST) interface{} {
	ast.FinalCode.PushBack("\tH = " + v.Numero + ";\n")
	return nil
}
