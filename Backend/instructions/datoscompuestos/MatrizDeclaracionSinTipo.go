package datoscompuestos

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"container/list"
	"strconv"
)

type MatrizDeclaracionSinTipo struct {
	Lin  int
	Col  int
	Name string
	Def  interfaces.Instruction
}

func NewMatrizDeclaracionSinTipo(lin int, col int, name string, def interfaces.Instruction) MatrizDeclaracionSinTipo {
	return MatrizDeclaracionSinTipo{lin, col, name, def}
}

func (v MatrizDeclaracionSinTipo) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	matexit := ast.GetMatriz(v.Name)
	if matexit != nil {
		Errores := environment.Errores{
			Descripcion: "Ya existe la matriz en este ambito",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		ast.QuitarNiveles()
		return nil
	}

	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	v.Def.Ejecutar(ast, gen)
	if !ast.IsMain(ast.ObtenerAmbito()) {
		gen.MainCodeT()
	}
	tipo := ast.Lista_Matriz_Val.Back().Value.(environment.Valores_Matriz).Valor
	tipo.Scope = ast.ObtenerAmbito()

	Valor := ast.Lista_Matriz_Val.Back()
	Condiciones := Valor.Value.(environment.Valores_Matriz)

	n1 := -1
	n2 := -1
	n3 := -1
	n4 := -1
	n5 := -1
	contador := 0
	if Condiciones.Matriztam != nil {
		for nivel := ast.Pila_Matriz_Val.Front(); nivel != nil; nivel = nivel.Next() {
			lista := nivel.Value.(*list.List)
			for elem := lista.Front(); elem != nil; elem = elem.Next() {
				valores := elem.Value.(environment.Valores_Matriz)
				if valores.Tipo == tipo.Tipo {
					for e := valores.Matriztam.Front(); e != nil; e = e.Next() {
						switch contador {
						case 0:
							n1 = e.Value.(int)
						case 1:
							n2 = e.Value.(int)
						case 2:
							n3 = e.Value.(int)
						case 3:
							n4 = e.Value.(int)
						case 4:
							n5 = e.Value.(int)
						}
						contador++
					}
				} else {
					Errores := environment.Errores{
						Descripcion: "Los tipos de datos no corresponde con el tipo de matriz",
						Fila:        strconv.Itoa(v.Lin),
						Columna:     strconv.Itoa(v.Col),
						Tipo:        "Error Semantico",
						Ambito:      ast.ObtenerAmbito(),
					}
					ast.ErroresHTML(Errores)
					ast.QuitarNiveles()
					return nil
				}
			}
		}

		var matriz environment.Matriz
		if n1 > 0 && n2 > 0 && n3 > 0 && n4 > 0 && n5 > 0 {
			matriz = ast.NuevaMatriz(v.Name, true, tipo, n1, n2, n3, n4, n5)
			var valormatri interface{}
			if matriz.Symbols.Tipo == environment.INTEGER || matriz.Symbols.Tipo == environment.FLOAT {
				newTmp := gen.NewTemp()
				gen.AddAssign(newTmp, "H")
				valormatri1 := ast.Lista_Matriz_Val.Back().Value.(environment.Valores_Matriz).Valor.Valor
				gen.AddSetHeap("(int)H", valormatri1.(string))
				gen.AddExpression("H", "H", "1", "+")
				valormatri = newTmp
			} else {
				valormatri1 := ast.Lista_Matriz_Val.Back().Value.(environment.Valores_Matriz).Valor.Valor
				valormatri = valormatri1
			}
			ast.SustituirValores5(&matriz, valormatri, n1, n2, n3, n4, n5)
		}
		if n1 > 0 && n2 > 0 && n3 > 0 && n4 > 0 && n5 == -1 {
			matriz = ast.NuevaMatriz(v.Name, true, tipo, n1, n2, n3, n4)
			var valormatri interface{}
			if matriz.Symbols.Tipo == environment.INTEGER || matriz.Symbols.Tipo == environment.FLOAT {
				newTmp := gen.NewTemp()
				gen.AddAssign(newTmp, "H")
				valormatri1 := ast.Lista_Matriz_Val.Back().Value.(environment.Valores_Matriz).Valor.Valor
				gen.AddSetHeap("(int)H", valormatri1.(string))
				gen.AddExpression("H", "H", "1", "+")
				valormatri = newTmp
			} else {
				valormatri1 := ast.Lista_Matriz_Val.Back().Value.(environment.Valores_Matriz).Valor.Valor
				valormatri = valormatri1
			}
			ast.SustituirValores4(&matriz, valormatri, n1, n2, n3, n4)
		}
		if n1 > 0 && n2 > 0 && n3 > 0 && n4 == -1 && n5 == -1 {
			matriz = ast.NuevaMatriz(v.Name, true, tipo, n1, n2, n3)
			var valormatri interface{}
			if matriz.Symbols.Tipo == environment.INTEGER || matriz.Symbols.Tipo == environment.FLOAT {
				newTmp := gen.NewTemp()
				gen.AddAssign(newTmp, "H")
				valormatri1 := ast.Lista_Matriz_Val.Back().Value.(environment.Valores_Matriz).Valor.Valor
				gen.AddSetHeap("(int)H", valormatri1.(string))
				gen.AddExpression("H", "H", "1", "+")
				valormatri = newTmp
			} else {
				valormatri1 := ast.Lista_Matriz_Val.Back().Value.(environment.Valores_Matriz).Valor.Valor
				valormatri = valormatri1
			}
			ast.SustituirValores3(&matriz, valormatri, n1, n2, n3)
		}
		if n1 > 0 && n2 > 0 && n3 == -1 && n4 == -1 && n5 == -1 {
			matriz = ast.NuevaMatriz(v.Name, true, tipo, n1, n2)
			var valormatri interface{}
			if matriz.Symbols.Tipo == environment.INTEGER || matriz.Symbols.Tipo == environment.FLOAT {
				newTmp := gen.NewTemp()
				gen.AddAssign(newTmp, "H")
				valormatri1 := ast.Lista_Matriz_Val.Back().Value.(environment.Valores_Matriz).Valor.Valor
				gen.AddSetHeap("(int)H", valormatri1.(string))
				gen.AddExpression("H", "H", "1", "+")
				valormatri = newTmp
			} else {
				valormatri1 := ast.Lista_Matriz_Val.Back().Value.(environment.Valores_Matriz).Valor.Valor
				valormatri = valormatri1
			}
			ast.SustituirValores2(&matriz, valormatri, n1, n2)
		}
		ast.GuardarMatriz(matriz)
		ast.QuitarNiveles()
		gen.MainCodeF()
		return nil
	}

	tipo2 := ast.Lista_Matriz_Val.Front().Value.(environment.Valores_Matriz).Elements.Front().Value.(environment.Symbol)

	if Condiciones.Elements != nil {

		ast.ImprimirArreglovalores()
		contadorpila := 0
		for nivel := ast.Pila_Matriz_Val.Front(); nivel != nil; nivel = nivel.Next() {
			lista := nivel.Value.(*list.List)
			if lista.Back() != nil {
				contadorpila++
			}
		}

		//comprobacion del tipo en todos
		for nivel := ast.Pila_Matriz_Val.Front(); nivel != nil; nivel = nivel.Next() {
			lista := nivel.Value.(*list.List)
			for elem := lista.Front(); elem != nil; elem = elem.Next() {
				valores := elem.Value.(environment.Valores_Matriz)
				if valores.Elements != nil {
					for e := valores.Elements.Front(); e != nil; e = e.Next() {
						if tipo2.Tipo != e.Value.(environment.Symbol).Tipo {
							Errores := environment.Errores{
								Descripcion: "Algun dato dentro de la matriz no es del mismo tipo que se declaro la matriz",
								Fila:        strconv.Itoa(v.Lin),
								Columna:     strconv.Itoa(v.Col),
								Tipo:        "Error Semantico",
								Ambito:      ast.ObtenerAmbito(),
							}
							ast.ErroresHTML(Errores)
							ast.QuitarNiveles()
							return nil
						}
					}
				}
			}
		}
		//comprobacion de elementos de mas
		for nivel := ast.Pila_Matriz_Val.Front(); nivel != nil; nivel = nivel.Next() {
			lista := nivel.Value.(*list.List)
			for elem := lista.Front(); elem != nil; elem = elem.Next() {
				valores := elem.Value.(environment.Valores_Matriz)
				if valores.Elements != nil {
					if elem.Next() != nil {
						if valores.Elements.Len() != elem.Next().Value.(environment.Valores_Matriz).Elements.Len() {
							Errores := environment.Errores{
								Descripcion: "Hay mas datos de los que se pueden asignar a la matriz porfavor verifique",
								Fila:        strconv.Itoa(v.Lin),
								Columna:     strconv.Itoa(v.Col),
								Tipo:        "Error Semantico",
								Ambito:      ast.ObtenerAmbito(),
							}
							ast.ErroresHTML(Errores)
							ast.QuitarNiveles()
							return nil
						}
					}
				}
			}
		}

		var matriz environment.Matriz
		lista := ast.Lista_Matriz_Val.Len()
		elementos := ast.Lista_Matriz_Val.Back().Value.(environment.Valores_Matriz).Elements.Len()
		etiqueta := false

		if contadorpila < 2 {
			matriz = ast.NuevaMatriz(v.Name, true, tipo2, lista, elementos)
			etiqueta = true
		} else {
			matriz = ast.NuevaMatriz(v.Name, true, tipo2, contadorpila, lista, elementos)
		}

		contadorpila--

		for nivel := ast.Pila_Matriz_Val.Front(); nivel != nil; nivel = nivel.Next() {
			lista := nivel.Value.(*list.List)
			contadorlista := 0
			for elem := lista.Front(); elem != nil; elem = elem.Next() {
				valores := elem.Value.(environment.Valores_Matriz)
				if valores.Elements != nil {
					contadorvalor := 0
					for e := valores.Elements.Front(); e != nil; e = e.Next() {
						if etiqueta {
							if matriz.Symbols.Tipo == environment.INTEGER || matriz.Symbols.Tipo == environment.FLOAT {
								var valormatri interface{}
								newTmp := gen.NewTemp()
								gen.AddAssign(newTmp, "H")
								valormatri1 := e.Value.(environment.Value).Value
								gen.AddSetHeap("(int)H", valormatri1)
								gen.AddExpression("H", "H", "1", "+")
								valormatri = newTmp
								ast.IngresarValor(&matriz, []int{contadorlista, contadorvalor}, valormatri)
							} else {
								ast.IngresarValor(&matriz, []int{contadorlista, contadorvalor}, e.Value.(environment.Value).Value)
							}
						} else {
							if matriz.Symbols.Tipo == environment.INTEGER || matriz.Symbols.Tipo == environment.FLOAT {
								var valormatri interface{}
								newTmp := gen.NewTemp()
								gen.AddAssign(newTmp, "H")
								valormatri1 := e.Value.(environment.Value).Value
								gen.AddSetHeap("(int)H", valormatri1)
								gen.AddExpression("H", "H", "1", "+")
								valormatri = newTmp
								ast.IngresarValor(&matriz, []int{contadorpila, contadorlista, contadorvalor}, valormatri)
							} else {
								ast.IngresarValor(&matriz, []int{contadorpila, contadorlista, contadorvalor}, e.Value.(environment.Value).Value)
							}
						}
						contadorvalor++
					}
				}
				contadorlista++
			}
			contadorpila--
		}

		ast.GuardarMatriz(matriz)
		ast.QuitarNiveles()
		gen.MainCodeF()
		return nil
	}

	ast.QuitarNiveles()
	gen.MainCodeF()
	return nil
}
