package funciones

import (
	"Backend/environment"
)

type FuncionesDeclaracionR struct {
	Lin    int
	Col    int
	Name   string
	Tipo   environment.TipoExpresion
	Bloque []interface{}
}

func NewFuncionesDeclaracionR(lin int, col int, name string, tipo environment.TipoExpresion, bloque []interface{}) FuncionesDeclaracionR {
	return FuncionesDeclaracionR{Lin: lin, Col: col, Name: name, Tipo: tipo, Bloque: bloque}
}

/*
func (v FuncionesDeclaracionR) Ejecutar(ast *environment.AST) interface{} {
	funcion := environment.Funcion{
		Lin:           v.Lin,
		Col:           v.Col,
		Nombre:        v.Name,
		IsReturn:      true,
		IsParame:      false,
		CodigoFuncion: v.Bloque,
		Tipo:          v.Tipo,
	}

	ast.GuardarFuncion(funcion)
	ast.Lista_Funciones_Var.Init()
	return nil
}
*/
