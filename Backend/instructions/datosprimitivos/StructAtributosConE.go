package datosprimitivos

import (
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

/*
func (v StructAtributosConE) Ejecutar(ast *environment.AST) interface{} {
	symbol := v.Expr.Ejecutar(ast)
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
*/
