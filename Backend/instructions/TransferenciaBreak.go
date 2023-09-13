package instructions

import (
	"Backend/environment"
)

type TransferenciaBreak struct {
	Lin int
	Col int
}

func NewTransferenciaBreak(lin int, col int) TransferenciaBreak {
	return TransferenciaBreak{Lin: lin, Col: col}
}

func (v TransferenciaBreak) Ejecutar(ast *environment.AST) interface{} {
	symbol := environment.Symbol{
		Lin:   v.Lin,
		Col:   v.Col,
		Tipo:  environment.BOOLEAN,
		Valor: true,
		Scope: ast.ObtenerAmbito(),
	}
	Variable := environment.Variable{
		Name:        "Break",
		Symbols:     symbol,
		Mutable:     false,
		TipoSimbolo: "Sentencia de Transferencia",
	}

	ast.GuardarVariable(Variable)
	return nil
}
