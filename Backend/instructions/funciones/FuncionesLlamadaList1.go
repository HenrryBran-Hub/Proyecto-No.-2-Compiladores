package funciones

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
)

type FuncionesLlamadaList1 struct {
	Lin   int
	Col   int
	Name  string
	Lista interfaces.Instruction
}

func NewFuncionesLlamadaList1(lin int, col int, name string, lista interfaces.Instruction) FuncionesLlamadaList1 {
	return FuncionesLlamadaList1{lin, col, name, lista}
}

func (v FuncionesLlamadaList1) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {

	aux := ast.GetVariable(v.Name)
	symbol := environment.Symbol{
		Lin:   v.Lin,
		Col:   v.Col,
		Tipo:  aux.Symbols.Tipo,
		Valor: aux.Symbols.Valor,
		Scope: ast.ObtenerAmbito(),
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
	v.Lista.Ejecutar(ast, gen)
	return true
}
