package funciones

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"container/list"
	"strconv"
)

type FuncionesControlPR struct {
	Lin   int
	Col   int
	Name  string
	Lista interfaces.Instruction
}

func NewFuncionesControlPR(lin int, col int, name string, lista interfaces.Instruction) FuncionesControlPR {
	return FuncionesControlPR{lin, col, name, lista}
}

func (v FuncionesControlPR) Ejecutar(ast *environment.AST, gen *generator.Generator) environment.Value {
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
		return environment.NewValue("0", false, environment.NULL, false, false, false, environment.Variable{})
	}

	var listaparametros = list.New()
	for e := ast.Lista_Funciones_Par.Front(); e != nil; e = e.Next() {
		vari := e.Value.(environment.VariableFuncion)
		listaparametros.PushBack(vari)
	}
	ast.Lista_Funciones_Par.Init()

	if !existfun.IsParame && listaparametros.Len() > 0 {
		Errores := environment.Errores{
			Descripcion: "La funcion no debe de tener parametros",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		listaparametros.Init()
		ast.Lista_Funciones_Par.Init()
		return environment.NewValue("0", false, environment.NULL, false, false, false, environment.Variable{})
	}

	if existfun.Parametros.Len() != listaparametros.Len() {
		Errores := environment.Errores{
			Descripcion: "La cantidad de parametros ingresados no coincide con la cantidad de parametros de la funcion",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		listaparametros.Init()
		ast.Lista_Funciones_Par.Init()
		return environment.NewValue("0", false, environment.NULL, false, false, false, environment.Variable{})
	}

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
			ast.Lista_Funciones_Par.Init()
			return environment.NewValue("0", false, environment.NULL, false, false, false, environment.Variable{})
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
				ast.Lista_Funciones_Par.Init()
				return environment.NewValue("0", false, environment.NULL, false, false, false, environment.Variable{})
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
				ast.Lista_Funciones_Par.Init()
				return environment.NewValue("0", false, environment.NULL, false, false, false, environment.Variable{})
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
			ast.Lista_Funciones_Par.Init()
			return environment.NewValue("0", false, environment.NULL, false, false, false, environment.Variable{})
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
			ast.Lista_Funciones_Par.Init()
			return environment.NewValue("0", false, environment.NULL, false, false, false, environment.Variable{})
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

	newtmp := gen.NewTemp()
	gen.AddExpression("P", "P", strconv.Itoa(ast.PosicionStack+1), "+")
	gen.AddCall(v.Name)
	gen.AddExpression("P", "P", strconv.Itoa(ast.PosicionStack+1), "-")
	gen.AddGetStack(newtmp, "(int)P")

	gen.AddBr()

	gen.MainCodeF()
	ast.DisminuirAmbito()
	gen.MainCodeF()
	ast.Lista_Funciones_Par.Init()
	listaparametros.Init()
	return environment.NewValue(newtmp, false, existfun.Tipo, false, false, false, environment.Variable{})

}
