package funciones

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
)

type FuncionesLlamadaList4 struct {
	Lin   int
	Col   int
	Expr  interfaces.Expression
	Lista interfaces.Instruction
}

func NewFuncionesLlamadaList4(lin int, col int, expr interfaces.Expression, lista interfaces.Instruction) FuncionesLlamadaList4 {
	return FuncionesLlamadaList4{lin, col, expr, lista}
}

func (v FuncionesLlamadaList4) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
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
		Symbols:     symbol,
		Mutable:     true,
		TipoSimbolo: "Funcion",
		Inout:       false,
		EI:          false,
	}

	ast.Lista_Funciones_Par.PushBack(variablefuncion)
	gen.MainCodeF()
	v.Lista.Ejecutar(ast, gen)
	return true
}
