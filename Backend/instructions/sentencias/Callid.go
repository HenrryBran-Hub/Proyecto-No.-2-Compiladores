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
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	var result environment.Value
	variable := ast.GetVariable(o.Name)
	gen.AddComment("Estoy dentro de la sentencia Call-Id")
	if variable != nil {
		result = environment.NewValue(fmt.Sprintf("%v", variable.Symbols.Valor), true, variable.Symbols.Tipo, false, false, false, *variable)
		result.IntValue = variable.Symbols.ValorInt
		result.FloatValue = variable.Symbols.ValorFloat
		result.StringValue = variable.Symbols.ValorString
		return result
	} else {
		Errores := environment.Errores{
			Descripcion: "La variale que esta intentando llamar no existe: \n Variable: " + o.Name,
			Fila:        strconv.Itoa(o.Lin),
			Columna:     strconv.Itoa(o.Col),
			Tipo:        "Error Semantico",
		}
		ast.ErroresHTML(Errores)
		result = environment.NewValue(fmt.Sprintf("%v", "99999"), true, environment.NULL, false, false, false, *variable)
		return result
	}
}
