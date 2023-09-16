package expressions

import (
	"Backend/environment"
	"Backend/generator"
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

func (o Operation) Ejecutar(ast *environment.AST, gen *generator.Generator) environment.Value {
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

	var op1, op2, result environment.Value
	newTemp := gen.NewTemp()

	switch o.Operador {
	case "+":
		{
			op1 = o.Op_izq.Ejecutar(ast, gen)
			op2 = o.Op_der.Ejecutar(ast, gen)
			//validar tipo dominante
			dominante = tabla_dominante[op1.Type][op2.Type]
			//valida el tipo
			if dominante == environment.INTEGER {
				gen.AddExpression(newTemp, op1.Value, op2.Value, "+")
				result = environment.NewValue(newTemp, true, dominante, false, false, false)
				result.IntValue = op1.IntValue + op2.IntValue
				return result
			} else if dominante == environment.FLOAT {
				gen.AddExpression(newTemp, op1.Value, op2.Value, "+")
				result = environment.NewValue(newTemp, true, dominante, false, false, false)
				result.IntValue = op1.IntValue + op2.IntValue
				return result
			} else {
				r1 := fmt.Sprintf("%v", op1.Value)
				r2 := fmt.Sprintf("%v", op2.Value)
				Errores := environment.Errores{
					Descripcion: "No es posible sumar los dos valores, operacion1:" + r1 + ", operacion2:" + r2 + ".",
					Fila:        strconv.Itoa(o.Lin),
					Columna:     strconv.Itoa(o.Col),
					Tipo:        "Error Semantico",
					Ambito:      ast.ObtenerAmbito(),
				}
				ast.ErroresHTML(Errores)
			}
		}
	case "-":
		{
			op1 = o.Op_izq.Ejecutar(ast, gen)
			op2 = o.Op_der.Ejecutar(ast, gen)
			//validar tipo dominante
			dominante = tabla_dominante[op1.Type][op2.Type]
			//valida el tipo
			if dominante == environment.INTEGER {
				gen.AddExpression(newTemp, op1.Value, op2.Value, "-")
				result = environment.NewValue(newTemp, true, dominante, false, false, false)
				result.IntValue = op1.IntValue + op2.IntValue
				return result
			} else if dominante == environment.FLOAT {
				gen.AddExpression(newTemp, op1.Value, op2.Value, "-")
				result = environment.NewValue(newTemp, true, dominante, false, false, false)
				result.IntValue = op1.IntValue + op2.IntValue
				return result
			} else {
				r1 := fmt.Sprintf("%v", op1.Value)
				r2 := fmt.Sprintf("%v", op2.Value)
				Errores := environment.Errores{
					Descripcion: "No es posible sumar los dos valores, operacion1:" + r1 + ", operacion2:" + r2 + ".",
					Fila:        strconv.Itoa(o.Lin),
					Columna:     strconv.Itoa(o.Col),
					Tipo:        "Error Semantico",
					Ambito:      ast.ObtenerAmbito(),
				}
				ast.ErroresHTML(Errores)
			}
		}
	case "*":
		{
			op1 = o.Op_izq.Ejecutar(ast, gen)
			op2 = o.Op_der.Ejecutar(ast, gen)
			//validar tipo dominante
			dominante = tabla_dominante[op1.Type][op2.Type]
			//valida el tipo
			if dominante == environment.INTEGER {
				gen.AddExpression(newTemp, op1.Value, op2.Value, "*")
				result = environment.NewValue(newTemp, true, dominante, false, false, false)
				result.IntValue = op1.IntValue + op2.IntValue
				return result
			} else if dominante == environment.FLOAT {
				gen.AddExpression(newTemp, op1.Value, op2.Value, "*")
				result = environment.NewValue(newTemp, true, dominante, false, false, false)
				result.IntValue = op1.IntValue + op2.IntValue
				return result
			} else {
				r1 := fmt.Sprintf("%v", op1.Value)
				r2 := fmt.Sprintf("%v", op2.Value)
				Errores := environment.Errores{
					Descripcion: "No es posible sumar los dos valores, operacion1:" + r1 + ", operacion2:" + r2 + ".",
					Fila:        strconv.Itoa(o.Lin),
					Columna:     strconv.Itoa(o.Col),
					Tipo:        "Error Semantico",
					Ambito:      ast.ObtenerAmbito(),
				}
				ast.ErroresHTML(Errores)
			}
		}
	case "/":
		{
			op1 = o.Op_izq.Ejecutar(ast, gen)
			op2 = o.Op_der.Ejecutar(ast, gen)
			//validar tipo dominante
			dominante = tabla_dominante[op1.Type][op2.Type]
			//valida el tipo
			if dominante == environment.INTEGER {
				gen.AddExpression(newTemp, op1.Value, op2.Value, "/")
				result = environment.NewValue(newTemp, true, dominante, false, false, false)
				result.IntValue = op1.IntValue + op2.IntValue
				return result
			} else if dominante == environment.FLOAT {
				gen.AddExpression(newTemp, op1.Value, op2.Value, "/")
				result = environment.NewValue(newTemp, true, dominante, false, false, false)
				result.IntValue = op1.IntValue + op2.IntValue
				return result
			} else {
				r1 := fmt.Sprintf("%v", op1.Value)
				r2 := fmt.Sprintf("%v", op2.Value)
				Errores := environment.Errores{
					Descripcion: "No es posible sumar los dos valores, operacion1:" + r1 + ", operacion2:" + r2 + ".",
					Fila:        strconv.Itoa(o.Lin),
					Columna:     strconv.Itoa(o.Col),
					Tipo:        "Error Semantico",
					Ambito:      ast.ObtenerAmbito(),
				}
				ast.ErroresHTML(Errores)
			}
		}
	case "%":
		{
			op1 = o.Op_izq.Ejecutar(ast, gen)
			op2 = o.Op_der.Ejecutar(ast, gen)
			//validar tipo dominante
			dominante = tabla_dominante[op1.Type][op2.Type]
			//valida el tipo
			if dominante == environment.INTEGER {
				gen.AddExpression(newTemp, op1.Value, op2.Value, "%")
				result = environment.NewValue(newTemp, true, dominante, false, false, false)
				result.IntValue = op1.IntValue + op2.IntValue
				return result
			} else if dominante == environment.FLOAT {
				gen.AddExpression(newTemp, op1.Value, op2.Value, "%")
				result = environment.NewValue(newTemp, true, dominante, false, false, false)
				result.IntValue = op1.IntValue + op2.IntValue
				return result
			} else {
				r1 := fmt.Sprintf("%v", op1.Value)
				r2 := fmt.Sprintf("%v", op2.Value)
				Errores := environment.Errores{
					Descripcion: "No es posible sumar los dos valores, operacion1:" + r1 + ", operacion2:" + r2 + ".",
					Fila:        strconv.Itoa(o.Lin),
					Columna:     strconv.Itoa(o.Col),
					Tipo:        "Error Semantico",
					Ambito:      ast.ObtenerAmbito(),
				}
				ast.ErroresHTML(Errores)
			}
		}
	case "<":
		{
			return result
		}
	case ">":
		{
			return result
		}
	case "<=":
		{
			return result
		}
	case ">=":
		{
			return result
		}
	case "==":
		{
			return result
		}
	case "!=":
		{
			return result
		}
	case "&&":
		{
			return result
		}
	case "||":
		{
			return result
		}
	case "!":
		{
			return result
		}
	}

	return result
}
