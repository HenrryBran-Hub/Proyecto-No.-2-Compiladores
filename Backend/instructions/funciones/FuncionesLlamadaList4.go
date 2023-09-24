package funciones

import (
	"Backend/interfaces"
)

type FuncionesLlamadaList4 struct {
	Lin   int
	Col   int
	Expr  interfaces.Expression
	Lista interfaces.Instruction
}

func NewFuncionesLlamadaList4(lin int, col int, expr interfaces.Expression, lista interfaces.Instruction) FuncionesLlamadaList4 {
	return FuncionesLlamadaList4{lin, col, expr, lista}
}

/*
func (v FuncionesLlamadaList4) Ejecutar(ast *environment.AST) interface{} {
	valor := v.Expr.Ejecutar(ast)
	symbol := environment.Symbol{
		Lin:   v.Lin,
		Col:   v.Col,
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
	v.Lista.Ejecutar(ast)
	return true
}
*/
