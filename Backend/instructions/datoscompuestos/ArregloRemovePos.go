package datoscompuestos

import (
	"Backend/environment"
	"Backend/generator"
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

func (v ArregloRemovePos) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
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

	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	Posicion := v.Pos.Ejecutar(ast, gen)
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}

	if Posicion.Type != environment.INTEGER {
		Errores := environment.Errores{
			Descripcion: "El valor de posicion tiene que ser entero o tipo int" + v.Remove,
			Fila:        strconv.Itoa(v.Line),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		gen.MainCodeF()
		return nil
	}

	if Posicion.IntValue > Remove.Elements.Len() || Posicion.IntValue < 0 {
		if Posicion.IntValue > Remove.Elements.Len() {
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
		gen.MainCodeF()
		return nil
	}

	var elementToDelete *list.Element
	var elemntptToDelete *list.Element
	e := Remove.Elements.Front()
	ee := Remove.ElementsPt.Front()
	for i := 0; e != nil && ee != nil; i++ {
		if i == Posicion.IntValue {
			elementToDelete = e
			elemntptToDelete = ee
			break
		} else {
			e = e.Next()
			ee = ee.Next()
		}
	}

	Remove.Elements.Remove(elementToDelete)
	Remove.ElementsPt.Remove(elemntptToDelete)

	ast.ActualizarArreglo(v.Remove, Remove)
	gen.MainCodeF()
	return nil
}
