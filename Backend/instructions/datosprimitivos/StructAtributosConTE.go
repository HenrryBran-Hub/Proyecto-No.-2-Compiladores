package datosprimitivos

import (
	"Backend/environment"
	"Backend/interfaces"
)

type StructAtributosConTE struct {
	Line  int
	Col   int
	Tipo  string
	Name  string
	Type  environment.TipoExpresion
	Expr  interfaces.Expression
	Type2 string
}

func NewStructAtributosConTE(line, col int, tipo string, name string, typet environment.TipoExpresion, expr interfaces.Expression) StructAtributosConTE {
	return StructAtributosConTE{line, col, tipo, name, typet, expr, ""}
}

func NewStructAtributosConTE2(line, col int, tipo string, name string, typet string, expr interfaces.Expression) StructAtributosConTE {
	return StructAtributosConTE{line, col, tipo, name, environment.STRUCT, expr, typet}
}

/*
func (v StructAtributosConTE) Ejecutar(ast *environment.AST) interface{} {

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

	valor := v.Expr.Ejecutar(ast)
	symbol := environment.Symbol{
		Lin:   v.Line,
		Col:   v.Col,
		Tipo:  v.Type,
		Valor: valor.Valor,
		Scope: ast.ObtenerAmbito(),
	}

	if valor.Tipo != v.Type {
		Errores := environment.Errores{
			Descripcion: "Se ha querido asignar un valor no correspondiente a el tipo de dato: \nTipo de dato:" + v.Tipo,
			Fila:        strconv.Itoa(v.Line),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		symbol.Tipo = valor.Tipo
		symbol.Valor = nil
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
*/
