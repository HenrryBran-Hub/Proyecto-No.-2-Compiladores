package instructionsc3d

import (
	"Optimizar/environmentc3d"
)

type Heapop4 struct {
	Numero string
}

func NewHeapop4(numero string) Heapop4 {
	return Heapop4{numero}
}

func (v Heapop4) Ejecutar(ast *environmentc3d.AST) interface{} {
	ast.FinalCode.PushBack("\tH = H + " + v.Numero + ";\n")
	return nil
}
