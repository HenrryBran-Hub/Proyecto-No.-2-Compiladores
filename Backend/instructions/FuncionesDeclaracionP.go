package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
	"container/list"
	"strconv"
)

type FuncionesDeclaracionP struct {
	Lin        int
	Col        int
	Name       string
	Parametros interfaces.Instruction
	Bloque     []interface{}
}

func NewFuncionesDeclaracionP(lin int, col int, name string, parametros interfaces.Instruction, bloque []interface{}) FuncionesDeclaracionP {
	return FuncionesDeclaracionP{lin, col, name, parametros, bloque}
}

func (v FuncionesDeclaracionP) Ejecutar(ast *environment.AST) interface{} {
	listavalores := list.New()
	v.Parametros.Ejecutar(ast)
	for e := ast.Lista_Funciones_Var.Front(); e != nil; e = e.Next() {
		valor := e.Value.(environment.VariableFuncion)
		valor.Symbols.Scope = "Funcion-" + v.Name
		listavalores.PushBack(valor)
	}

	for e := ast.Lista_Funciones_Var.Front(); e != nil; e = e.Next() {
		valore := e.Value.(environment.VariableFuncion)

		// Compara con los elementos anteriores en la lista
		for i := ast.Lista_Funciones_Var.Front(); i != e; i = i.Next() {
			valori := i.Value.(environment.VariableFuncion)
			if valore.EI == true && valore.ExternoInterno != "_" {
				if valore.ExternoInterno == valori.ExternoInterno {
					Errores := environment.Errores{
						Descripcion: "Esta ingresando variables externas/internas iguales debe de cambiarles los nombres",
						Fila:        strconv.Itoa(v.Lin),
						Columna:     strconv.Itoa(v.Col),
						Tipo:        "Error Semántico",
						Ambito:      "Funcion-" + v.Name,
					}
					ast.ErroresHTML(Errores)
					ast.Lista_Funciones_Var.Init()
					return nil
				}
			}
		}

		// Compara con los elementos siguientes en la lista
		for i := e.Next(); i != nil; i = i.Next() {
			valori := i.Value.(environment.VariableFuncion)
			if valore.EI == true && valore.ExternoInterno != "_" {
				if valore.ExternoInterno == valori.ExternoInterno {
					Errores := environment.Errores{
						Descripcion: "Esta ingresando variables externas/internas iguales debe de cambiarles los nombres",
						Fila:        strconv.Itoa(v.Lin),
						Columna:     strconv.Itoa(v.Col),
						Tipo:        "Error Semántico",
						Ambito:      "Funcion-" + v.Name,
					}
					ast.ErroresHTML(Errores)
					ast.Lista_Funciones_Var.Init()
					return nil
				}
			}
		}
	}

	for e := ast.Lista_Funciones_Var.Front(); e != nil; e = e.Next() {
		valore := e.Value.(environment.VariableFuncion)

		// Compara con los elementos anteriores en la lista
		for i := ast.Lista_Funciones_Var.Front(); i != e; i = i.Next() {
			valori := i.Value.(environment.VariableFuncion)
			if valore.Name == valori.Name {
				Errores := environment.Errores{
					Descripcion: "Estas identificando los parametros con los mismos nombres tienes que cambiarlos",
					Fila:        strconv.Itoa(v.Lin),
					Columna:     strconv.Itoa(v.Col),
					Tipo:        "Error Semántico",
					Ambito:      "Funcion-" + v.Name,
				}
				ast.ErroresHTML(Errores)
				ast.Lista_Funciones_Var.Init()
				return nil
			}
		}

		// Compara con los elementos siguientes en la lista
		for i := e.Next(); i != nil; i = i.Next() {
			valori := i.Value.(environment.VariableFuncion)
			if valore.Name == valori.Name {
				Errores := environment.Errores{
					Descripcion: "Estas identificando los parametros con los mismos nombres tienes que cambiarlos",
					Fila:        strconv.Itoa(v.Lin),
					Columna:     strconv.Itoa(v.Col),
					Tipo:        "Error Semántico",
					Ambito:      "Funcion-" + v.Name,
				}
				ast.ErroresHTML(Errores)
				ast.Lista_Funciones_Var.Init()
				return nil
			}
		}
	}

	funcion := environment.Funcion{
		Lin:           v.Lin,
		Col:           v.Col,
		Nombre:        v.Name,
		IsReturn:      false,
		IsParame:      true,
		Parametros:    listavalores,
		CodigoFuncion: v.Bloque,
	}

	ast.GuardarFuncion(funcion)
	ast.Lista_Funciones_Var.Init()
	return nil
}
