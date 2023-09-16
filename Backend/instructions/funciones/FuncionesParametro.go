package funciones

import (
	"Backend/environment"
)

type FuncionesParametro struct {
	Lin            int
	Col            int
	ExternoInterno string
	Nombre         string
	Tipo           environment.TipoExpresion
	Inout          bool
	EI             bool
}

func NewFuncionesParametro(lin int, col int, externointerno, name string, tipo environment.TipoExpresion, inout, ei bool) FuncionesParametro {
	return FuncionesParametro{lin, col, externointerno, name, tipo, inout, ei}
}

/*
func (v FuncionesParametro) Ejecutar(ast *environment.AST) interface{} {
	symbol := environment.Symbol{
		Lin:   v.Lin,
		Col:   v.Col,
		Tipo:  v.Tipo,
		Valor: nil,
		Scope: ast.ObtenerAmbito(),
	}

	variablefuncion := environment.VariableFuncion{
		Name:           v.Nombre,
		Symbols:        symbol,
		Mutable:        true,
		TipoSimbolo:    "Funcion",
		Inout:          v.Inout,
		EI:             v.EI,
		ExternoInterno: v.ExternoInterno,
	}

	ast.Lista_Funciones_Var.PushBack(variablefuncion)
	return true
}
*/
