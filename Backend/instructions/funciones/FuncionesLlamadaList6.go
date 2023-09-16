package funciones

import (
	"Backend/interfaces"
)

type FuncionesLlamadaList6 struct {
	Expr interfaces.Expression
}

func NewFuncionesLlamadaList6(expr interfaces.Expression) FuncionesLlamadaList6 {
	return FuncionesLlamadaList6{expr}
}

/*
func (v FuncionesLlamadaList6) Ejecutar(ast *environment.AST) interface{} {
	valor := v.Expr.Ejecutar(ast)
	symbol := environment.Symbol{
		Lin:   valor.Lin,
		Col:   valor.Col,
		Tipo:  valor.Tipo,
		Valor: valor.Valor,
		Scope: ast.ObtenerAmbito(),
	}

	variablefuncion := environment.VariableFuncion{
		Symbols:     symbol,
		Mutable:     true,
		TipoSimbolo: "Funcion",
		Inout:       false,
		EI:          false,
	}

	ast.Lista_Funciones_Par.PushBack(variablefuncion)
	return true
}
*/
