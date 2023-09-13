package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
	"container/list"
	"strconv"
)

type ArregloRemovePos struct {
	Line   int
	Col    int
	Remove string
	Pos    interfaces.Expression
}

func NewArregloRemovePos(line, col int, remove string, pos interfaces.Expression) ArregloRemovePos {
	return ArregloRemovePos{line, col, remove, pos}
}

func (v ArregloRemovePos) Ejecutar(ast *environment.AST) interface{} {
	Remove := ast.GetArreglo(v.Remove)
	if Remove == nil {
		Errores := environment.Errores{
			Descripcion: "El arreglo que esta intentando quitar el ultimo valor no existe: \n Arreglo: " + v.Remove,
			Fila:        strconv.Itoa(v.Line),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return nil
	}

	if Remove.Elements.Len() == 0 {
		Errores := environment.Errores{
			Descripcion: "El arreglo esta vacio, no se pueden eliminar mas valores \n Arreglo: " + v.Remove,
			Fila:        strconv.Itoa(v.Line),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
	}

	Posicion := v.Pos.Ejecutar(ast)

	if Posicion.Tipo != environment.INTEGER {
		Errores := environment.Errores{
			Descripcion: "El valor de posicion tiene que ser entero o tipo int" + v.Remove,
			Fila:        strconv.Itoa(v.Line),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return nil
	}

	if Posicion.Valor.(int) > Remove.Elements.Len() || Posicion.Valor.(int) < 0 {
		if Posicion.Valor.(int) > Remove.Elements.Len() {
			Errores := environment.Errores{
				Descripcion: "Esta ingresando un valor mayor que el tamaño del vector" + v.Remove,
				Fila:        strconv.Itoa(v.Line),
				Columna:     strconv.Itoa(v.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
		} else {
			Errores := environment.Errores{
				Descripcion: "El valor de posicion tiene que ser mayor o igual a 0" + v.Remove,
				Fila:        strconv.Itoa(v.Line),
				Columna:     strconv.Itoa(v.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
		}
		return nil
	}

	var elementToDelete *list.Element
	for e, i := Remove.Elements.Front(), 0; e != nil; e, i = e.Next(), i+1 {
		if i == Posicion.Valor {
			elementToDelete = e
			break
		}
	}

	Remove.Elements.Remove(elementToDelete)

	ast.ActualizarArreglo(v.Remove, Remove)
	return nil
}
