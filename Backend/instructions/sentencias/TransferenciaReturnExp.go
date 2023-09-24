package sentencias

import (
	"Backend/interfaces"
)

type TransferenciaReturnExp struct {
	Lin   int
	Col   int
	Value interfaces.Expression
}

func NewTransferenciaReturnExp(lin int, col int, value interfaces.Expression) TransferenciaReturnExp {
	return TransferenciaReturnExp{Lin: lin, Col: col, Value: value}
}

/*
func (v TransferenciaReturnExp) Ejecutar(ast *environment.AST) interface{} {
	value := v.Value.Ejecutar(ast)
	symbol := environment.Symbol{
		Lin:   v.Lin,
		Col:   v.Col,
		Tipo:  value.Tipo,
		Valor: value.Valor,
		Scope: ast.ObtenerAmbito(),
	}
	Variable := environment.Variable{
		Name:        "ReturnExp",
		Symbols:     symbol,
		Mutable:     false,
		TipoSimbolo: "Sentencia de Transferencia",
	}

	ast.GuardarVariable(Variable)
	return nil
}
*/
