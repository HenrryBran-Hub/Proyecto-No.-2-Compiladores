package datoscompuestos

import (
	"Backend/environment"
	"Backend/generator"
	"strconv"
)

type ArregloCount struct {
	Line   int
	Col    int
	VCount string
}

func NewArregloCount(line, col int, vcount string) ArregloCount {
	return ArregloCount{line, col, vcount}
}

func (v ArregloCount) Ejecutar(ast *environment.AST, gen *generator.Generator) environment.Value {
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	VCount := ast.GetArreglo(v.VCount)
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	if VCount == nil {
		Errores := environment.Errores{
			Descripcion: "El arreglo que esta intentando ver si esta vacio no existe: \n Arreglo: " + v.VCount,
			Fila:        strconv.Itoa(v.Line),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return environment.NewValue("0", false, environment.NULL, false, false, false, environment.Variable{})
	}

	symbol := environment.Symbol{
		Lin:      v.Line,
		Col:      v.Col,
		Tipo:     VCount.Symbols.Tipo,
		Scope:    VCount.Symbols.Scope,
		TipoDato: environment.VECTOR,
		Posicion: VCount.Symbols.Posicion,
	}
	Variable := environment.Variable{
		Name:        VCount.Name,
		Symbols:     symbol,
		Mutable:     false,
		TipoSimbolo: "Vector",
	}

	gen.AddComment("Conteo de vector")
	newTemp := gen.NewTemp()
	gen.AddAssign(newTemp, strconv.Itoa(VCount.Elements.Len()))
	gen.AddSetStack(strconv.Itoa(VCount.Symbols.Posicion), newTemp)
	gen.AddBr()
	result := environment.NewValue(newTemp, false, environment.INTEGER, false, false, false, Variable)
	result.IntValue = VCount.Elements.Len()
	gen.MainCodeF()
	return result
}
