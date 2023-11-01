package funciones

import (
	"Backend/environment"
	"Backend/generator"
	"strconv"
)

type FuncionesControlR struct {
	Lin  int
	Col  int
	Name string
}

func NewFuncionesControlR(lin int, col int, name string) FuncionesControlR {
	return FuncionesControlR{lin, col, name}
}

func (v FuncionesControlR) Ejecutar(ast *environment.AST, gen *generator.Generator) environment.Value {
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
		return environment.NewValue("0", false, environment.NULL, false, false, false, environment.Variable{})
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
		return environment.NewValue("0", false, environment.NULL, false, false, false, environment.Variable{})
	}

	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}

	newtmp := gen.NewTemp()
	gen.AddExpression("P", "P", strconv.Itoa(ast.PosicionStack), "+")
	gen.AddCall(v.Name)
	gen.AddGetStack(newtmp, "(int)P")
	gen.AddExpression("P", "P", strconv.Itoa(ast.PosicionStack), "-")
	gen.AddBr()

	gen.MainCodeF()
	ast.Lista_Funciones_Par.Init()
	return environment.NewValue(newtmp, false, existfun.Tipo, false, false, false, environment.Variable{})
}
