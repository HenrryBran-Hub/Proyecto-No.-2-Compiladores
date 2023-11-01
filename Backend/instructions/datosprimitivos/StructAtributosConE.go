package datosprimitivos

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
)

type StructAtributosConE struct {
	Line int
	Col  int
	Tipo string
	Name string
	Expr interfaces.Expression
}

func NewStructAtributosConE(line, col int, tipo, name string, expr interfaces.Expression) StructAtributosConE {
	return StructAtributosConE{line, col, tipo, name, expr}
}

func (v StructAtributosConE) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
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
		Scope:       valor.Val.Symbols.Scope,
		ValorInt:    valor.IntValue,
		ValorFloat:  valor.FloatValue,
		ValorString: valor.StringValue,
		Posicion:    valor.Val.Symbols.Posicion,
	}

	Variable := environment.Variable{
		Name:        v.Name,
		Symbols:     symbol,
		Mutable:     true,
		TipoSimbolo: "Variable",
	}

	if v.Tipo == "let" {
		Variable.Mutable = false
	}

	ast.AtributosStruct.PushBack(Variable)
	return nil
}
