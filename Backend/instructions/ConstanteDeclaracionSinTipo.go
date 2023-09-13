package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
)

type ConstanteDeclaracionSinTipo struct {
	Lin   int
	Col   int
	Name  string
	Type  environment.TipoExpresion
	Value interfaces.Expression
}

func NewConstanteDeclaracionSinTipo(lin int, col int, name string, value interfaces.Expression) ConstanteDeclaracionSinTipo {
	return ConstanteDeclaracionSinTipo{Lin: lin, Col: col, Name: name, Value: value}
}

func (v ConstanteDeclaracionSinTipo) Ejecutar(ast *environment.AST) interface{} {
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
		Mutable:     false,
		TipoSimbolo: "Constante",
	}

	ast.GuardarVariable(Variable)
	return nil
}
