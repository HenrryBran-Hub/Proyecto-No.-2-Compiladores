package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
)

type StruckFunction struct {
	Matuting bool
	Op       interfaces.Instruction
}

func NewStruckFunction(op interfaces.Instruction) StruckFunction {
	return StruckFunction{false, op}
}

func NewStruckFunctionMutating(op interfaces.Instruction) StruckFunction {
	return StruckFunction{true, op}
}

func (v StruckFunction) Ejecutar(ast *environment.AST) interface{} {

	v.Op.Ejecutar(ast)

	funcionstruc := ast.Lista_Funciones.Back().Value.(environment.Funcion)

	funcion2 := environment.Funcion{
		Lin:           funcionstruc.Lin,
		Col:           funcionstruc.Col,
		Nombre:        funcionstruc.Nombre,
		IsReturn:      funcionstruc.IsReturn,
		IsParame:      funcionstruc.IsParame,
		Parametros:    funcionstruc.Parametros,
		CodigoFuncion: funcionstruc.CodigoFuncion,
		Tipo:          funcionstruc.Tipo,
		Mutating:      v.Matuting,
	}

	ast.FuncionesStruct.PushBack(funcion2)
	ast.Lista_Funciones.Remove(ast.Lista_Funciones.Back())
	return nil
}
