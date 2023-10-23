package instructionsc3d

import (
	"Optimizar/environmentc3d"
)

type Print2 struct {
	Cadena string
	Tipo   string
	Op     string
}

func NewPrint2(cadena, tipo, op string) Print2 {
	return Print2{cadena, tipo, op}
}

func (v Print2) Ejecutar(ast *environmentc3d.AST) interface{} {
	ast.FinalCode.PushBack("\tprintf(" + v.Cadena + " , " + v.Tipo + v.Op + ");\n")
	return nil
}
