package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
	"container/list"
	"strconv"
)

type MatrizDeclaracion struct {
	Lin  int
	Col  int
	Name string
	Type interfaces.Expression
	Def  interfaces.Instruction
}

func NewMatrizDeclaracion(lin int, col int, name string, tipo interfaces.Expression, def interfaces.Instruction) MatrizDeclaracion {
	return MatrizDeclaracion{lin, col, name, tipo, def}
}

func (v MatrizDeclaracion) Ejecutar(ast *environment.AST) interface{} {
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

	tipo := v.Type.Ejecutar(ast)
	v.Def.Ejecutar(ast)
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
			valormatri := ast.Lista_Matriz_Val.Back().Value.(environment.Valores_Matriz).Valor.Valor
			ast.SustituirValores5(&matriz, valormatri, n1, n2, n3, n4, n5)
		}
		if n1 > 0 && n2 > 0 && n3 > 0 && n4 > 0 && n5 == -1 {
			matriz = ast.NuevaMatriz(v.Name, true, tipo, n1, n2, n3, n4)
			valormatri := ast.Lista_Matriz_Val.Back().Value.(environment.Valores_Matriz).Valor.Valor
			ast.SustituirValores4(&matriz, valormatri, n1, n2, n3, n4)
		}
		if n1 > 0 && n2 > 0 && n3 > 0 && n4 == -1 && n5 == -1 {
			matriz = ast.NuevaMatriz(v.Name, true, tipo, n1, n2, n3)
			valormatri := ast.Lista_Matriz_Val.Back().Value.(environment.Valores_Matriz).Valor.Valor
			ast.SustituirValores3(&matriz, valormatri, n1, n2, n3)
		}
		if n1 > 0 && n2 > 0 && n3 == -1 && n4 == -1 && n5 == -1 {
			matriz = ast.NuevaMatriz(v.Name, true, tipo, n1, n2)
			valormatri := ast.Lista_Matriz_Val.Back().Value.(environment.Valores_Matriz).Valor.Valor
			ast.SustituirValores2(&matriz, valormatri, n1, n2)
		}

		if ast.Lista_Matriz_Val.Len() == tipo.Valor {
			ast.GuardarMatriz(matriz)
			ast.QuitarNiveles()
			return nil
		} else {
			Errores := environment.Errores{
				Descripcion: "Los niveles declarados no son los mismos ingresados",
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

	if Condiciones.Elements != nil {

		//comprobacion del tipo en todos
		for nivel := ast.Pila_Matriz_Val.Front(); nivel != nil; nivel = nivel.Next() {
			lista := nivel.Value.(*list.List)
			for elem := lista.Front(); elem != nil; elem = elem.Next() {
				valores := elem.Value.(environment.Valores_Matriz)
				if valores.Elements != nil {
					for e := valores.Elements.Front(); e != nil; e = e.Next() {
						if tipo.Tipo != e.Value.(environment.Symbol).Tipo {
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
		contadorpila := tipo.Valor.(int) - 1
		pila := ast.Pila_Matriz_Val.Len()
		lista := ast.Lista_Matriz_Val.Len()
		elementos := ast.Lista_Matriz_Val.Back().Value.(environment.Valores_Matriz).Elements.Len()
		pila2 := tipo.Valor.(int)

		if pila <= 3 {
			matriz = ast.NuevaMatriz(v.Name, true, tipo, lista, elementos)
		} else {
			matriz = ast.NuevaMatriz(v.Name, true, tipo, pila2, lista, elementos)
		}

		for nivel := ast.Pila_Matriz_Val.Front(); nivel != nil; nivel = nivel.Next() {
			lista := nivel.Value.(*list.List)
			contadorlista := 0
			for elem := lista.Front(); elem != nil; elem = elem.Next() {
				valores := elem.Value.(environment.Valores_Matriz)
				if valores.Elements != nil {
					contadorvalor := 0
					for e := valores.Elements.Front(); e != nil; e = e.Next() {
						if tipo.Valor.(int) == 2 {
							ast.IngresarValor(&matriz, []int{contadorlista, contadorvalor}, e.Value.(environment.Symbol).Valor)
						} else {
							ast.IngresarValor(&matriz, []int{contadorpila, contadorlista, contadorvalor}, e.Value.(environment.Symbol).Valor)
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
		return nil
	}

	ast.QuitarNiveles()
	return nil
}
