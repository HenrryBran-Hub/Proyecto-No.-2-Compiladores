package instructions

import (
	"Backend/environment"
)

type FuncionesDeclaracion struct {
	Lin    int
	Col    int
	Name   string
	Bloque []interface{}
}

func NewFuncionesDeclaracion(lin int, col int, name string, bloque []interface{}) FuncionesDeclaracion {
	return FuncionesDeclaracion{lin, col, name, bloque}
}

func (v FuncionesDeclaracion) Ejecutar(ast *environment.AST) interface{} {

	funcion := environment.Funcion{
		Lin:           v.Lin,
		Col:           v.Col,
		Nombre:        v.Name,
		IsReturn:      false,
		IsParame:      false,
		CodigoFuncion: v.Bloque,
	}

	ast.GuardarFuncion(funcion)
	ast.Lista_Funciones_Var.Init()
	return nil
}
