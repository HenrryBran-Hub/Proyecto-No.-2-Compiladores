package instructionsc3d

import (
	"Optimizar/environmentc3d"
)

type Print1 struct {
	Cadena string
	Numero string
}

func NewPrint1(cadena, numero string) Print1 {
	return Print1{cadena, numero}
}

func (v Print1) Ejecutar(ast *environmentc3d.AST) interface{} {
	ast.FinalCode.PushBack("\tprintf(" + v.Cadena + " , " + v.Numero + ");\n")
	return nil
}
