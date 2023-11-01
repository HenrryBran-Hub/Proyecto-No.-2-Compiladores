package funciones

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
)

type FuncionesLlamadaList5 struct {
	Lin  int
	Col  int
	Name string
	Expr interfaces.Expression
}

func NewFuncionesLlamadaList5(lin int, col int, name string, expr interfaces.Expression) FuncionesLlamadaList5 {
	return FuncionesLlamadaList5{lin, col, name, expr}
}

func (v FuncionesLlamadaList5) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	valor := v.Expr.Ejecutar(ast, gen)
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	symbol := environment.Symbol{
		Lin:         v.Lin,
		Col:         v.Col,
		Tipo:        valor.Type,
		Valor:       valor.Value,
		ValorInt:    valor.IntValue,
		ValorFloat:  valor.FloatValue,
		ValorString: valor.StringValue,
		Scope:       ast.ObtenerAmbito(),
	}

	variablefuncion := environment.VariableFuncion{
		Name:        v.Name,
		Symbols:     symbol,
		Mutable:     true,
		TipoSimbolo: "Funcion",
		Inout:       false,
		EI:          true,
	}

	ast.Lista_Funciones_Par.PushBack(variablefuncion)
	gen.MainCodeF()
	return true
}
