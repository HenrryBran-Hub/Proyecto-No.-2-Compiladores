package funciones

import (
	"Backend/environment"
	"Backend/generator"
	"strconv"
)

type FuncionesControl struct {
	Lin  int
	Col  int
	Name string
}

func NewFuncionesControl(lin int, col int, name string) FuncionesControl {
	return FuncionesControl{lin, col, name}
}

func (v FuncionesControl) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	existfun := ast.GetFuncion(v.Name)
	if existfun == nil {
		Errores := environment.Errores{
			Descripcion: "No existe la funcion",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		ast.Lista_Funciones_Par.Init()
		return nil
	}

	if existfun.IsParame && ast.Lista_Funciones_Par.Len() > 0 {
		Errores := environment.Errores{
			Descripcion: "La funcion no debe de tener parametros",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		ast.Lista_Funciones_Par.Init()
		return nil
	}

	if existfun.IsReturn {
		Errores := environment.Errores{
			Descripcion: "Se esta declarando una funcion con retorno fuera de un ambiente que acepte esto",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		ast.Lista_Funciones_Par.Init()
		return nil
	}

	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}

	gen.AddCall(v.Name)

	gen.MainCodeF()
	ast.Lista_Funciones_Par.Init()
	return nil
}
