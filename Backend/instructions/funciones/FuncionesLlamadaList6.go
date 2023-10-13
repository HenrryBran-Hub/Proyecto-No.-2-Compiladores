package funciones

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
)

type FuncionesLlamadaList6 struct {
	Expr interfaces.Expression
}

func NewFuncionesLlamadaList6(expr interfaces.Expression) FuncionesLlamadaList6 {
	return FuncionesLlamadaList6{expr}
}

func (v FuncionesLlamadaList6) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	valor := v.Expr.Ejecutar(ast, gen)
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	symbol := environment.Symbol{
		Lin:         valor.Val.Symbols.Lin,
		Col:         valor.Val.Symbols.Col,
		Tipo:        valor.Val.Symbols.Tipo,
		Valor:       valor.Value,
		ValorInt:    valor.IntValue,
		ValorFloat:  valor.FloatValue,
		ValorString: valor.StringValue,
		Scope:       ast.ObtenerAmbito(),
	}

	variablefuncion := environment.VariableFuncion{
		Symbols:     symbol,
		Mutable:     true,
		TipoSimbolo: "Funcion",
		Inout:       false,
		EI:          false,
	}

	ast.Lista_Funciones_Par.PushBack(variablefuncion)
	gen.MainCodeF()
	return true
}
