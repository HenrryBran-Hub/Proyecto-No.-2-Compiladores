package datoscompuestos

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"strconv"
)

type ArregloAppend struct {
	Name  string
	Value interfaces.Expression
}

func NewArregloAppend(name string, value interfaces.Expression) ArregloAppend {
	return ArregloAppend{name, value}
}

func (v ArregloAppend) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	arreglo := ast.GetArreglo(v.Name)
	expre := v.Value.Ejecutar(ast, gen)
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	if arreglo == nil {
		Errores := environment.Errores{
			Descripcion: "El arreglo que esta intentando agregar no existe: \n Arreglo: " + v.Name,
			Fila:        strconv.Itoa(expre.Val.Symbols.Lin),
			Columna:     strconv.Itoa(expre.Val.Symbols.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		gen.MainCodeF()
		return nil
	} else {
		if arreglo.Symbols.Tipo != expre.Type {
			Errores := environment.Errores{
				Descripcion: "El valor que esta intentando asignar es de diferente tipo que del arreglo",
				Fila:        strconv.Itoa(expre.Val.Symbols.Lin),
				Columna:     strconv.Itoa(expre.Val.Symbols.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
			gen.MainCodeF()
			return nil
		}
	}

	arreglo.Elements.PushBack(expre)

	//llamar a generar concatstring
	gen.AppendVector()
	//concat
	gen.AddComment("Datos Compuestos Arreglo-Append")
	envSize := strconv.Itoa(ast.PosicionStack)
	tmp1 := gen.NewTemp()
	tmp2 := gen.NewTemp()
	gen.AddExpression(tmp1, "P", envSize, "+")
	gen.AddExpression(tmp1, tmp1, "1", "+")
	gen.AddSetStack("(int)"+tmp1, expre.Value)
	gen.AddExpression("P", "P", envSize, "+")
	gen.AddCall("AppendVector")
	gen.AddGetStack(tmp2, "(int)P")
	gen.AddExpression("P", "P", envSize, "-")

	newTemp := gen.NewTemp()
	gen.AddAssign(newTemp, "H")           //Creamos un nuevo puntero Head para que en este vayan los valores
	gen.AddSetHeap("(int)H", expre.Value) // Agregamos el caracter de escape
	gen.AddExpression("H", "H", "1", "+") // agregamos una nueva expresion igual que arriba de addexpression
	gen.AddBr()
	arreglo.ElementsPt.PushBack(newTemp)

	ast.ActualizarArreglo(v.Name, arreglo)
	gen.AddBr()
	gen.MainCodeF()
	return nil
}
