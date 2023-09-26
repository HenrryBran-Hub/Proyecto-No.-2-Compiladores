package sentencias

import (
	"Backend/environment"
	"Backend/generator"
)

type TransferenciaBreak struct {
	Lin int
	Col int
}

func NewTransferenciaBreak(lin int, col int) TransferenciaBreak {
	return TransferenciaBreak{Lin: lin, Col: col}
}

func (v TransferenciaBreak) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}

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
	gen.MainCodeF()
	return nil
}
