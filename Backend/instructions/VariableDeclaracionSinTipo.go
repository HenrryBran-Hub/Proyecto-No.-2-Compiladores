package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
)

type VariableDeclaracionSinTipo struct {
	Lin   int
	Col   int
	Name  string
	Type  environment.TipoExpresion
	Value interfaces.Expression
}

func NewVariableDeclaracionSinTipo(lin int, col int, name string, value interfaces.Expression) VariableDeclaracionSinTipo {
	return VariableDeclaracionSinTipo{Lin: lin, Col: col, Name: name, Value: value}
}

func (v VariableDeclaracionSinTipo) Ejecutar(ast *environment.AST) interface{} {
	value := v.Value.Ejecutar(ast)
	symbol := environment.Symbol{
		Lin:   value.Lin,
		Col:   value.Col,
		Tipo:  value.Tipo,
		Valor: value.Valor,
		Scope: ast.ObtenerAmbito(),
	}
	Variable := environment.Variable{
		Name:        v.Name,
		Symbols:     symbol,
		Mutable:     true,
		TipoSimbolo: "Variable",
	}

	ast.GuardarVariable(Variable)
	return nil
}
