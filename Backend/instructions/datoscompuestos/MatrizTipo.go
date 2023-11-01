package datoscompuestos

import (
	"Backend/environment"
	"Backend/generator"
)

type MatrizTipo struct {
	Lin int
	Col int
	Op  environment.TipoExpresion
}

func NewMatrizTipo(lin, col int, op environment.TipoExpresion) MatrizTipo {
	exp := MatrizTipo{Lin: lin, Col: col, Op: op}
	return exp
}

func (o MatrizTipo) Ejecutar(ast *environment.AST, gen *generator.Generator) environment.Value {
	symbol := environment.Symbol{
		Lin:      o.Lin,
		Col:      o.Col,
		Tipo:     o.Op,
		Valor:    1,
		Scope:    ast.ObtenerAmbito(),
		Posicion: ast.PosicionStack,
	}
	Variable := environment.Variable{
		Name:        "matriztipo",
		Symbols:     symbol,
		Mutable:     true,
		TipoSimbolo: "Variable",
	}

	return environment.NewValue("", true, o.Op, false, false, false, Variable)
}
