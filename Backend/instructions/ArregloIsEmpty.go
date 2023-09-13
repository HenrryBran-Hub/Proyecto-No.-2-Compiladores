package instructions

import (
	"Backend/environment"
	"strconv"
)

type ArregloIsEmpty struct {
	Line   int
	Col    int
	VEmpty string
}

func NewArregloIsEmpty(line, col int, vempty string) ArregloIsEmpty {
	return ArregloIsEmpty{line, col, vempty}
}

func (v ArregloIsEmpty) Ejecutar(ast *environment.AST) environment.Symbol {
	VEmpty := ast.GetArreglo(v.VEmpty)
	if VEmpty == nil {
		Errores := environment.Errores{
			Descripcion: "El arreglo que esta intentando ver si esta vacio no existe: \n Arreglo: " + v.VEmpty,
			Fila:        strconv.Itoa(v.Line),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return environment.Symbol{Lin: v.Line, Col: v.Col, Tipo: environment.NULL, Valor: nil}
	}

	if VEmpty.Elements.Len() == 0 {
		return environment.Symbol{Lin: v.Line, Col: v.Col, Tipo: environment.BOOLEAN, Valor: true}
	} else {
		return environment.Symbol{Lin: v.Line, Col: v.Col, Tipo: environment.BOOLEAN, Valor: false}
	}
}
