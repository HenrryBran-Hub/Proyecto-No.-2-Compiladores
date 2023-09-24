package datosprimitivos

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"strconv"
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

func (v VariableDeclaracionSinTipo) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	if ast.ObtenerAmbito() == "Global" {
		gen.MainCodeT()
	}

	value := v.Value.Ejecutar(ast, gen)
	symbol := environment.Symbol{
		Lin:      v.Lin,
		Col:      v.Col,
		Tipo:     value.Type,
		Scope:    ast.ObtenerAmbito(),
		TipoDato: environment.VARIABLE,
		Posicion: ast.Lista_Variables.Len(),
		Valor:    value.Value,
	}
	Variable := environment.Variable{
		Name:        v.Name,
		Symbols:     symbol,
		Mutable:     true,
		TipoSimbolo: "Variable",
	}

	gen.AddComment("Declaracion de Variable")

	if value.Type == environment.BOOLEAN {
		gen.AddSetStack(strconv.Itoa(symbol.Posicion), value.Value)
		gen.AddBr()
	} else {
		//si es temp (num,string,etc)
		gen.AddSetStack(strconv.Itoa(symbol.Posicion), value.Value)
		gen.AddBr()
	}

	ast.GuardarVariable(Variable)
	gen.MainCodeF()
	return value
}
