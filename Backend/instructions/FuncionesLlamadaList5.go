package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
)

type FuncionesLlamadaList5 struct {
	Lin  int
	Col  int
	Name string
	Expr interfaces.Expression
}

func NewFuncionesLlamadaList5(lin int, col int, name string, expr interfaces.Expression) FuncionesLlamadaList5 {
	return FuncionesLlamadaList5{lin, col, name, expr}
}

func (v FuncionesLlamadaList5) Ejecutar(ast *environment.AST) interface{} {
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
	return true
}
