package instructionsc3d

import (
	"Optimizar/environmentc3d"
)

type AcumuladorEncabezado struct {
	H string
	S string
}

func NewAcumuladorEncabezado(h, s string) AcumuladorEncabezado {
	return AcumuladorEncabezado{h, s}
}

func (v AcumuladorEncabezado) Ejecutar(ast *environmentc3d.AST) interface{} {
	ast.FinalCode.PushBack("#include <stdio.h>\n")
	ast.FinalCode.PushBack("double heap[" + v.H + "];\n")
	ast.FinalCode.PushBack("double stack[" + v.S + "];\n")
	ast.FinalCode.PushBack("double P;\n")
	ast.FinalCode.PushBack("double H;\n")
	return nil
}
