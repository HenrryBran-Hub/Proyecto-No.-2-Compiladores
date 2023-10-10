package datoscompuestos

import (
	"Backend/environment"
	"Backend/generator"
	"strconv"
)

type ArregloRemoveLast struct {
	Line   int
	Col    int
	Remove string
}

func NewArregloRemoveLast(line, col int, remove string) ArregloRemoveLast {
	return ArregloRemoveLast{line, col, remove}
}

func (v ArregloRemoveLast) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	Remove := ast.GetArreglo(v.Remove)
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
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

	lastElementpt := Remove.ElementsPt.Back()
	if lastElementpt != nil {
		Remove.ElementsPt.Remove(lastElementpt)
	}

	ast.ActualizarArreglo(v.Remove, Remove)
	gen.MainCodeF()
	return nil
}
