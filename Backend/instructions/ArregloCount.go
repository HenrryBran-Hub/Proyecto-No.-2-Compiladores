package instructions

import (
	"Backend/environment"
	"strconv"
)

type ArregloCount struct {
	Line   int
	Col    int
	VCount string
}

func NewArregloCount(line, col int, vcount string) ArregloCount {
	return ArregloCount{line, col, vcount}
}

func (v ArregloCount) Ejecutar(ast *environment.AST) environment.Symbol {
	VCount := ast.GetArreglo(v.VCount)
	if VCount == nil {
		Errores := environment.Errores{
			Descripcion: "El arreglo que esta intentando ver si esta vacio no existe: \n Arreglo: " + v.VCount,
			Fila:        strconv.Itoa(v.Line),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return environment.Symbol{Lin: v.Line, Col: v.Col, Tipo: environment.NULL, Valor: nil}
	}

	return environment.Symbol{Lin: v.Line, Col: v.Col, Tipo: environment.INTEGER, Valor: VCount.Elements.Len()}
}
