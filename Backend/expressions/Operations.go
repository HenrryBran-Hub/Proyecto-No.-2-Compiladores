package expressions

import (
	"Backend/environment"
	"Backend/interfaces"
	"fmt"
	"strconv"
)

type Operation struct {
	Lin      int
	Col      int
	Op_izq   interfaces.Expression
	Operador string
	Op_der   interfaces.Expression
}

func NewOperation(lin int, col int, Op1 interfaces.Expression, Operador string, Op2 interfaces.Expression) Operation {
	exp := Operation{Lin: lin, Col: col, Op_izq: Op1, Operador: Operador, Op_der: Op2}
	return exp
}

func (o Operation) Ejecutar(ast *environment.AST) environment.Symbol {
	var dominante environment.TipoExpresion

	tabla_dominante := [6][6]environment.TipoExpresion{
		//INTEGER					FLOAT				STRING 					BOOLEANG				CHARACTER			NULL
		/*INTEGER*/ {environment.INTEGER, environment.FLOAT, environment.STRING, environment.BOOLEAN, environment.NULL, environment.NULL},
		/*FLOAT*/ {environment.FLOAT, environment.FLOAT, environment.STRING, environment.NULL, environment.NULL, environment.NULL},
		/*STRING*/ {environment.STRING, environment.STRING, environment.STRING, environment.STRING, environment.STRING, environment.NULL},
		/*BOOLEAN*/ {environment.BOOLEAN, environment.NULL, environment.STRING, environment.BOOLEAN, environment.NULL, environment.NULL},
		/*CHAR*/ {environment.NULL, environment.NULL, environment.STRING, environment.NULL, environment.STRING, environment.NULL},
		/*NULL*/ {environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
	}

	var op1, op2 environment.Symbol
	op1 = o.Op_izq.Ejecutar(ast)
	op2 = o.Op_der.Ejecutar(ast)
	//quitamos el operador coma
	switch o.Operador {
	case "+":
		{
			//validar tipo dominante
			dominante = tabla_dominante[op1.Tipo][op2.Tipo]
			//valida el tipo
			if dominante == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: dominante, Valor: op1.Valor.(int) + op2.Valor.(int)}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Valor), 64)
				num2 := fmt.Sprintf("%.6f", val1+val2)
				num3, err := strconv.ParseFloat(num2, 64)
				if err != nil {
					fmt.Println(err)
				}
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: dominante, Valor: num3}
			} else if dominante == environment.STRING {
				r1 := fmt.Sprintf("%v", op1.Valor)
				r2 := fmt.Sprintf("%v", op2.Valor)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: dominante, Valor: r1 + r2}
			} else {
				r1 := fmt.Sprintf("%v", op1.Valor)
				r2 := fmt.Sprintf("%v", op2.Valor)
				Errores := environment.Errores{
					Descripcion: "No es posible sumar los dos valores, operacion1:" + r1 + ", operacion2:" + r2 + ".",
					Fila:        strconv.Itoa(op1.Lin),
					Columna:     strconv.Itoa(op1.Col),
					Tipo:        "Error Semantico",
					Ambito:      op1.Scope,
				}
				ast.ErroresHTML(Errores)
			}
		}
	case "-":
		{
			dominante = tabla_dominante[op1.Tipo][op2.Tipo]

			if dominante == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: dominante, Valor: op1.Valor.(int) - op2.Valor.(int)}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Valor), 64)
				num2 := fmt.Sprintf("%.6f", val1-val2)
				num3, err := strconv.ParseFloat(num2, 64)
				if err != nil {
					fmt.Println(err)
				}
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: dominante, Valor: num3}
			} else {
				r1 := fmt.Sprintf("%v", op1.Valor)
				r2 := fmt.Sprintf("%v", op2.Valor)
				Errores := environment.Errores{
					Descripcion: "No es posible restar los dos valores, operacion1:" + r1 + ", operacion2:" + r2 + ".",
					Fila:        strconv.Itoa(op1.Lin),
					Columna:     strconv.Itoa(op1.Col),
					Tipo:        "Error Semantico",
					Ambito:      op1.Scope,
				}
				ast.ErroresHTML(Errores)
			}
		}
	case "*":
		{
			dominante = tabla_dominante[op1.Tipo][op2.Tipo]
			if dominante == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: dominante, Valor: op1.Valor.(int) * op2.Valor.(int)}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Valor), 64)
				num2 := fmt.Sprintf("%.6f", val1*val2)
				num3, err := strconv.ParseFloat(num2, 64)
				if err != nil {
					fmt.Println(err)
				}
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: dominante, Valor: num3}
			} else {
				r1 := fmt.Sprintf("%v", op1.Valor)
				r2 := fmt.Sprintf("%v", op2.Valor)
				Errores := environment.Errores{
					Descripcion: "No es posible multiplicar los dos valores, operacion1:" + r1 + ", operacion2:" + r2 + ".",
					Fila:        strconv.Itoa(op1.Lin),
					Columna:     strconv.Itoa(op1.Col),
					Tipo:        "Error Semantico",
					Ambito:      op1.Scope,
				}
				ast.ErroresHTML(Errores)
			}
		}
	case "/":
		{
			dominante = tabla_dominante[op1.Tipo][op2.Tipo]
			if dominante == environment.INTEGER {
				if op2.Valor.(int) != 0 {
					return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: dominante, Valor: op1.Valor.(int) / op2.Valor.(int)}
				} else {
					r1 := fmt.Sprintf("%v", op1.Valor)
					r2 := fmt.Sprintf("%v", op2.Valor)
					Errores := environment.Errores{
						Descripcion: "No es posible dividir entre 0, operacion1:" + r1 + ", operacion2:" + r2 + ".",
						Fila:        strconv.Itoa(op1.Lin),
						Columna:     strconv.Itoa(op1.Col),
						Tipo:        "Error Semantico",
						Ambito:      op1.Scope,
					}
					ast.ErroresHTML(Errores)
				}

			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Valor), 64)
				if val2 != 0 {
					num2 := fmt.Sprintf("%.6f", val1/val2)
					num3, err := strconv.ParseFloat(num2, 64)
					if err != nil {
						fmt.Println(err)
					}
					return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: dominante, Valor: num3}
				} else {
					r1 := fmt.Sprintf("%v", op1.Valor)
					r2 := fmt.Sprintf("%v", op2.Valor)
					Errores := environment.Errores{
						Descripcion: "No es posible dividir entre 0, operacion1:" + r1 + ", operacion2:" + r2 + ".",
						Fila:        strconv.Itoa(op1.Lin),
						Columna:     strconv.Itoa(op1.Col),
						Tipo:        "Error Semantico",
						Ambito:      op1.Scope,
					}
					ast.ErroresHTML(Errores)
				}
			} else {
				r1 := fmt.Sprintf("%v", op1.Valor)
				r2 := fmt.Sprintf("%v", op2.Valor)
				Errores := environment.Errores{
					Descripcion: "No es posible dividir los dos valores, operacion1:" + r1 + ", operacion2:" + r2 + ".",
					Fila:        strconv.Itoa(op1.Lin),
					Columna:     strconv.Itoa(op1.Col),
					Tipo:        "Error Semantico",
					Ambito:      op1.Scope,
				}
				ast.ErroresHTML(Errores)
			}

		}
	case "%":
		{
			dominante = tabla_dominante[op1.Tipo][op2.Tipo]
			if dominante == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: dominante, Valor: op1.Valor.(int) % op2.Valor.(int)}
			} else {
				r1 := fmt.Sprintf("%v", op1.Valor)
				r2 := fmt.Sprintf("%v", op2.Valor)
				Errores := environment.Errores{
					Descripcion: "No es posible modular los dos valores, operacion1:" + r1 + ", operacion2:" + r2 + ".",
					Fila:        strconv.Itoa(op1.Lin),
					Columna:     strconv.Itoa(op1.Col),
					Tipo:        "Error Semantico",
					Ambito:      op1.Scope,
				}
				ast.ErroresHTML(Errores)
			}
		}
	case "<":
		{
			dominante = tabla_dominante[op1.Tipo][op2.Tipo]
			if dominante == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: op1.Valor.(int) < op2.Valor.(int)}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Valor), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: val1 < val2}
			} else if dominante == environment.STRING && (op1.Tipo == op2.Tipo) {
				r1 := fmt.Sprintf("%v", op1.Valor)
				r2 := fmt.Sprintf("%v", op2.Valor)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: r1 < r2}
			} else if dominante == environment.CHARACTER && (op1.Tipo == op2.Tipo) {
				r1 := rune(fmt.Sprintf("%v", op1.Valor)[0])
				r2 := rune(fmt.Sprintf("%v", op1.Valor)[0])
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: r1 < r2}
			} else {
				r1 := fmt.Sprintf("%v", op1.Valor)
				r2 := fmt.Sprintf("%v", op2.Valor)
				Errores := environment.Errores{
					Descripcion: "No es posible comparar los dos valores(<), operacion1:" + r1 + ", operacion2:" + r2 + ".",
					Fila:        strconv.Itoa(op1.Lin),
					Columna:     strconv.Itoa(op1.Col),
					Tipo:        "Error Semantico",
					Ambito:      op1.Scope,
				}
				ast.ErroresHTML(Errores)
			}
		}
	case ">":
		{
			dominante = tabla_dominante[op1.Tipo][op2.Tipo]
			if dominante == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: op1.Valor.(int) > op2.Valor.(int)}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Valor), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: val1 > val2}
			} else if dominante == environment.STRING && (op1.Tipo == op2.Tipo) {
				r1 := fmt.Sprintf("%v", op1.Valor)
				r2 := fmt.Sprintf("%v", op2.Valor)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: r1 > r2}
			} else if dominante == environment.CHARACTER && (op1.Tipo == op2.Tipo) {
				r1 := rune(fmt.Sprintf("%v", op1.Valor)[0])
				r2 := rune(fmt.Sprintf("%v", op1.Valor)[0])
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: r1 > r2}
			} else {
				r1 := fmt.Sprintf("%v", op1.Valor)
				r2 := fmt.Sprintf("%v", op2.Valor)
				Errores := environment.Errores{
					Descripcion: "No es posible comparar los dos valores(>), operacion1:" + r1 + ", operacion2:" + r2 + ".",
					Fila:        strconv.Itoa(op1.Lin),
					Columna:     strconv.Itoa(op1.Col),
					Tipo:        "Error Semantico",
					Ambito:      op1.Scope,
				}
				ast.ErroresHTML(Errores)
			}
		}
	case "<=":
		{
			dominante = tabla_dominante[op1.Tipo][op2.Tipo]
			if dominante == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: op1.Valor.(int) <= op2.Valor.(int)}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Valor), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: val1 <= val2}
			} else if dominante == environment.STRING && (op1.Tipo == op2.Tipo) {
				r1 := fmt.Sprintf("%v", op1.Valor)
				r2 := fmt.Sprintf("%v", op2.Valor)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: r1 <= r2}
			} else if dominante == environment.CHARACTER && (op1.Tipo == op2.Tipo) {
				r1 := rune(fmt.Sprintf("%v", op1.Valor)[0])
				r2 := rune(fmt.Sprintf("%v", op1.Valor)[0])
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: r1 <= r2}
			} else {
				r1 := fmt.Sprintf("%v", op1.Valor)
				r2 := fmt.Sprintf("%v", op2.Valor)
				Errores := environment.Errores{
					Descripcion: "No es posible comparar los dos valores(<=), operacion1:" + r1 + ", operacion2:" + r2 + ".",
					Fila:        strconv.Itoa(op1.Lin),
					Columna:     strconv.Itoa(op1.Col),
					Tipo:        "Error Semantico",
					Ambito:      op1.Scope,
				}
				ast.ErroresHTML(Errores)
			}
		}
	case ">=":
		{
			dominante = tabla_dominante[op1.Tipo][op2.Tipo]
			if dominante == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: op1.Valor.(int) >= op2.Valor.(int)}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Valor), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: val1 >= val2}
			} else if dominante == environment.STRING && (op1.Tipo == op2.Tipo) {
				r1 := fmt.Sprintf("%v", op1.Valor)
				r2 := fmt.Sprintf("%v", op2.Valor)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: r1 >= r2}
			} else if dominante == environment.CHARACTER && (op1.Tipo == op2.Tipo) {
				r1 := rune(fmt.Sprintf("%v", op1.Valor)[0])
				r2 := rune(fmt.Sprintf("%v", op1.Valor)[0])
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: r1 >= r2}
			} else {
				r1 := fmt.Sprintf("%v", op1.Valor)
				r2 := fmt.Sprintf("%v", op2.Valor)
				Errores := environment.Errores{
					Descripcion: "No es posible comparar los dos valores(>=), operacion1:" + r1 + ", operacion2:" + r2 + ".",
					Fila:        strconv.Itoa(op1.Lin),
					Columna:     strconv.Itoa(op1.Col),
					Tipo:        "Error Semantico",
					Ambito:      op1.Scope,
				}
				ast.ErroresHTML(Errores)
			}
		}
	case "==":
		{
			if op1.Tipo == op2.Tipo {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: op1.Valor == op2.Valor}
			} else {
				r1 := fmt.Sprintf("%v", op1.Valor)
				r2 := fmt.Sprintf("%v", op2.Valor)
				Errores := environment.Errores{
					Descripcion: "No es posible comparar los dos valores(==), operacion1:" + r1 + ", operacion2:" + r2 + ".",
					Fila:        strconv.Itoa(op1.Lin),
					Columna:     strconv.Itoa(op1.Col),
					Tipo:        "Error Semantico",
					Ambito:      op1.Scope,
				}
				ast.ErroresHTML(Errores)
			}
		}
	case "!=":
		{
			if op1.Tipo == op2.Tipo {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: op1.Valor != op2.Valor}
			} else {
				r1 := fmt.Sprintf("%v", op1.Valor)
				r2 := fmt.Sprintf("%v", op2.Valor)
				Errores := environment.Errores{
					Descripcion: "No es posible comparar los dos valores(!=), operacion1:" + r1 + ", operacion2:" + r2 + ".",
					Fila:        strconv.Itoa(op1.Lin),
					Columna:     strconv.Itoa(op1.Col),
					Tipo:        "Error Semantico",
					Ambito:      op1.Scope,
				}
				ast.ErroresHTML(Errores)
			}
		}
	case "&&":
		{
			if (op1.Tipo == environment.BOOLEAN) && (op2.Tipo == environment.BOOLEAN) {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: op1.Valor.(bool) && op2.Valor.(bool)}
			} else {
				r1 := fmt.Sprintf("%v", op1.Valor)
				r2 := fmt.Sprintf("%v", op2.Valor)
				Errores := environment.Errores{
					Descripcion: "Los valores no son compatibles(&&), operacion1:" + r1 + ", operacion2:" + r2 + ".",
					Fila:        strconv.Itoa(op1.Lin),
					Columna:     strconv.Itoa(op1.Col),
					Tipo:        "Error Semantico",
					Ambito:      op1.Scope,
				}
				ast.ErroresHTML(Errores)
			}
		}
	case "||":
		{
			if (op1.Tipo == environment.BOOLEAN) && (op2.Tipo == environment.BOOLEAN) {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: op1.Valor.(bool) || op2.Valor.(bool)}
			} else {
				r1 := fmt.Sprintf("%v", op1.Valor)
				r2 := fmt.Sprintf("%v", op2.Valor)
				Errores := environment.Errores{
					Descripcion: "Los valores no son compatibles(&&), operacion1:" + r1 + ", operacion2:" + r2 + ".",
					Fila:        strconv.Itoa(op1.Lin),
					Columna:     strconv.Itoa(op1.Col),
					Tipo:        "Error Semantico",
					Ambito:      op1.Scope,
				}
				ast.ErroresHTML(Errores)
			}
		}
	case "!":
		{
			if op1.Tipo == environment.BOOLEAN {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: !op1.Valor.(bool)}
			} else {
				r1 := fmt.Sprintf("%v", op1.Valor)
				Errores := environment.Errores{
					Descripcion: "Los valores no son compatibles(!), operacion1:" + r1 + ".",
					Fila:        strconv.Itoa(op1.Lin),
					Columna:     strconv.Itoa(op1.Col),
					Tipo:        "Error Semantico",
					Ambito:      op1.Scope,
				}
				ast.ErroresHTML(Errores)
			}
		}
	}

	var result interface{}
	return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.NULL, Valor: result}
}
