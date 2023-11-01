package instructionsc3d

import (
	"Optimizar/environmentc3d"
)

type Heapop3 struct {
	Tipo   string
	Numero string
}

func NewHeapop3(tipo, numero string) Heapop3 {
	return Heapop3{tipo, numero}
}

func (v Heapop3) Ejecutar(ast *environmentc3d.AST) interface{} {
	ast.FinalCode.PushBack("\theap[" + v.Tipo + "H] = " + v.Numero + ";\n")
	return nil
}
