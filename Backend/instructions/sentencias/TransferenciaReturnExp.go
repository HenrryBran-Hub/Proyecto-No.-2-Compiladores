package sentencias

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"strconv"
)

type TransferenciaReturnExp struct {
	Lin   int
	Col   int
	Value interfaces.Expression
}

func NewTransferenciaReturnExp(lin int, col int, value interfaces.Expression) TransferenciaReturnExp {
	return TransferenciaReturnExp{Lin: lin, Col: col, Value: value}
}

func (v TransferenciaReturnExp) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	value := v.Value.Ejecutar(ast, gen)

	symbol := environment.Symbol{
		Lin:   v.Lin,
		Col:   v.Col,
		Tipo:  value.Type,
		Valor: value.Value,
		Scope: ast.ObtenerAmbito(),
	}
	Variable := environment.Variable{
		Name:        "ReturnExp",
		Symbols:     symbol,
		Mutable:     false,
		TipoSimbolo: "Sentencia de Transferencia",
	}

	gen.AddComment("Retorno de variable")

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
	return nil
}
