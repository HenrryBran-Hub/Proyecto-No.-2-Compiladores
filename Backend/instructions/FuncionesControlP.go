package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
	"container/list"
	"strconv"
)

type FuncionesControlP struct {
	Lin   int
	Col   int
	Name  string
	Lista interfaces.Instruction
}

func NewFuncionesControlP(lin int, col int, name string, lista interfaces.Instruction) FuncionesControlP {
	return FuncionesControlP{lin, col, name, lista}
}

func (v FuncionesControlP) Ejecutar(ast *environment.AST) interface{} {
	v.Lista.Ejecutar(ast)

	existfun := ast.GetFuncion(v.Name)
	if existfun == nil {
		Errores := environment.Errores{
			Descripcion: "No existe la funcion",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		ast.Lista_Funciones_Par.Init()
		return nil
	}

	if existfun.IsParame == false && ast.Lista_Funciones_Par.Len() > 0 {
		Errores := environment.Errores{
			Descripcion: "La funcion no debe de tener parametros",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		ast.Lista_Funciones_Par.Init()
		return nil
	}

	if existfun.Parametros.Len() != ast.Lista_Funciones_Par.Len() {
		Errores := environment.Errores{
			Descripcion: "La cantidad de parametros ingresados no coincide con la cantidad de parametros de la funcion",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		ast.Lista_Funciones_Par.Init()
		return nil
	}

	var listaparametros = list.New()
	for e := ast.Lista_Funciones_Par.Front(); e != nil; e = e.Next() {
		vari := e.Value.(environment.VariableFuncion)
		listaparametros.PushBack(vari)
	}
	ast.Lista_Funciones_Par.Init()

	e1 := existfun.Parametros.Front()
	e2 := listaparametros.Front()
	listavariablesinterna := list.New()

	for e1 != nil && e2 != nil {
		valor1 := e1.Value.(environment.VariableFuncion)
		valor2 := e2.Value.(environment.VariableFuncion)

		symbol := environment.Symbol{
			Lin:   valor1.Symbols.Lin,
			Col:   valor1.Symbols.Col,
			Tipo:  valor1.Symbols.Tipo,
			Valor: valor2.Symbols.Valor,
			Scope: valor1.Symbols.Scope,
		}
		Variable := environment.Variable{
			Name:        valor1.Name,
			Symbols:     symbol,
			Mutable:     true,
			TipoSimbolo: "Variable",
		}

		if valor1.Symbols.Tipo != valor2.Symbols.Tipo {
			Errores := environment.Errores{
				Descripcion: "El tipo ingresado a la funcion no es el mismo que esta definido en la funcion",
				Fila:        strconv.Itoa(v.Lin),
				Columna:     strconv.Itoa(v.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
			listaparametros.Init()
			return nil
		}

		if valor1.EI == true && valor2.EI == false {
			if valor1.ExternoInterno != "_" {
				Errores := environment.Errores{
					Descripcion: "Tiene que agregar las variables internas osea \"oper: expr\", y solo esta poniendo \"expr\"",
					Fila:        strconv.Itoa(v.Lin),
					Columna:     strconv.Itoa(v.Col),
					Tipo:        "Error Semantico",
					Ambito:      ast.ObtenerAmbito(),
				}
				ast.ErroresHTML(Errores)
				listaparametros.Init()
				return nil
			}
		} else if valor1.EI == true && valor2.EI == true {
			if valor1.ExternoInterno != valor2.Name {
				Errores := environment.Errores{
					Descripcion: "Tiene que agregar las variables internas osea \"oper: expr\" correctamente osea en orden, o poner las que son correctas en caso las puso en orden",
					Fila:        strconv.Itoa(v.Lin),
					Columna:     strconv.Itoa(v.Col),
					Tipo:        "Error Semantico",
					Ambito:      ast.ObtenerAmbito(),
				}
				ast.ErroresHTML(Errores)
				listaparametros.Init()
				return nil
			}
		}

		if valor1.Inout == true && valor2.Inout == false {
			Errores := environment.Errores{
				Descripcion: "No se esta enviando el valor como referencia",
				Fila:        strconv.Itoa(v.Lin),
				Columna:     strconv.Itoa(v.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
			listaparametros.Init()
			return nil
		} else if valor1.Inout == false && valor2.Inout == true {
			Errores := environment.Errores{
				Descripcion: "Se esta enviando el valor como referencia",
				Fila:        strconv.Itoa(v.Lin),
				Columna:     strconv.Itoa(v.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
			listaparametros.Init()
			return nil
		}

		listavariablesinterna.PushBack(Variable)

		e1 = e1.Next()
		e2 = e2.Next()
	}

	ast.AumentarAmbito("Funcion-" + v.Name)
	for e := listavariablesinterna.Front(); e != nil; e = e.Next() {
		vari := e.Value.(environment.Variable)
		ast.GuardarVariable(vari)
	}

	for _, inst := range existfun.CodigoFuncion {
		if inst == nil {
			continue
		}
		instruction, ok := inst.(interfaces.Instruction)
		if !ok {
			continue
		}
		instruction.Ejecutar(ast)
		rvari := ast.GetVariable("Return")
		if rvari != nil {
			break
		}
		revari := ast.GetVariable("ReturnExp")
		if revari != nil {
			Errores := environment.Errores{
				Descripcion: "Se esta devolviendo un valor en una funcion que no debe de tener retorno",
				Fila:        strconv.Itoa(v.Lin),
				Columna:     strconv.Itoa(v.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
			listaparametros.Init()
			ast.DisminuirAmbito()
			return nil
		}
	}

	listavariablesinterna2 := list.New()
	for e := ast.Lista_Variables.Front(); e != nil; e = e.Next() {
		listavariablesinterna2.PushBack(e.Value)
	}
	ast.DisminuirAmbito()

	e1 = existfun.Parametros.Front()
	e2 = listaparametros.Front()
	for e1 != nil && e2 != nil {
		valor1 := e1.Value.(environment.VariableFuncion)
		valor2 := e2.Value.(environment.VariableFuncion)

		for e := listavariablesinterna2.Front(); e != nil; e = e.Next() {
			vari2 := e.Value.(environment.Variable)
			if valor1.Name == vari2.Name && valor2.Inout == true {
				variable := ast.GetVariable(valor2.Name)
				variable.Symbols.Valor = vari2.Symbols.Valor
				ast.ActualizarVariable(variable)
			}
		}

		e1 = e1.Next()
		e2 = e2.Next()
	}

	if existfun.IsReturn == true {
		Errores := environment.Errores{
			Descripcion: "Se esta declarando una funcion con retorno fuera de un ambiente que acepte esto",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		listaparametros.Init()
		return nil
	} else {
		listaparametros.Init()
		return nil
	}
}
