package datoscompuestos

type ArregloRemoveLast struct {
	Line   int
	Col    int
	Remove string
}

func NewArregloRemoveLast(line, col int, remove string) ArregloRemoveLast {
	return ArregloRemoveLast{line, col, remove}
}

/*
func (v ArregloRemoveLast) Ejecutar(ast *environment.AST) interface{} {
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

	lastElement := Remove.Elements.Back()
	if lastElement != nil {
		Remove.Elements.Remove(lastElement)
	}

	ast.ActualizarArreglo(v.Remove, Remove)
	return nil
}
*/
