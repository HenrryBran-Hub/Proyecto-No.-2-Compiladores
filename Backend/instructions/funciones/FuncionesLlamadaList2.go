package funciones

import (
	"Backend/environment"
	"Backend/generator"
)

type FuncionesLlamadaList2 struct {
	Lin  int
	Col  int
	Name string
}

func NewFuncionesLlamadaList2(lin int, col int, name string) FuncionesLlamadaList2 {
	return FuncionesLlamadaList2{lin, col, name}
}

func (v FuncionesLlamadaList2) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {

	aux := ast.GetVariable(v.Name)

	symbol := environment.Symbol{
		Lin:         v.Lin,
		Col:         v.Col,
		Tipo:        aux.Symbols.Tipo,
		Valor:       aux.Symbols.Valor,
		ValorInt:    aux.Symbols.ValorInt,
		ValorFloat:  aux.Symbols.ValorFloat,
		ValorString: aux.Symbols.ValorString,
		Scope:       ast.ObtenerAmbito(),
	}

	variablefuncion := environment.VariableFuncion{
		Name:        v.Name,
		Symbols:     symbol,
		Mutable:     true,
		TipoSimbolo: "Funcion",
		Inout:       true,
		EI:          false,
	}

	ast.Lista_Funciones_Par.PushBack(variablefuncion)
	return true
}
