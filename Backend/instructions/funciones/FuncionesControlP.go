package funciones

import (
	"Backend/environment"
	"Backend/generator"
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

func (v FuncionesControlP) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	v.Lista.Ejecutar(ast, gen)

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

	if !existfun.IsParame && ast.Lista_Funciones_Par.Len() > 0 {
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
			Lin:         valor1.Symbols.Lin,
			Col:         valor1.Symbols.Col,
			Tipo:        valor1.Symbols.Tipo,
			Valor:       valor2.Symbols.Valor,
			Scope:       valor1.Symbols.Scope,
			ValorInt:    valor2.Symbols.ValorInt,
			ValorFloat:  valor2.Symbols.ValorFloat,
			ValorString: valor2.Symbols.ValorString,
			Posicion:    valor1.Symbols.Posicion,
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

		if valor1.EI && !valor2.EI {
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
		} else if valor1.EI && valor2.EI {
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

		if valor1.Inout && !valor2.Inout {
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
		} else if !valor1.Inout && valor2.Inout {
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

	ambito := ast.ObtenerAmbito()
	ast.AumentarAmbito(ambito)
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}

	gen.AddComment("Parametros declaracion")
	contador := 0
	for e := listavariablesinterna.Front(); e != nil; e = e.Next() {
		vari := e.Value.(environment.Variable)
		gen.AddSetStack(strconv.Itoa(existfun.Inicio+contador), vari.Symbols.Valor.(string))
		gen.AddBr()
		contador++
	}

	gen.AddComment("Llamada de funcion")
	newtem1 := gen.NewTemp()
	gen.AddExpression(newtem1, "P", strconv.Itoa(ast.PosicionStack), "+")
	contador = 0
	for e := listavariablesinterna.Front(); e != nil; e = e.Next() {
		gen.AddExpression(newtem1, newtem1, "1", "+")
		newtem2 := gen.NewTemp()
		gen.AddGetStack(newtem2, strconv.Itoa(existfun.Inicio+contador))
		gen.AddSetStack("(int)"+newtem1, newtem2)
		contador++

	}

	gen.AddExpression("P", "P", strconv.Itoa(ast.PosicionStack), "+")
	gen.AddCall(existfun.Nombre)
	gen.AddExpression("P", "P", strconv.Itoa(ast.PosicionStack), "-")

	listavariablesinterna2 := list.New()
	for e := ast.Lista_Variables.Front(); e != nil; e = e.Next() {
		listavariablesinterna2.PushBack(e.Value)
	}
	gen.MainCodeF()
	ast.DisminuirAmbito()

	if existfun.IsReturn {
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
