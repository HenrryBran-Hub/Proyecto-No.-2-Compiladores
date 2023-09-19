package sentencias

import (
	"Backend/environment"
	"Backend/generator"
	"fmt"
	"strconv"
)

type Callid struct {
	Lin  int
	Col  int
	Name string
}

func NewCallid(lin int, col int, name string) Callid {
	exp := Callid{Lin: lin, Col: col, Name: name}
	return exp
}

func (o Callid) Ejecutar(ast *environment.AST, gen *generator.Generator) environment.Value {
	var result environment.Value
	variable := ast.GetVariable(o.Name)
	gen.AddComment("LLamamos la funcion Callid ")
	if variable != nil {
		if ast.ObtenerAmbito() == "Global" {
			gen.MainCodeT()
		}
		result = environment.NewValue(fmt.Sprintf("%v", variable.Symbols.Valor), true, variable.Symbols.Tipo, false, false, false)
		return result
	} else {
		Errores := environment.Errores{
			Descripcion: "La variale que esta intentando llamar no existe: \n Variable: " + o.Name,
			Fila:        strconv.Itoa(o.Lin),
			Columna:     strconv.Itoa(o.Col),
			Tipo:        "Error Semantico",
		}
		ast.ErroresHTML(Errores)
		result = environment.NewValue(fmt.Sprintf("%v", "99999"), true, environment.NULL, false, false, false)
		return result
	}
}
