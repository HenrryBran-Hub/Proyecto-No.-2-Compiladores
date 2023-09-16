package datoscompuestos

import (
	"Backend/interfaces"
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

/*
func (v ArregloAppendArreglo) Ejecutar(ast *environment.AST) interface{} {
	Principal := ast.GetArreglo(v.Prin)
	Prinop := v.Pop.Ejecutar(ast)
	Secundaria := ast.GetArreglo(v.Secu)
	Secuop := v.Sop.Ejecutar(ast)
	if Principal == nil || Secundaria == nil {
		if Principal == nil {
			Errores := environment.Errores{
				Descripcion: "El arreglo que esta intentando agregar no existe: \n Arreglo: " + v.Prin,
				Fila:        strconv.Itoa(Prinop.Lin),
				Columna:     strconv.Itoa(Prinop.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
		} else {
			Errores := environment.Errores{
				Descripcion: "El arreglo que esta intentando agregar no existe: \n Arreglo: " + v.Secu,
				Fila:        strconv.Itoa(Secuop.Lin),
				Columna:     strconv.Itoa(Secuop.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
		}
		return nil
	}

	if Principal.Symbols.Tipo != Secundaria.Symbols.Tipo {
		Errores := environment.Errores{
			Descripcion: "El valor que esta intentando asignar es de diferente tipo que del arreglo",
			Fila:        strconv.Itoa(Prinop.Lin),
			Columna:     strconv.Itoa(Prinop.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return nil
	}

	if Prinop.Tipo != environment.INTEGER || Secuop.Tipo != environment.INTEGER {
		Errores := environment.Errores{
			Descripcion: "Tiene que ingresar valores enteros para poder ingresar a los vectores",
			Fila:        strconv.Itoa(Prinop.Lin),
			Columna:     strconv.Itoa(Prinop.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return nil
	}

	if Prinop.Valor.(int) >= Principal.Elements.Len() || Prinop.Valor.(int) < 0 {
		if Prinop.Valor.(int) >= Principal.Elements.Len() {
			Errores := environment.Errores{
				Descripcion: "El tamaño de poscion es mayor al tamaño del vector Principal",
				Fila:        strconv.Itoa(Prinop.Lin),
				Columna:     strconv.Itoa(Prinop.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
		} else {
			Errores := environment.Errores{
				Descripcion: "El tamaño de poscion debe ser mayor o igual a 0",
				Fila:        strconv.Itoa(Prinop.Lin),
				Columna:     strconv.Itoa(Prinop.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
		}
		return nil
	}

	if Secuop.Valor.(int) >= Secundaria.Elements.Len() || Secuop.Valor.(int) < 0 {
		if Secuop.Valor.(int) >= Secundaria.Elements.Len() {
			Errores := environment.Errores{
				Descripcion: "El tamaño de poscion es mayor al tamaño del vector Secundario",
				Fila:        strconv.Itoa(Prinop.Lin),
				Columna:     strconv.Itoa(Prinop.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
		} else {
			Errores := environment.Errores{
				Descripcion: "El tamaño de poscion debe ser mayor o igual a 0",
				Fila:        strconv.Itoa(Prinop.Lin),
				Columna:     strconv.Itoa(Prinop.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
		}
		return nil
	}

	ValPrin := Principal.Elements.Front()
	for i := 0; i < Prinop.Valor.(int) && ValPrin != nil; i++ {
		ValPrin = ValPrin.Next()
	}

	ValSecu := Secundaria.Elements.Front()
	for i := 0; i < Secuop.Valor.(int) && ValSecu != nil; i++ {
		ValSecu = ValSecu.Next()
	}

	if ValPrin != nil && ValSecu != nil {
		ValPrin.Value = ValSecu.Value
	}

	ast.ActualizarArreglo(v.Prin, Principal)
	return nil
}
*/
