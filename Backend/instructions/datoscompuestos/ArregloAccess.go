package datoscompuestos

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"strconv"
)

type ArregloAccess struct {
	Line    int
	Col     int
	VAccess string
	Pos     interfaces.Expression
}

func NewArregloAccess(line, col int, vaccess string, pos interfaces.Expression) ArregloAccess {
	return ArregloAccess{line, col, vaccess, pos}
}

func (v ArregloAccess) Ejecutar(ast *environment.AST, gen *generator.Generator) environment.Value {
	VAccess := ast.GetArreglo(v.VAccess)
	if VAccess == nil {
		Errores := environment.Errores{
			Descripcion: "El arreglo que esta intentando accedes no existe: \n Arreglo: " + v.VAccess,
			Fila:        strconv.Itoa(v.Line),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return environment.NewValue("0", false, environment.NULL, false, false, false, environment.Variable{})
	}

	if VAccess.Elements.Len() == 0 {
		Errores := environment.Errores{
			Descripcion: "El arreglo esta vacio, no se pueden acceder a los datos \n Arreglo: " + v.VAccess,
			Fila:        strconv.Itoa(v.Line),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return environment.NewValue("0", false, environment.NULL, false, false, false, environment.Variable{})
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
			Descripcion: "El valor de posicion tiene que ser entero o tipo int" + v.VAccess,
			Fila:        strconv.Itoa(v.Line),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return environment.NewValue("0", false, environment.NULL, false, false, false, environment.Variable{})
	}

	if Posicion.IntValue > VAccess.Elements.Len() || Posicion.IntValue < 0 {
		if Posicion.IntValue > VAccess.Elements.Len() {
			Errores := environment.Errores{
				Descripcion: "Esta ingresando un valor mayor que el tamaño del vector" + v.VAccess,
				Fila:        strconv.Itoa(v.Line),
				Columna:     strconv.Itoa(v.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
		} else {
			Errores := environment.Errores{
				Descripcion: "El valor de posicion tiene que ser mayor o igual a 0" + v.VAccess,
				Fila:        strconv.Itoa(v.Line),
				Columna:     strconv.Itoa(v.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
		}
		return environment.NewValue("0", false, environment.NULL, false, false, false, environment.Variable{})
	}

	e := VAccess.Elements.Front()
	ee := VAccess.ElementsPt.Front()
	for i := 0; e != nil && ee != nil; i++ {
		if i == Posicion.IntValue {
			break
		} else {
			e = e.Next()
			ee = ee.Next()
		}
	}

	symbol := environment.Symbol{
		Lin:      v.Line,
		Col:      v.Col,
		Tipo:     VAccess.Symbols.Tipo,
		Scope:    VAccess.Symbols.Scope,
		TipoDato: environment.VECTOR,
		Posicion: VAccess.Symbols.Posicion,
	}
	Variable := environment.Variable{
		Name:        VAccess.Name,
		Symbols:     symbol,
		Mutable:     false,
		TipoSimbolo: "Vector",
	}

	gen.AddComment("Datos Compuestos Arreglo-Acceso")
	newTemp := gen.NewTemp()
	gen.AddGetHeap(newTemp, "(int)"+ee.Value.(string))
	gen.AddSetStack(strconv.Itoa(VAccess.Symbols.Posicion), newTemp)
	gen.AddBr()
	result := environment.NewValue(newTemp, false, VAccess.Symbols.Tipo, false, false, false, Variable)
	result.IntValue = VAccess.Elements.Len()
	gen.AddBr()
	gen.MainCodeF()
	return result
}
