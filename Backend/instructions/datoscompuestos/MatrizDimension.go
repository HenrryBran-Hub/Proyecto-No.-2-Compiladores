package datoscompuestos

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
)

type MatrizDimension struct {
	Lin  int
	Col  int
	Type interfaces.Expression
}

func NewMatrizDimension(lin, col int, tipo interfaces.Expression) MatrizDimension {
	exp := MatrizDimension{Lin: lin, Col: col, Type: tipo}
	return exp
}

func (o MatrizDimension) Ejecutar(ast *environment.AST, gen *generator.Generator) environment.Value {
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	valor := o.Type.Ejecutar(ast, gen)
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	result := valor.Val.Symbols.Valor.(int) + 1
	symbol := environment.Symbol{
		Lin:      o.Lin,
		Col:      o.Col,
		Tipo:     valor.Type,
		Valor:    result,
		Scope:    ast.ObtenerAmbito(),
		Posicion: ast.PosicionStack,
	}
	Variable := environment.Variable{
		Name:        "matriztipo",
		Symbols:     symbol,
		Mutable:     true,
		TipoSimbolo: "Variable",
	}

	gen.MainCodeF()
	return environment.NewValue("", true, valor.Type, false, false, false, Variable)
}
