package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
	"strconv"
)

type ArregloAppendExp struct {
	Prin string
	Pop  interfaces.Expression
	Sop  interfaces.Expression
}

func NewArregloAppendExp(prin string, pop interfaces.Expression, sop interfaces.Expression) ArregloAppendExp {
	return ArregloAppendExp{prin, pop, sop}
}

func (v ArregloAppendExp) Ejecutar(ast *environment.AST) interface{} {
	Principal := ast.GetArreglo(v.Prin)
	Prinop := v.Pop.Ejecutar(ast)
	Secuop := v.Sop.Ejecutar(ast)
	if Principal == nil {
		Errores := environment.Errores{
			Descripcion: "El arreglo que esta intentando agregar no existe: \n Arreglo: " + v.Prin,
			Fila:        strconv.Itoa(Prinop.Lin),
			Columna:     strconv.Itoa(Prinop.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return nil
	}

	if Principal.Symbols.Tipo != Secuop.Tipo {
		if Secuop.Tipo == environment.NULL {
			Errores := environment.Errores{
				Descripcion: "Se ingreso un valor nulo a el arreglo",
				Fila:        strconv.Itoa(Prinop.Lin),
				Columna:     strconv.Itoa(Prinop.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
		} else {
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
	}

	if Prinop.Tipo != environment.INTEGER {
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

	ValPrin := Principal.Elements.Front()
	for i := 0; i < Prinop.Valor.(int) && ValPrin != nil; i++ {
		ValPrin = ValPrin.Next()
	}

	if ValPrin != nil {
		ValPrin.Value = Secuop
	}

	ast.ActualizarArreglo(v.Prin, Principal)
	return nil
}
