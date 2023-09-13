package instructions

import (
	"Backend/environment"
)

type StructAtributosConT struct {
	Line  int
	Col   int
	Tipo  string
	Name  string
	Type  environment.TipoExpresion
	Type2 string
}

func NewStructAtributosConT(line, col int, tipo string, name string, typet environment.TipoExpresion) StructAtributosConT {
	return StructAtributosConT{line, col, tipo, name, typet, ""}
}

func NewStructAtributosConT2(line, col int, tipo string, name string, typet string) StructAtributosConT {
	return StructAtributosConT{line, col, tipo, name, environment.STRUCT, typet}
}

func (v StructAtributosConT) Ejecutar(ast *environment.AST) interface{} {

	// var tipoexpstr environment.TipoExpresion
	// switch v.Type {
	// case "Int":
	// 	tipoexpstr = environment.INTEGER
	// case "Float":
	// 	tipoexpstr = environment.FLOAT
	// case "String":
	// 	tipoexpstr = environment.STRING
	// case "Bool":
	// 	tipoexpstr = environment.BOOLEAN
	// case "Character":
	// 	tipoexpstr = environment.CHARACTER
	// default:
	// 	tipoexpstr = environment.NULL
	// }

	symbol := environment.Symbol{
		Lin:    v.Line,
		Col:    v.Col,
		Tipo:   v.Type,
		Scope:  ast.ObtenerAmbito(),
		Struct: v.Tipo,
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
