package instructions

import (
	"Backend/environment"
)

type TransferenciaReturn struct {
	Lin int
	Col int
}

func NewTransferenciaReturn(lin int, col int) TransferenciaReturn {
	return TransferenciaReturn{Lin: lin, Col: col}
}

func (v TransferenciaReturn) Ejecutar(ast *environment.AST) interface{} {
	symbol := environment.Symbol{
		Lin:   v.Lin,
		Col:   v.Col,
		Tipo:  environment.BOOLEAN,
		Valor: true,
		Scope: ast.ObtenerAmbito(),
	}
	Variable := environment.Variable{
		Name:        "Return",
		Symbols:     symbol,
		Mutable:     false,
		TipoSimbolo: "Sentencia de Transferencia",
	}

	ast.GuardarVariable(Variable)
	return nil
}
