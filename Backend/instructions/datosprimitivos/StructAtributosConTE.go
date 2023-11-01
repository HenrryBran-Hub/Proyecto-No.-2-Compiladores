package datosprimitivos

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"strconv"
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

func (v StructAtributosConTE) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {

	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	valor := v.Expr.Ejecutar(ast, gen)
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	symbol := environment.Symbol{
		Lin:         v.Line,
		Col:         v.Col,
		Tipo:        v.Type,
		Scope:       ast.ObtenerAmbito(),
		TipoDato:    environment.VARIABLE,
		Posicion:    ast.PosicionStack,
		Valor:       valor.Value,
		ValorInt:    valor.IntValue,
		ValorFloat:  valor.FloatValue,
		ValorString: valor.StringValue,
	}

	if valor.Val.Symbols.Tipo != v.Type {
		Errores := environment.Errores{
			Descripcion: "Se ha querido asignar un valor no correspondiente a el tipo de dato: \nTipo de dato:" + v.Tipo,
			Fila:        strconv.Itoa(v.Line),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		symbol.Tipo = valor.Val.Symbols.Tipo
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
