package instructionsc3d

import (
	"Optimizar/environmentc3d"
)

type ArregloTemporales struct {
	Id_valido string
}

func NewArregloTemporales(id string) ArregloTemporales {
	return ArregloTemporales{id}
}

func (v ArregloTemporales) Ejecutar(ast *environmentc3d.AST) interface{} {
	tempo := environmentc3d.Temporal{
		Id: v.Id_valido,
	}
	ast.Lista_Temporales.PushBack(tempo)
	return nil
}
