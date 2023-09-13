package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
)

type FuncionesLlamadaList3 struct {
	Lin   int
	Col   int
	Name  string
	Expr  interfaces.Expression
	Lista interfaces.Instruction
}

func NewFuncionesLlamadaList3(lin int, col int, name string, expr interfaces.Expression, lista interfaces.Instruction) FuncionesLlamadaList3 {
	return FuncionesLlamadaList3{lin, col, name, expr, lista}
}

func (v FuncionesLlamadaList3) Ejecutar(ast *environment.AST) interface{} {
	valor := v.Expr.Ejecutar(ast)
	symbol := environment.Symbol{
		Lin:   v.Lin,
		Col:   v.Col,
		Tipo:  valor.Tipo,
		Valor: valor.Valor,
		Scope: ast.ObtenerAmbito(),
	}

	variablefuncion := environment.VariableFuncion{
		Name:        v.Name,
		Symbols:     symbol,
		Mutable:     true,
		TipoSimbolo: "Funcion",
		Inout:       false,
		EI:          true,
	}

	ast.Lista_Funciones_Par.PushBack(variablefuncion)
	v.Lista.Ejecutar(ast)
	return true
}
