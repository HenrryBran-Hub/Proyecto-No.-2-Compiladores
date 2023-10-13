package funciones

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
)

type FuncionesListaParametro struct {
	Lin            int
	Col            int
	ExternoInterno string
	Nombre         string
	Tipo           environment.TipoExpresion
	Inout          bool
	EI             bool
	Lista          interfaces.Instruction
}

func NewFuncionesListaParametro(lin int, col int, externointerno, name string, tipo environment.TipoExpresion, inout, ei bool, lista interfaces.Instruction) FuncionesListaParametro {
	return FuncionesListaParametro{lin, col, externointerno, name, tipo, inout, ei, lista}
}

func (v FuncionesListaParametro) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
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
	v.Lista.Ejecutar(ast, gen)
	return true
}
