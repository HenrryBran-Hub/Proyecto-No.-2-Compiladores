package funciones

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
)

type FuncionesLlamadaList3 struct {
	Lin   int
	Col   int
	Name  string
	Expr  interfaces.Expression
	Lista interfaces.Instruction
}

func NewFuncionesLlamadaList3(lin int, col int, name string, expr interfaces.Expression, lista interfaces.Instruction) FuncionesLlamadaList3 {
	return FuncionesLlamadaList3{lin, col, name, expr, lista}
}

func (v FuncionesLlamadaList3) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {

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
		Tipo:        valor.Val.Symbols.Tipo,
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
	v.Lista.Ejecutar(ast, gen)
	return true
}
