package instructions

import (
	"Backend/environment"
)

type StructAtributos struct {
	Line int
	Col  int
	Tipo string
	Name string
}

func NewStructAtributos(line, col int, tipo, name string) StructAtributos {
	return StructAtributos{line, col, tipo, name}
}

func (v StructAtributos) Ejecutar(ast *environment.AST) interface{} {

	symbol := environment.Symbol{
		Lin:   v.Line,
		Col:   v.Col,
		Tipo:  environment.NULL,
		Valor: nil,
		Scope: ast.ObtenerAmbito(),
	}

	Variable := environment.Variable{
		Name:        v.Name,
		Symbols:     symbol,
		Mutable:     true,
		TipoSimbolo: "Struct",
	}

	if v.Tipo == "let" {
		Variable.Mutable = false
	}

	ast.AtributosStruct.PushBack(Variable)
	return nil
}
