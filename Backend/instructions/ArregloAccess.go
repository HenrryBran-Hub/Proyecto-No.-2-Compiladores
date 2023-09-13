package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
	"container/list"
	"strconv"
)

type ArregloAccess struct {
	Line    int
	Col     int
	VAccess string
	Pos     interfaces.Expression
}

func NewArregloAccess(line, col int, vaccess string, pos interfaces.Expression) ArregloAccess {
	return ArregloAccess{line, col, vaccess, pos}
}

func (v ArregloAccess) Ejecutar(ast *environment.AST) environment.Symbol {
	VAccess := ast.GetArreglo(v.VAccess)
	if VAccess == nil {
		Errores := environment.Errores{
			Descripcion: "El arreglo que esta intentando accedes no existe: \n Arreglo: " + v.VAccess,
			Fila:        strconv.Itoa(v.Line),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return environment.Symbol{Lin: v.Line, Col: v.Col, Tipo: environment.NULL, Valor: nil}
	}

	if VAccess.Elements.Len() == 0 {
		Errores := environment.Errores{
			Descripcion: "El arreglo esta vacio, no se pueden acceder a los datos \n Arreglo: " + v.VAccess,
			Fila:        strconv.Itoa(v.Line),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return environment.Symbol{Lin: v.Line, Col: v.Col, Tipo: environment.NULL, Valor: nil}
	}

	Posicion := v.Pos.Ejecutar(ast)

	if Posicion.Tipo != environment.INTEGER {
		Errores := environment.Errores{
			Descripcion: "El valor de posicion tiene que ser entero o tipo int" + v.VAccess,
			Fila:        strconv.Itoa(v.Line),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return environment.Symbol{Lin: v.Line, Col: v.Col, Tipo: environment.NULL, Valor: nil}
	}

	if Posicion.Valor.(int) > VAccess.Elements.Len() || Posicion.Valor.(int) < 0 {
		if Posicion.Valor.(int) > VAccess.Elements.Len() {
			Errores := environment.Errores{
				Descripcion: "Esta ingresando un valor mayor que el tamaño del vector" + v.VAccess,
				Fila:        strconv.Itoa(v.Line),
				Columna:     strconv.Itoa(v.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
		} else {
			Errores := environment.Errores{
				Descripcion: "El valor de posicion tiene que ser mayor o igual a 0" + v.VAccess,
				Fila:        strconv.Itoa(v.Line),
				Columna:     strconv.Itoa(v.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
		}
		return environment.Symbol{Lin: v.Line, Col: v.Col, Tipo: environment.NULL, Valor: nil}
	}

	var eVAccess *list.Element
	for e, i := VAccess.Elements.Front(), 0; e != nil; e, i = e.Next(), i+1 {
		if i == Posicion.Valor {
			eVAccess = e
			break
		}
	}

	data := eVAccess.Value.(environment.Symbol)

	return environment.Symbol{Lin: v.Line, Col: v.Col, Tipo: data.Tipo, Valor: data.Valor}
}
