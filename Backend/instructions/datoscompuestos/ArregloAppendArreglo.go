package datoscompuestos

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"strconv"
)

type ArregloAppendArreglo struct {
	Prin string
	Pop  interfaces.Expression
	Secu string
	Sop  interfaces.Expression
}

func NewArregloAppendArreglo(prin string, pop interfaces.Expression, secu string, sop interfaces.Expression) ArregloAppendArreglo {
	return ArregloAppendArreglo{prin, pop, secu, sop}
}

func (v ArregloAppendArreglo) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	Principal := ast.GetArreglo(v.Prin)
	Prinop := v.Pop.Ejecutar(ast, gen)
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	Secundaria := ast.GetArreglo(v.Secu)
	Secuop := v.Sop.Ejecutar(ast, gen)
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	if Principal == nil || Secundaria == nil {
		if Principal == nil {
			Errores := environment.Errores{
				Descripcion: "El arreglo que esta intentando agregar no existe: \n Arreglo: " + v.Prin,
				Fila:        strconv.Itoa(Prinop.Val.Symbols.Lin),
				Columna:     strconv.Itoa(Prinop.Val.Symbols.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
		} else {
			Errores := environment.Errores{
				Descripcion: "El arreglo que esta intentando agregar no existe: \n Arreglo: " + v.Secu,
				Fila:        strconv.Itoa(Secuop.Val.Symbols.Lin),
				Columna:     strconv.Itoa(Secuop.Val.Symbols.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
		}
		gen.MainCodeF()
		return nil
	}

	if Principal.Symbols.Tipo != Secundaria.Symbols.Tipo {
		Errores := environment.Errores{
			Descripcion: "El valor que esta intentando asignar es de diferente tipo que del arreglo",
			Fila:        strconv.Itoa(Prinop.Val.Symbols.Lin),
			Columna:     strconv.Itoa(Prinop.Val.Symbols.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		gen.MainCodeF()
		return nil
	}

	if Prinop.Type != environment.INTEGER || Secuop.Type != environment.INTEGER {
		Errores := environment.Errores{
			Descripcion: "Tiene que ingresar valores enteros para poder ingresar a los vectores",
			Fila:        strconv.Itoa(Prinop.Val.Symbols.Lin),
			Columna:     strconv.Itoa(Prinop.Val.Symbols.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		gen.MainCodeF()
		return nil
	}

	if Prinop.IntValue >= Principal.Elements.Len() || Prinop.IntValue < 0 {
		if Prinop.IntValue >= Principal.Elements.Len() {
			Errores := environment.Errores{
				Descripcion: "El tamaño de poscion es mayor al tamaño del vector Principal",
				Fila:        strconv.Itoa(Prinop.Val.Symbols.Lin),
				Columna:     strconv.Itoa(Prinop.Val.Symbols.Lin),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
		} else {
			Errores := environment.Errores{
				Descripcion: "El tamaño de poscion debe ser mayor o igual a 0",
				Fila:        strconv.Itoa(Prinop.Val.Symbols.Lin),
				Columna:     strconv.Itoa(Prinop.Val.Symbols.Lin),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
		}
		gen.MainCodeF()
		return nil
	}

	if Secuop.IntValue >= Secundaria.Elements.Len() || Secuop.IntValue < 0 {
		if Secuop.IntValue >= Secundaria.Elements.Len() {
			Errores := environment.Errores{
				Descripcion: "El tamaño de poscion es mayor al tamaño del vector Secundario",
				Fila:        strconv.Itoa(Secuop.Val.Symbols.Lin),
				Columna:     strconv.Itoa(Secuop.Val.Symbols.Lin),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
		} else {
			Errores := environment.Errores{
				Descripcion: "El tamaño de poscion debe ser mayor o igual a 0",
				Fila:        strconv.Itoa(Secuop.Val.Symbols.Lin),
				Columna:     strconv.Itoa(Secuop.Val.Symbols.Lin),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
		}
		gen.MainCodeF()
		return nil
	}

	ValPrin := Principal.Elements.Front()
	ValPrinpt := Principal.ElementsPt.Front()
	for i := 0; i < Prinop.IntValue && ValPrin != nil; i++ {
		ValPrin = ValPrin.Next()
		ValPrinpt = ValPrinpt.Next()
	}

	ValSecu := Secundaria.Elements.Front()
	for i := 0; i < Secuop.IntValue && ValSecu != nil; i++ {
		ValSecu = ValSecu.Next()
	}

	gen.AddComment("Datos Compuestos Arreglo-Append-Arreglo")
	if ValPrin != nil && ValSecu != nil {
		newTemp2 := gen.NewTemp()
		cambio := ValPrin.Value.(environment.Value)
		cambio.Value = newTemp2
		gen.AddAssign(newTemp2, ValSecu.Value.(environment.Value).Value)
		newTemp := gen.NewTemp()
		gen.AddAssign(newTemp, "H")           //Creamos un nuevo puntero Head para que en este vayan los valores
		gen.AddSetHeap("(int)H", newTemp2)    // Agregamos el caracter de escape
		gen.AddExpression("H", "H", "1", "+") // agregamos una nueva expresion igual que arriba de addexpression
		gen.AddBr()

		ValPrinpt.Value = newTemp2
		ValPrin.Value = ValSecu.Value
	}

	ast.ActualizarArreglo(v.Prin, Principal)
	gen.AddBr()
	gen.MainCodeF()
	return nil
}
