package datoscompuestos

import (
	"Backend/environment"
	"Backend/generator"
	"strconv"
)

type ArregloIsEmpty struct {
	Line   int
	Col    int
	VEmpty string
}

func NewArregloIsEmpty(line, col int, vempty string) ArregloIsEmpty {
	return ArregloIsEmpty{line, col, vempty}
}

func (v ArregloIsEmpty) Ejecutar(ast *environment.AST, gen *generator.Generator) environment.Value {
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	VEmpty := ast.GetArreglo(v.VEmpty)
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	if VEmpty == nil {
		Errores := environment.Errores{
			Descripcion: "El arreglo que esta intentando ver si esta vacio no existe: \n Arreglo: " + v.VEmpty,
			Fila:        strconv.Itoa(v.Line),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return environment.NewValue("201314439", false, environment.NULL, false, false, false, environment.Variable{})
	}

	symbol := environment.Symbol{
		Lin:      v.Line,
		Col:      v.Col,
		Tipo:     VEmpty.Symbols.Tipo,
		Scope:    VEmpty.Symbols.Scope,
		TipoDato: environment.VECTOR,
		Posicion: VEmpty.Symbols.Posicion,
	}
	Variable := environment.Variable{
		Name:        VEmpty.Name,
		Symbols:     symbol,
		Mutable:     false,
		TipoSimbolo: "Vector",
	}

	if VEmpty.Elements.Len() == 0 {
		newTemp := gen.NewTemp()
		gen.AddAssign(newTemp, "1")
		gen.AddSetStack(strconv.Itoa(VEmpty.Symbols.Posicion), newTemp)
		gen.MainCodeF()
		return environment.NewValue("1", false, environment.BOOLEAN, false, false, false, Variable)
	} else {
		newTemp := gen.NewTemp()
		gen.AddAssign(newTemp, "0")
		gen.AddSetStack(strconv.Itoa(VEmpty.Symbols.Posicion), newTemp)
		gen.MainCodeF()
		return environment.NewValue("1", false, environment.BOOLEAN, false, false, false, Variable)
	}
}
