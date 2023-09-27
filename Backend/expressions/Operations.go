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
				newTemp1 := gen.NewTemp()
				newTemp2 := gen.NewTemp()

				if op1.IsTemp && ast.IsTempT(op1.Value) && op1.Val.Name == "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else if !op1.IsTemp && op1.Value != "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else {
					gen.AddGetStack(newTemp1, strconv.Itoa(op1.Val.Symbols.Posicion))
				}

				if op2.IsTemp && ast.IsTempT(op2.Value) && op2.Val.Name == "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else if !op2.IsTemp && op2.Value != "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else {
					gen.AddGetStack(newTemp2, strconv.Itoa(op2.Val.Symbols.Posicion))
				}

				gen.AddExpression(newTemp, newTemp1, newTemp2, "+")
				result = environment.NewValue(newTemp, true, dominante, false, false, false, environment.Variable{})
				result.IntValue = op1.IntValue + op2.IntValue
				return result
			} else if dominante == environment.FLOAT {
				newTemp1 := gen.NewTemp()
				newTemp2 := gen.NewTemp()

				if op1.IsTemp && ast.IsTempT(op1.Value) && op1.Val.Name == "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else if !op1.IsTemp && op1.Value != "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else {
					gen.AddGetStack(newTemp1, strconv.Itoa(op1.Val.Symbols.Posicion))
				}

				if op2.IsTemp && ast.IsTempT(op2.Value) && op2.Val.Name == "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else if !op2.IsTemp && op2.Value != "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else {
					gen.AddGetStack(newTemp2, strconv.Itoa(op2.Val.Symbols.Posicion))
				}

				gen.AddExpression(newTemp, newTemp1, newTemp2, "+")
				result = environment.NewValue(newTemp, true, dominante, false, false, false, environment.Variable{})
				result.IntValue = op1.IntValue + op2.IntValue
				return result
			} else if dominante == environment.STRING {
				//llamar a generar concatstring
				gen.GenerateConcatString()
				//concat
				gen.AddComment("Concatenando strings")
				envSize := strconv.Itoa(ast.PosicionStack + 1)
				tmp1 := gen.NewTemp()
				tmp2 := gen.NewTemp()
				gen.AddExpression(tmp1, "P", envSize, "+")
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op1.Value)
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op2.Value)
				gen.AddExpression("P", "P", envSize, "+")
				gen.AddCall("concatString")
				gen.AddGetStack(tmp2, "(int)P")
				gen.AddExpression("P", "P", envSize, "-")
				gen.AddBr()
				result = environment.NewValue(tmp2, true, dominante, false, false, false, environment.Variable{})
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
				newTemp1 := gen.NewTemp()
				newTemp2 := gen.NewTemp()

				if op1.IsTemp && ast.IsTempT(op1.Value) && op1.Val.Name == "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else if !op1.IsTemp && op1.Value != "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else {
					gen.AddGetStack(newTemp1, strconv.Itoa(op1.Val.Symbols.Posicion))
				}

				if op2.IsTemp && ast.IsTempT(op2.Value) && op2.Val.Name == "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else if !op2.IsTemp && op2.Value != "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else {
					gen.AddGetStack(newTemp2, strconv.Itoa(op2.Val.Symbols.Posicion))
				}

				gen.AddExpression(newTemp, newTemp1, newTemp2, "-")
				result = environment.NewValue(newTemp, true, dominante, false, false, false, environment.Variable{})
				result.IntValue = op1.IntValue + op2.IntValue
				return result
			} else if dominante == environment.FLOAT {
				newTemp1 := gen.NewTemp()
				newTemp2 := gen.NewTemp()

				if op1.IsTemp && ast.IsTempT(op1.Value) && op1.Val.Name == "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else if !op1.IsTemp && op1.Value != "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else {
					gen.AddGetStack(newTemp1, strconv.Itoa(op1.Val.Symbols.Posicion))
				}

				if op2.IsTemp && ast.IsTempT(op2.Value) && op2.Val.Name == "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else if !op2.IsTemp && op2.Value != "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else {
					gen.AddGetStack(newTemp2, strconv.Itoa(op2.Val.Symbols.Posicion))
				}

				gen.AddExpression(newTemp, newTemp1, newTemp2, "-")
				result = environment.NewValue(newTemp, true, dominante, false, false, false, environment.Variable{})
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
				newTemp1 := gen.NewTemp()
				newTemp2 := gen.NewTemp()

				if op1.IsTemp && ast.IsTempT(op1.Value) && op1.Val.Name == "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else if !op1.IsTemp && op1.Value != "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else {
					gen.AddGetStack(newTemp1, strconv.Itoa(op1.Val.Symbols.Posicion))
				}

				if op2.IsTemp && ast.IsTempT(op2.Value) && op2.Val.Name == "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else if !op2.IsTemp && op2.Value != "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else {
					gen.AddGetStack(newTemp2, strconv.Itoa(op2.Val.Symbols.Posicion))
				}

				gen.AddExpression(newTemp, newTemp1, newTemp2, "*")
				result = environment.NewValue(newTemp, true, dominante, false, false, false, environment.Variable{})
				result.IntValue = op1.IntValue + op2.IntValue
				return result
			} else if dominante == environment.FLOAT {
				newTemp1 := gen.NewTemp()
				newTemp2 := gen.NewTemp()

				if op1.IsTemp && ast.IsTempT(op1.Value) && op1.Val.Name == "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else if !op1.IsTemp && op1.Value != "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else {
					gen.AddGetStack(newTemp1, strconv.Itoa(op1.Val.Symbols.Posicion))
				}

				if op2.IsTemp && ast.IsTempT(op2.Value) && op2.Val.Name == "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else if !op2.IsTemp && op2.Value != "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else {
					gen.AddGetStack(newTemp2, strconv.Itoa(op2.Val.Symbols.Posicion))
				}

				gen.AddExpression(newTemp, newTemp1, newTemp2, "*")
				result = environment.NewValue(newTemp, true, dominante, false, false, false, environment.Variable{})
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
				if op2.Value != "0" {
					newTemp1 := gen.NewTemp()
					newTemp2 := gen.NewTemp()

					if op1.IsTemp && ast.IsTempT(op1.Value) && op1.Val.Name == "" {
						gen.AddAssign(newTemp1, op1.Value)
					} else if !op1.IsTemp && op1.Value != "" {
						gen.AddAssign(newTemp1, op1.Value)
					} else {
						gen.AddGetStack(newTemp1, strconv.Itoa(op1.Val.Symbols.Posicion))
					}

					if op2.IsTemp && ast.IsTempT(op2.Value) && op2.Val.Name == "" {
						gen.AddAssign(newTemp2, op2.Value)
					} else if !op2.IsTemp && op2.Value != "" {
						gen.AddAssign(newTemp2, op2.Value)
					} else {
						gen.AddGetStack(newTemp2, strconv.Itoa(op2.Val.Symbols.Posicion))
					}

					gen.AddExpression(newTemp, newTemp1, newTemp2, "/")
					result = environment.NewValue(newTemp, true, dominante, false, false, false, environment.Variable{})
					result.IntValue = op1.IntValue + op2.IntValue
					return result
				} else {
					Errores := environment.Errores{
						Descripcion: "No es posible dividir entre 0, operacion1 entre operacion 2.",
						Fila:        strconv.Itoa(o.Lin),
						Columna:     strconv.Itoa(o.Col),
						Tipo:        "Error Semantico",
						Ambito:      ast.ObtenerAmbito(),
					}
					ast.ErroresHTML(Errores)
					return result
				}
			} else if dominante == environment.FLOAT {
				if op2.Value != "0" {
					newTemp1 := gen.NewTemp()
					newTemp2 := gen.NewTemp()

					if op1.IsTemp && ast.IsTempT(op1.Value) && op1.Val.Name == "" {
						gen.AddAssign(newTemp1, op1.Value)
					} else if !op1.IsTemp && op1.Value != "" {
						gen.AddAssign(newTemp1, op1.Value)
					} else {
						gen.AddGetStack(newTemp1, strconv.Itoa(op1.Val.Symbols.Posicion))
					}

					if op2.IsTemp && ast.IsTempT(op2.Value) && op2.Val.Name == "" {
						gen.AddAssign(newTemp2, op2.Value)
					} else if !op2.IsTemp && op2.Value != "" {
						gen.AddAssign(newTemp2, op2.Value)
					} else {
						gen.AddGetStack(newTemp2, strconv.Itoa(op2.Val.Symbols.Posicion))
					}

					gen.AddExpression(newTemp, newTemp1, newTemp2, "/")
					result = environment.NewValue(newTemp, true, dominante, false, false, false, environment.Variable{})
					result.IntValue = op1.IntValue + op2.IntValue
					return result
				} else {
					Errores := environment.Errores{
						Descripcion: "No es posible dividir entre 0, operacion1 entre operacion 2.",
						Fila:        strconv.Itoa(o.Lin),
						Columna:     strconv.Itoa(o.Col),
						Tipo:        "Error Semantico",
						Ambito:      ast.ObtenerAmbito(),
					}
					ast.ErroresHTML(Errores)
					return result
				}
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
				newTemp1 := gen.NewTemp()
				newTemp2 := gen.NewTemp()

				if op1.IsTemp && ast.IsTempT(op1.Value) && op1.Val.Name == "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else if !op1.IsTemp && op1.Value != "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else {
					gen.AddGetStack(newTemp1, strconv.Itoa(op1.Val.Symbols.Posicion))
				}

				if op2.IsTemp && ast.IsTempT(op2.Value) && op2.Val.Name == "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else if !op2.IsTemp && op2.Value != "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else {
					gen.AddGetStack(newTemp2, strconv.Itoa(op2.Val.Symbols.Posicion))
				}

				gen.AddExpression(newTemp, "(int)"+newTemp1, "(int)"+newTemp2, "%")
				result = environment.NewValue(newTemp, true, dominante, false, false, false, environment.Variable{})
				return result
			} else if dominante == environment.FLOAT {
				newTemp1 := gen.NewTemp()
				newTemp2 := gen.NewTemp()

				if op1.IsTemp && ast.IsTempT(op1.Value) && op1.Val.Name == "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else if !op1.IsTemp && op1.Value != "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else {
					gen.AddGetStack(newTemp1, strconv.Itoa(op1.Val.Symbols.Posicion))
				}

				if op2.IsTemp && ast.IsTempT(op2.Value) && op2.Val.Name == "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else if !op2.IsTemp && op2.Value != "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else {
					gen.AddGetStack(newTemp2, strconv.Itoa(op2.Val.Symbols.Posicion))
				}

				gen.AddExpression(newTemp, "(int)"+newTemp1, "(int)"+newTemp2, "%")
				result = environment.NewValue(newTemp, true, dominante, false, false, false, environment.Variable{})
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
			op1 = o.Op_izq.Ejecutar(ast, gen)
			op2 = o.Op_der.Ejecutar(ast, gen)
			//validar tipo dominante
			dominante = tabla_dominante[op1.Type][op2.Type]
			//valida el tipo
			if dominante == environment.INTEGER {
				tlabel := gen.NewLabel()
				flabel := gen.NewLabel()
				newTemp1 := gen.NewTemp()
				newTemp2 := gen.NewTemp()

				if op1.IsTemp && ast.IsTempT(op1.Value) && op1.Val.Name == "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else if !op1.IsTemp && op1.Value != "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else {
					gen.AddGetStack(newTemp1, strconv.Itoa(op1.Val.Symbols.Posicion))
				}

				if op2.IsTemp && ast.IsTempT(op2.Value) && op2.Val.Name == "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else if !op2.IsTemp && op2.Value != "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else {
					gen.AddGetStack(newTemp2, strconv.Itoa(op2.Val.Symbols.Posicion))
				}

				gen.AddIf(newTemp1, newTemp2, "<", tlabel)
				gen.AddGoto(flabel)
				result = environment.NewValue("", false, environment.BOOLEAN, false, false, false, environment.Variable{})
				result.TrueLabel.PushBack(tlabel)
				result.FalseLabel.PushBack(flabel)
				result.Val.FEti = flabel
				result.Val.TEti = tlabel

				return result
			} else if dominante == environment.FLOAT {
				tlabel := gen.NewLabel()
				flabel := gen.NewLabel()
				newTemp1 := gen.NewTemp()
				newTemp2 := gen.NewTemp()

				if op1.IsTemp && ast.IsTempT(op1.Value) && op1.Val.Name == "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else if !op1.IsTemp && op1.Value != "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else {
					gen.AddGetStack(newTemp1, strconv.Itoa(op1.Val.Symbols.Posicion))
				}

				if op2.IsTemp && ast.IsTempT(op2.Value) && op2.Val.Name == "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else if !op2.IsTemp && op2.Value != "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else {
					gen.AddGetStack(newTemp2, strconv.Itoa(op2.Val.Symbols.Posicion))
				}

				gen.AddIf(newTemp1, newTemp2, "<", tlabel)
				gen.AddGoto(flabel)
				result = environment.NewValue("", false, environment.BOOLEAN, false, false, false, environment.Variable{})
				result.TrueLabel.PushBack(tlabel)
				result.FalseLabel.PushBack(flabel)
				result.Val.FEti = flabel
				result.Val.TEti = tlabel

				return result
			} else if dominante == environment.STRING && (op1.Type == op2.Type) {
				//llamar a generar concatstring
				gen.GenerateCompareMenorString()
				//concat
				gen.AddComment("Comparando strings")
				envSize := strconv.Itoa(ast.PosicionStack)
				tmp1 := gen.NewTemp()
				tmp2 := gen.NewTemp()
				gen.AddExpression(tmp1, "P", envSize, "+")
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op1.Value)
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op2.Value)
				gen.AddExpression("P", "P", envSize, "+")
				gen.AddCall("comparemenorString")
				gen.AddGetStack(tmp2, "(int)P")
				gen.AddExpression("P", "P", envSize, "-")
				gen.AddBr()
				//result = environment.NewValue(tmp2, true, environment.BOOLEAN, false, false, false, environment.Variable{})

				tlabel := gen.NewLabel()
				flabel := gen.NewLabel()

				gen.AddIf(tmp2, "1", "==", tlabel)
				gen.AddGoto(flabel)
				result = environment.NewValue(tmp2, false, environment.BOOLEAN, false, false, false, environment.Variable{})
				result.TrueLabel.PushBack(tlabel)
				result.FalseLabel.PushBack(flabel)
				result.Val.FEti = flabel
				result.Val.TEti = tlabel

				return result
			} else if dominante == environment.CHARACTER && (op1.Type == op2.Type) {
				//llamar a generar concatstring
				gen.GenerateCompareMenorString()
				//concat
				gen.AddComment("Comparando strings")
				envSize := strconv.Itoa(ast.PosicionStack)
				tmp1 := gen.NewTemp()
				tmp2 := gen.NewTemp()
				gen.AddExpression(tmp1, "P", envSize, "+")
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op1.Value)
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op2.Value)
				gen.AddExpression("P", "P", envSize, "+")
				gen.AddCall("comparemenorString")
				gen.AddGetStack(tmp2, "(int)P")
				gen.AddExpression("P", "P", envSize, "-")
				gen.AddBr()
				//result = environment.NewValue(tmp2, true, environment.BOOLEAN, false, false, false, environment.Variable{})

				tlabel := gen.NewLabel()
				flabel := gen.NewLabel()

				gen.AddIf(tmp2, "1", "==", tlabel)
				gen.AddGoto(flabel)
				result = environment.NewValue(tmp2, false, environment.BOOLEAN, false, false, false, environment.Variable{})
				result.TrueLabel.PushBack(tlabel)
				result.FalseLabel.PushBack(flabel)
				result.Val.FEti = flabel
				result.Val.TEti = tlabel

				return result
			} else {
				Errores := environment.Errores{
					Descripcion: "No es posible comparar los dos valores(<)",
					Fila:        strconv.Itoa(o.Lin),
					Columna:     strconv.Itoa(o.Col),
					Tipo:        "Error Semantico",
					Ambito:      ast.ObtenerAmbito(),
				}
				ast.ErroresHTML(Errores)
				return result
			}
		}
	case ">":
		{
			op1 = o.Op_izq.Ejecutar(ast, gen)
			op2 = o.Op_der.Ejecutar(ast, gen)
			//validar tipo dominante
			dominante = tabla_dominante[op1.Type][op2.Type]
			//valida el tipo
			if dominante == environment.INTEGER {
				tlabel := gen.NewLabel()
				flabel := gen.NewLabel()
				newTemp1 := gen.NewTemp()
				newTemp2 := gen.NewTemp()

				if op1.IsTemp && ast.IsTempT(op1.Value) && op1.Val.Name == "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else if !op1.IsTemp && op1.Value != "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else {
					gen.AddGetStack(newTemp1, strconv.Itoa(op1.Val.Symbols.Posicion))
				}

				if op2.IsTemp && ast.IsTempT(op2.Value) && op2.Val.Name == "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else if !op2.IsTemp && op2.Value != "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else {
					gen.AddGetStack(newTemp2, strconv.Itoa(op2.Val.Symbols.Posicion))
				}

				gen.AddIf(newTemp1, newTemp2, ">", tlabel)
				gen.AddGoto(flabel)
				result = environment.NewValue("", false, environment.BOOLEAN, false, false, false, environment.Variable{})
				result.TrueLabel.PushBack(tlabel)
				result.FalseLabel.PushBack(flabel)
				result.Val.FEti = flabel
				result.Val.TEti = tlabel

				return result
			} else if dominante == environment.FLOAT {
				tlabel := gen.NewLabel()
				flabel := gen.NewLabel()

				newTemp1 := gen.NewTemp()
				newTemp2 := gen.NewTemp()

				if op1.IsTemp && ast.IsTempT(op1.Value) && op1.Val.Name == "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else if !op1.IsTemp && op1.Value != "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else {
					gen.AddGetStack(newTemp1, strconv.Itoa(op1.Val.Symbols.Posicion))
				}

				if op2.IsTemp && ast.IsTempT(op2.Value) && op2.Val.Name == "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else if !op2.IsTemp && op2.Value != "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else {
					gen.AddGetStack(newTemp2, strconv.Itoa(op2.Val.Symbols.Posicion))
				}

				gen.AddIf(newTemp1, newTemp2, ">", tlabel)
				gen.AddGoto(flabel)
				result = environment.NewValue("", false, environment.BOOLEAN, false, false, false, environment.Variable{})
				result.TrueLabel.PushBack(tlabel)
				result.FalseLabel.PushBack(flabel)
				result.Val.FEti = flabel
				result.Val.TEti = tlabel

				return result
			} else if dominante == environment.STRING && (op1.Type == op2.Type) {
				//llamar a generar concatstring
				gen.GenerateCompareMaryorString()
				//concat
				gen.AddComment("Comparando strings")
				envSize := strconv.Itoa(ast.PosicionStack)
				tmp1 := gen.NewTemp()
				tmp2 := gen.NewTemp()
				gen.AddExpression(tmp1, "P", envSize, "+")
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op1.Value)
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op2.Value)
				gen.AddExpression("P", "P", envSize, "+")
				gen.AddCall("comparemayorString")
				gen.AddGetStack(tmp2, "(int)P")
				gen.AddExpression("P", "P", envSize, "-")
				gen.AddBr()
				//result = environment.NewValue(tmp2, true, environment.BOOLEAN, false, false, false, environment.Variable{})

				tlabel := gen.NewLabel()
				flabel := gen.NewLabel()

				gen.AddIf(tmp2, "1", "==", tlabel)
				gen.AddGoto(flabel)
				result = environment.NewValue(tmp2, false, environment.BOOLEAN, false, false, false, environment.Variable{})
				result.TrueLabel.PushBack(tlabel)
				result.FalseLabel.PushBack(flabel)
				result.Val.FEti = flabel
				result.Val.TEti = tlabel

				return result
			} else if dominante == environment.CHARACTER && (op1.Type == op2.Type) {
				//llamar a generar concatstring
				gen.GenerateCompareMaryorString()
				//concat
				gen.AddComment("Comparando strings")
				envSize := strconv.Itoa(ast.PosicionStack)
				tmp1 := gen.NewTemp()
				tmp2 := gen.NewTemp()
				gen.AddExpression(tmp1, "P", envSize, "+")
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op1.Value)
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op2.Value)
				gen.AddExpression("P", "P", envSize, "+")
				gen.AddCall("comparemayorString")
				gen.AddGetStack(tmp2, "(int)P")
				gen.AddExpression("P", "P", envSize, "-")
				gen.AddBr()
				//result = environment.NewValue(tmp2, true, environment.BOOLEAN, false, false, false, environment.Variable{})

				tlabel := gen.NewLabel()
				flabel := gen.NewLabel()

				gen.AddIf(tmp2, "1", "==", tlabel)
				gen.AddGoto(flabel)
				result = environment.NewValue(tmp2, false, environment.BOOLEAN, false, false, false, environment.Variable{})
				result.TrueLabel.PushBack(tlabel)
				result.FalseLabel.PushBack(flabel)
				result.Val.FEti = flabel
				result.Val.TEti = tlabel

				return result
			} else {
				Errores := environment.Errores{
					Descripcion: "No es posible comparar los dos valores(>)",
					Fila:        strconv.Itoa(o.Lin),
					Columna:     strconv.Itoa(o.Col),
					Tipo:        "Error Semantico",
					Ambito:      ast.ObtenerAmbito(),
				}
				ast.ErroresHTML(Errores)
				return result
			}
		}
	case "<=":
		{
			op1 = o.Op_izq.Ejecutar(ast, gen)
			op2 = o.Op_der.Ejecutar(ast, gen)
			//validar tipo dominante
			dominante = tabla_dominante[op1.Type][op2.Type]
			//valida el tipo
			if dominante == environment.INTEGER {
				tlabel := gen.NewLabel()
				flabel := gen.NewLabel()

				newTemp1 := gen.NewTemp()
				newTemp2 := gen.NewTemp()

				if op1.IsTemp && ast.IsTempT(op1.Value) && op1.Val.Name == "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else if !op1.IsTemp && op1.Value != "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else {
					gen.AddGetStack(newTemp1, strconv.Itoa(op1.Val.Symbols.Posicion))
				}

				if op2.IsTemp && ast.IsTempT(op2.Value) && op2.Val.Name == "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else if !op2.IsTemp && op2.Value != "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else {
					gen.AddGetStack(newTemp2, strconv.Itoa(op2.Val.Symbols.Posicion))
				}

				gen.AddIf(newTemp1, newTemp2, "<=", tlabel)
				gen.AddGoto(flabel)
				result = environment.NewValue("", false, environment.BOOLEAN, false, false, false, environment.Variable{})
				result.TrueLabel.PushBack(tlabel)
				result.FalseLabel.PushBack(flabel)
				result.Val.FEti = flabel
				result.Val.TEti = tlabel

				return result
			} else if dominante == environment.FLOAT {
				tlabel := gen.NewLabel()
				flabel := gen.NewLabel()

				newTemp1 := gen.NewTemp()
				newTemp2 := gen.NewTemp()

				if op1.IsTemp && ast.IsTempT(op1.Value) && op1.Val.Name == "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else if !op1.IsTemp && op1.Value != "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else {
					gen.AddGetStack(newTemp1, strconv.Itoa(op1.Val.Symbols.Posicion))
				}

				if op2.IsTemp && ast.IsTempT(op2.Value) && op2.Val.Name == "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else if !op2.IsTemp && op2.Value != "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else {
					gen.AddGetStack(newTemp2, strconv.Itoa(op2.Val.Symbols.Posicion))
				}

				gen.AddIf(newTemp1, newTemp2, "<=", tlabel)
				gen.AddGoto(flabel)
				result = environment.NewValue("", false, environment.BOOLEAN, false, false, false, environment.Variable{})
				result.TrueLabel.PushBack(tlabel)
				result.FalseLabel.PushBack(flabel)
				result.Val.FEti = flabel
				result.Val.TEti = tlabel

				return result
			} else if dominante == environment.STRING && (op1.Type == op2.Type) {
				//llamar a generar concatstring
				gen.GenerateCompareMenorigString()
				//concat
				gen.AddComment("Comparando strings")
				envSize := strconv.Itoa(ast.PosicionStack)
				tmp1 := gen.NewTemp()
				tmp2 := gen.NewTemp()
				gen.AddExpression(tmp1, "P", envSize, "+")
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op1.Value)
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op2.Value)
				gen.AddExpression("P", "P", envSize, "+")
				gen.AddCall("comparemenorigString")
				gen.AddGetStack(tmp2, "(int)P")
				gen.AddExpression("P", "P", envSize, "-")
				gen.AddBr()
				//result = environment.NewValue(tmp2, true, environment.BOOLEAN, false, false, false, environment.Variable{})

				tlabel := gen.NewLabel()
				flabel := gen.NewLabel()

				gen.AddIf(tmp2, "1", "==", tlabel)
				gen.AddGoto(flabel)
				result = environment.NewValue(tmp2, false, environment.BOOLEAN, false, false, false, environment.Variable{})
				result.TrueLabel.PushBack(tlabel)
				result.FalseLabel.PushBack(flabel)
				result.Val.FEti = flabel
				result.Val.TEti = tlabel

				return result
			} else if dominante == environment.CHARACTER && (op1.Type == op2.Type) {
				//llamar a generar concatstring
				gen.GenerateCompareMenorigString()
				//concat
				gen.AddComment("Comparando strings")
				envSize := strconv.Itoa(ast.PosicionStack)
				tmp1 := gen.NewTemp()
				tmp2 := gen.NewTemp()
				gen.AddExpression(tmp1, "P", envSize, "+")
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op1.Value)
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op2.Value)
				gen.AddExpression("P", "P", envSize, "+")
				gen.AddCall("comparemenorigString")
				gen.AddGetStack(tmp2, "(int)P")
				gen.AddExpression("P", "P", envSize, "-")
				gen.AddBr()
				//result = environment.NewValue(tmp2, true, environment.BOOLEAN, false, false, false, environment.Variable{})

				tlabel := gen.NewLabel()
				flabel := gen.NewLabel()

				gen.AddIf(tmp2, "1", "==", tlabel)
				gen.AddGoto(flabel)
				result = environment.NewValue(tmp2, false, environment.BOOLEAN, false, false, false, environment.Variable{})
				result.TrueLabel.PushBack(tlabel)
				result.FalseLabel.PushBack(flabel)
				result.Val.FEti = flabel
				result.Val.TEti = tlabel

				return result
			} else {
				Errores := environment.Errores{
					Descripcion: "No es posible comparar los dos valores(<=)",
					Fila:        strconv.Itoa(o.Lin),
					Columna:     strconv.Itoa(o.Col),
					Tipo:        "Error Semantico",
					Ambito:      ast.ObtenerAmbito(),
				}
				ast.ErroresHTML(Errores)
				return result
			}
		}
	case ">=":
		{
			op1 = o.Op_izq.Ejecutar(ast, gen)
			op2 = o.Op_der.Ejecutar(ast, gen)
			//validar tipo dominante
			dominante = tabla_dominante[op1.Type][op2.Type]
			//valida el tipo
			if dominante == environment.INTEGER {
				tlabel := gen.NewLabel()
				flabel := gen.NewLabel()

				newTemp1 := gen.NewTemp()
				newTemp2 := gen.NewTemp()

				if op1.IsTemp && ast.IsTempT(op1.Value) && op1.Val.Name == "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else if !op1.IsTemp && op1.Value != "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else {
					gen.AddGetStack(newTemp1, strconv.Itoa(op1.Val.Symbols.Posicion))
				}

				if op2.IsTemp && ast.IsTempT(op2.Value) && op2.Val.Name == "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else if !op2.IsTemp && op2.Value != "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else {
					gen.AddGetStack(newTemp2, strconv.Itoa(op2.Val.Symbols.Posicion))
				}

				gen.AddIf(newTemp1, newTemp2, ">=", tlabel)
				gen.AddGoto(flabel)
				result = environment.NewValue("", false, environment.BOOLEAN, false, false, false, environment.Variable{})
				result.TrueLabel.PushBack(tlabel)
				result.FalseLabel.PushBack(flabel)
				result.Val.FEti = flabel
				result.Val.TEti = tlabel

				return result
			} else if dominante == environment.FLOAT {
				tlabel := gen.NewLabel()
				flabel := gen.NewLabel()

				newTemp1 := gen.NewTemp()
				newTemp2 := gen.NewTemp()

				if op1.IsTemp && ast.IsTempT(op1.Value) && op1.Val.Name == "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else if !op1.IsTemp && op1.Value != "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else {
					gen.AddGetStack(newTemp1, strconv.Itoa(op1.Val.Symbols.Posicion))
				}

				if op2.IsTemp && ast.IsTempT(op2.Value) && op2.Val.Name == "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else if !op2.IsTemp && op2.Value != "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else {
					gen.AddGetStack(newTemp2, strconv.Itoa(op2.Val.Symbols.Posicion))
				}

				gen.AddIf(newTemp1, newTemp2, ">=", tlabel)
				gen.AddGoto(flabel)
				result = environment.NewValue("", false, environment.BOOLEAN, false, false, false, environment.Variable{})
				result.TrueLabel.PushBack(tlabel)
				result.FalseLabel.PushBack(flabel)
				result.Val.FEti = flabel
				result.Val.TEti = tlabel

				return result
			} else if dominante == environment.STRING && (op1.Type == op2.Type) {
				//llamar a generar concatstring
				gen.GenerateCompareMayorigString()
				//concat
				gen.AddComment("Comparando strings")
				envSize := strconv.Itoa(ast.PosicionStack)
				tmp1 := gen.NewTemp()
				tmp2 := gen.NewTemp()
				gen.AddExpression(tmp1, "P", envSize, "+")
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op1.Value)
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op2.Value)
				gen.AddExpression("P", "P", envSize, "+")
				gen.AddCall("comparemayorigString")
				gen.AddGetStack(tmp2, "(int)P")
				gen.AddExpression("P", "P", envSize, "-")
				gen.AddBr()
				//result = environment.NewValue(tmp2, true, environment.BOOLEAN, false, false, false, environment.Variable{})

				tlabel := gen.NewLabel()
				flabel := gen.NewLabel()

				gen.AddIf(tmp2, "1", "==", tlabel)
				gen.AddGoto(flabel)
				result = environment.NewValue(tmp2, false, environment.BOOLEAN, false, false, false, environment.Variable{})
				result.TrueLabel.PushBack(tlabel)
				result.FalseLabel.PushBack(flabel)
				result.Val.FEti = flabel
				result.Val.TEti = tlabel

				return result
			} else if dominante == environment.CHARACTER && (op1.Type == op2.Type) {
				//llamar a generar concatstring
				gen.GenerateCompareMayorigString()
				//concat
				gen.AddComment("Comparando strings")
				envSize := strconv.Itoa(ast.PosicionStack)
				tmp1 := gen.NewTemp()
				tmp2 := gen.NewTemp()
				gen.AddExpression(tmp1, "P", envSize, "+")
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op1.Value)
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op2.Value)
				gen.AddExpression("P", "P", envSize, "+")
				gen.AddCall("comparemayorigString")
				gen.AddGetStack(tmp2, "(int)P")
				gen.AddExpression("P", "P", envSize, "-")
				gen.AddBr()
				//result = environment.NewValue(tmp2, true, environment.BOOLEAN, false, false, false, environment.Variable{})

				tlabel := gen.NewLabel()
				flabel := gen.NewLabel()

				gen.AddIf(tmp2, "1", "==", tlabel)
				gen.AddGoto(flabel)
				result = environment.NewValue(tmp2, false, environment.BOOLEAN, false, false, false, environment.Variable{})
				result.TrueLabel.PushBack(tlabel)
				result.FalseLabel.PushBack(flabel)
				result.Val.FEti = flabel
				result.Val.TEti = tlabel

				return result
			} else {
				Errores := environment.Errores{
					Descripcion: "No es posible comparar los dos valores(>=)",
					Fila:        strconv.Itoa(o.Lin),
					Columna:     strconv.Itoa(o.Col),
					Tipo:        "Error Semantico",
					Ambito:      ast.ObtenerAmbito(),
				}
				ast.ErroresHTML(Errores)
				return result
			}
		}
	case "==":
		{
			op1 = o.Op_izq.Ejecutar(ast, gen)
			op2 = o.Op_der.Ejecutar(ast, gen)
			//validar tipo dominante
			dominante = tabla_dominante[op1.Type][op2.Type]
			//valida el tipo
			if dominante == environment.INTEGER {
				tlabel := gen.NewLabel()
				flabel := gen.NewLabel()

				newTemp1 := gen.NewTemp()
				newTemp2 := gen.NewTemp()

				if op1.IsTemp && ast.IsTempT(op1.Value) && op1.Val.Name == "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else if !op1.IsTemp && op1.Value != "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else {
					gen.AddGetStack(newTemp1, strconv.Itoa(op1.Val.Symbols.Posicion))
				}

				if op2.IsTemp && ast.IsTempT(op2.Value) && op2.Val.Name == "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else if !op2.IsTemp && op2.Value != "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else {
					gen.AddGetStack(newTemp2, strconv.Itoa(op2.Val.Symbols.Posicion))
				}

				gen.AddIf(newTemp1, newTemp2, "==", tlabel)
				gen.AddGoto(flabel)
				result = environment.NewValue("", false, environment.BOOLEAN, false, false, false, environment.Variable{})
				result.TrueLabel.PushBack(tlabel)
				result.FalseLabel.PushBack(flabel)
				result.Val.FEti = flabel
				result.Val.TEti = tlabel

				return result
			} else if dominante == environment.FLOAT {
				tlabel := gen.NewLabel()
				flabel := gen.NewLabel()

				newTemp1 := gen.NewTemp()
				newTemp2 := gen.NewTemp()

				if op1.IsTemp && ast.IsTempT(op1.Value) && op1.Val.Name == "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else if !op1.IsTemp && op1.Value != "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else {
					gen.AddGetStack(newTemp1, strconv.Itoa(op1.Val.Symbols.Posicion))
				}

				if op2.IsTemp && ast.IsTempT(op2.Value) && op2.Val.Name == "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else if !op2.IsTemp && op2.Value != "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else {
					gen.AddGetStack(newTemp2, strconv.Itoa(op2.Val.Symbols.Posicion))
				}

				gen.AddIf(newTemp1, newTemp2, "==", tlabel)
				gen.AddGoto(flabel)
				result = environment.NewValue("", false, environment.BOOLEAN, false, false, false, environment.Variable{})
				result.TrueLabel.PushBack(tlabel)
				result.FalseLabel.PushBack(flabel)
				result.Val.FEti = flabel
				result.Val.TEti = tlabel

				return result
			} else if dominante == environment.STRING && (op1.Type == op2.Type) {
				//llamar a generar concatstring
				gen.GenerateCompareIgualString()
				//concat
				gen.AddComment("Comparando strings")
				envSize := strconv.Itoa(ast.PosicionStack)
				tmp1 := gen.NewTemp()
				tmp2 := gen.NewTemp()
				gen.AddExpression(tmp1, "P", envSize, "+")
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op1.Value)
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op2.Value)
				gen.AddExpression("P", "P", envSize, "+")
				gen.AddCall("compareigualString")
				gen.AddGetStack(tmp2, "(int)P")
				gen.AddExpression("P", "P", envSize, "-")
				gen.AddBr()
				//result = environment.NewValue(tmp2, true, environment.BOOLEAN, false, false, false, environment.Variable{})

				tlabel := gen.NewLabel()
				flabel := gen.NewLabel()

				gen.AddIf(tmp2, "1", "==", tlabel)
				gen.AddGoto(flabel)
				result = environment.NewValue(tmp2, false, environment.BOOLEAN, false, false, false, environment.Variable{})
				result.TrueLabel.PushBack(tlabel)
				result.FalseLabel.PushBack(flabel)
				result.Val.FEti = flabel
				result.Val.TEti = tlabel

				return result
			} else if dominante == environment.CHARACTER && (op1.Type == op2.Type) {
				//llamar a generar concatstring
				gen.GenerateCompareIgualString()
				//concat
				gen.AddComment("Comparando strings")
				envSize := strconv.Itoa(ast.PosicionStack)
				tmp1 := gen.NewTemp()
				tmp2 := gen.NewTemp()
				gen.AddExpression(tmp1, "P", envSize, "+")
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op1.Value)
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op2.Value)
				gen.AddExpression("P", "P", envSize, "+")
				gen.AddCall("compareigualString")
				gen.AddGetStack(tmp2, "(int)P")
				gen.AddExpression("P", "P", envSize, "-")
				gen.AddBr()
				//result = environment.NewValue(tmp2, true, environment.BOOLEAN, false, false, false, environment.Variable{})

				tlabel := gen.NewLabel()
				flabel := gen.NewLabel()

				gen.AddIf(tmp2, "1", "==", tlabel)
				gen.AddGoto(flabel)
				result = environment.NewValue(tmp2, false, environment.BOOLEAN, false, false, false, environment.Variable{})
				result.TrueLabel.PushBack(tlabel)
				result.FalseLabel.PushBack(flabel)
				result.Val.FEti = flabel
				result.Val.TEti = tlabel

				return result
			} else if dominante == environment.BOOLEAN && (op1.Type == op2.Type) {
				tlabel := gen.NewLabel()
				flabel := gen.NewLabel()
				newTemp1 := gen.NewTemp()
				newTemp3 := gen.NewTemp()
				newTemp4 := gen.NewTemp()
				newTemp6 := gen.NewTemp()

				if op1.Val.Name == "" {
					gen.AddExpression(newTemp1, op1.Value, "", "")
					gen.AddSetStack(strconv.Itoa(ast.PosicionStack+1), newTemp1)
					gen.AddGetStack(newTemp3, strconv.Itoa(ast.PosicionStack+1))
				} else {
					gen.AddGetStack(newTemp3, strconv.Itoa(op1.Val.Symbols.Posicion))
				}

				if op2.Val.Name == "" {
					gen.AddExpression(newTemp4, op2.Value, "", "")
					gen.AddSetStack(strconv.Itoa(ast.PosicionStack+1), newTemp4)
					gen.AddGetStack(newTemp6, strconv.Itoa(ast.PosicionStack+1))
				} else {
					gen.AddGetStack(newTemp6, strconv.Itoa(op2.Val.Symbols.Posicion))
				}
				gen.AddIf(newTemp3, newTemp6, "==", tlabel)
				gen.AddGoto(flabel)
				result = environment.NewValue("", false, environment.BOOLEAN, false, false, false, environment.Variable{})
				result.TrueLabel.PushBack(tlabel)
				result.FalseLabel.PushBack(flabel)
				result.Val.FEti = flabel
				result.Val.TEti = tlabel

				return result
			} else if dominante == environment.NULL && op1.Type != environment.NULL {
				if op1.Type == environment.INTEGER {
					tlabel := gen.NewLabel()
					flabel := gen.NewLabel()

					gen.AddIf(op1.Value, op2.Value, "==", tlabel)
					gen.AddGoto(flabel)
					result = environment.NewValue("", false, environment.BOOLEAN, false, false, false, environment.Variable{})
					result.TrueLabel.PushBack(tlabel)
					result.FalseLabel.PushBack(flabel)
					result.Val.FEti = flabel
					result.Val.TEti = tlabel

					return result
				} else if op1.Type == environment.FLOAT {
					tlabel := gen.NewLabel()
					flabel := gen.NewLabel()

					gen.AddIf(op1.Value, op2.Value, "==", tlabel)
					gen.AddGoto(flabel)
					result = environment.NewValue("", false, environment.BOOLEAN, false, false, false, environment.Variable{})
					result.TrueLabel.PushBack(tlabel)
					result.FalseLabel.PushBack(flabel)
					result.Val.FEti = flabel
					result.Val.TEti = tlabel

					return result
				} else if op1.Type == environment.STRING {
					//llamar a generar concatstring
					gen.GenerateCompareIgualString()
					//concat
					gen.AddComment("Comparando strings")
					envSize := strconv.Itoa(ast.PosicionStack)
					tmp1 := gen.NewTemp()
					tmp2 := gen.NewTemp()
					gen.AddExpression(tmp1, "P", envSize, "+")
					gen.AddExpression(tmp1, tmp1, "1", "+")
					temp3 := gen.NewTemp()
					gen.AddGetStack(temp3, strconv.Itoa(op1.Val.Symbols.Posicion))
					gen.AddSetStack("(int)"+tmp1, temp3)
					gen.AddExpression(tmp1, tmp1, "1", "+")
					temp4 := gen.NewTemp()                //Cremos un nuevo temporal en esta va a ir el Head osea la asignacion de temp = head
					gen.AddAssign(temp4, "H")             //Creamos un nuevo puntero Head para que en este vayan los valores
					gen.AddSetHeap("(int)H", "-1")        // Agregamos el caracter de escape
					gen.AddExpression("H", "H", "1", "+") // agregamos una nueva expresion igual que arriba de addexpression
					gen.AddSetStack("(int)"+tmp1, temp4)
					gen.AddExpression("P", "P", envSize, "+")
					gen.AddCall("compareigualString")
					gen.AddGetStack(tmp2, "(int)P")
					gen.AddExpression("P", "P", envSize, "-")
					gen.AddBr()
					//result = environment.NewValue(tmp2, true, environment.BOOLEAN, false, false, false, environment.Variable{})

					tlabel := gen.NewLabel()
					flabel := gen.NewLabel()

					gen.AddIf(tmp2, "1", "==", tlabel)
					gen.AddGoto(flabel)
					result = environment.NewValue(tmp2, false, environment.BOOLEAN, false, false, false, environment.Variable{})
					result.TrueLabel.PushBack(tlabel)
					result.FalseLabel.PushBack(flabel)
					result.Val.FEti = flabel
					result.Val.TEti = tlabel

					return result
				} else if op1.Type == environment.CHARACTER {
					//llamar a generar concatstring
					gen.GenerateCompareIgualString()
					//concat
					gen.AddComment("Comparando strings")
					envSize := strconv.Itoa(ast.PosicionStack)
					tmp1 := gen.NewTemp()
					tmp2 := gen.NewTemp()
					gen.AddExpression(tmp1, "P", envSize, "+")
					gen.AddExpression(tmp1, tmp1, "1", "+")
					temp3 := gen.NewTemp()
					gen.AddGetStack(temp3, strconv.Itoa(op1.Val.Symbols.Posicion))
					gen.AddSetStack("(int)"+tmp1, temp3)
					gen.AddExpression(tmp1, tmp1, "1", "+")
					temp4 := gen.NewTemp()                //Cremos un nuevo temporal en esta va a ir el Head osea la asignacion de temp = head
					gen.AddAssign(temp4, "H")             //Creamos un nuevo puntero Head para que en este vayan los valores
					gen.AddSetHeap("(int)H", "-1")        // Agregamos el caracter de escape
					gen.AddExpression("H", "H", "1", "+") // agregamos una nueva expresion igual que arriba de addexpression
					gen.AddSetStack("(int)"+tmp1, temp4)
					gen.AddExpression("P", "P", envSize, "+")
					gen.AddCall("compareigualString")
					gen.AddGetStack(tmp2, "(int)P")
					gen.AddExpression("P", "P", envSize, "-")
					gen.AddBr()
					//result = environment.NewValue(tmp2, true, environment.BOOLEAN, false, false, false, environment.Variable{})

					tlabel := gen.NewLabel()
					flabel := gen.NewLabel()

					gen.AddIf(tmp2, "1", "==", tlabel)
					gen.AddGoto(flabel)
					result = environment.NewValue(tmp2, false, environment.BOOLEAN, false, false, false, environment.Variable{})
					result.TrueLabel.PushBack(tlabel)
					result.FalseLabel.PushBack(flabel)
					result.Val.FEti = flabel
					result.Val.TEti = tlabel

					return result
				} else if op1.Type == environment.BOOLEAN {
					tlabel := gen.NewLabel()
					flabel := gen.NewLabel()
					newTemp1 := gen.NewTemp()
					newTemp3 := gen.NewTemp()
					newTemp4 := gen.NewTemp()
					newTemp6 := gen.NewTemp()

					if op1.Val.Name == "" {
						gen.AddExpression(newTemp1, op1.Value, "", "")
						gen.AddSetStack(strconv.Itoa(ast.PosicionStack+1), newTemp1)
						gen.AddGetStack(newTemp3, strconv.Itoa(ast.PosicionStack+1))
					} else {
						gen.AddGetStack(newTemp3, strconv.Itoa(op1.Val.Symbols.Posicion))
					}

					if op2.Val.Name == "" {
						gen.AddExpression(newTemp4, "0", "", "")
						gen.AddSetStack(strconv.Itoa(ast.PosicionStack+1), newTemp4)
						gen.AddGetStack(newTemp6, strconv.Itoa(ast.PosicionStack+1))
					} else {
						gen.AddGetStack(newTemp6, strconv.Itoa(op2.Val.Symbols.Posicion))
					}
					gen.AddIf(newTemp3, newTemp6, "==", tlabel)
					gen.AddGoto(flabel)
					result = environment.NewValue("", false, environment.BOOLEAN, false, false, false, environment.Variable{})
					result.TrueLabel.PushBack(tlabel)
					result.FalseLabel.PushBack(flabel)
					result.Val.FEti = flabel
					result.Val.TEti = tlabel

					return result
				}

			} else if dominante == environment.NULL && op2.Type != environment.NULL {
				if op2.Type == environment.INTEGER {
					tlabel := gen.NewLabel()
					flabel := gen.NewLabel()

					gen.AddIf(op1.Value, op2.Value, "==", tlabel)
					gen.AddGoto(flabel)
					result = environment.NewValue("", false, environment.BOOLEAN, false, false, false, environment.Variable{})
					result.TrueLabel.PushBack(tlabel)
					result.FalseLabel.PushBack(flabel)
					result.Val.FEti = flabel
					result.Val.TEti = tlabel

					return result
				} else if op2.Type == environment.FLOAT {
					tlabel := gen.NewLabel()
					flabel := gen.NewLabel()

					gen.AddIf(op1.Value, op2.Value, "==", tlabel)
					gen.AddGoto(flabel)
					result = environment.NewValue("", false, environment.BOOLEAN, false, false, false, environment.Variable{})
					result.TrueLabel.PushBack(tlabel)
					result.FalseLabel.PushBack(flabel)
					result.Val.FEti = flabel
					result.Val.TEti = tlabel

					return result
				} else if op2.Type == environment.STRING {
					//llamar a generar concatstring
					gen.GenerateCompareIgualString()
					//concat
					gen.AddComment("Comparando strings")
					envSize := strconv.Itoa(ast.PosicionStack)
					tmp1 := gen.NewTemp()
					tmp2 := gen.NewTemp()
					gen.AddExpression(tmp1, "P", envSize, "+")
					gen.AddExpression(tmp1, tmp1, "1", "+")
					temp3 := gen.NewTemp()
					gen.AddGetStack(temp3, strconv.Itoa(op2.Val.Symbols.Posicion))
					gen.AddSetStack("(int)"+tmp1, temp3)
					gen.AddExpression(tmp1, tmp1, "1", "+")
					temp4 := gen.NewTemp()                //Cremos un nuevo temporal en esta va a ir el Head osea la asignacion de temp = head
					gen.AddAssign(temp4, "H")             //Creamos un nuevo puntero Head para que en este vayan los valores
					gen.AddSetHeap("(int)H", "-1")        // Agregamos el caracter de escape
					gen.AddExpression("H", "H", "1", "+") // agregamos una nueva expresion igual que arriba de addexpression
					gen.AddSetStack("(int)"+tmp1, temp4)
					gen.AddExpression("P", "P", envSize, "+")
					gen.AddCall("compareigualString")
					gen.AddGetStack(tmp2, "(int)P")
					gen.AddExpression("P", "P", envSize, "-")
					gen.AddBr()
					//result = environment.NewValue(tmp2, true, environment.BOOLEAN, false, false, false, environment.Variable{})

					tlabel := gen.NewLabel()
					flabel := gen.NewLabel()

					gen.AddIf(tmp2, "1", "==", tlabel)
					gen.AddGoto(flabel)
					result = environment.NewValue(tmp2, false, environment.BOOLEAN, false, false, false, environment.Variable{})
					result.TrueLabel.PushBack(tlabel)
					result.FalseLabel.PushBack(flabel)
					result.Val.FEti = flabel
					result.Val.TEti = tlabel

					return result
				} else if op2.Type == environment.CHARACTER {
					//llamar a generar concatstring
					gen.GenerateCompareIgualString()
					//concat
					gen.AddComment("Comparando strings")
					envSize := strconv.Itoa(ast.PosicionStack)
					tmp1 := gen.NewTemp()
					tmp2 := gen.NewTemp()
					gen.AddExpression(tmp1, "P", envSize, "+")
					gen.AddExpression(tmp1, tmp1, "1", "+")
					temp3 := gen.NewTemp()
					gen.AddGetStack(temp3, strconv.Itoa(op2.Val.Symbols.Posicion))
					gen.AddSetStack("(int)"+tmp1, temp3)
					gen.AddExpression(tmp1, tmp1, "1", "+")
					temp4 := gen.NewTemp()                //Cremos un nuevo temporal en esta va a ir el Head osea la asignacion de temp = head
					gen.AddAssign(temp4, "H")             //Creamos un nuevo puntero Head para que en este vayan los valores
					gen.AddSetHeap("(int)H", "-1")        // Agregamos el caracter de escape
					gen.AddExpression("H", "H", "1", "+") // agregamos una nueva expresion igual que arriba de addexpression
					gen.AddSetStack("(int)"+tmp1, temp4)
					gen.AddExpression("P", "P", envSize, "+")
					gen.AddCall("compareigualString")
					gen.AddGetStack(tmp2, "(int)P")
					gen.AddExpression("P", "P", envSize, "-")
					gen.AddBr()
					//result = environment.NewValue(tmp2, true, environment.BOOLEAN, false, false, false, environment.Variable{})

					tlabel := gen.NewLabel()
					flabel := gen.NewLabel()

					gen.AddIf(tmp2, "1", "==", tlabel)
					gen.AddGoto(flabel)
					result = environment.NewValue(tmp2, false, environment.BOOLEAN, false, false, false, environment.Variable{})
					result.TrueLabel.PushBack(tlabel)
					result.FalseLabel.PushBack(flabel)
					result.Val.FEti = flabel
					result.Val.TEti = tlabel

					return result
				} else if op2.Type == environment.BOOLEAN {
					tlabel := gen.NewLabel()
					flabel := gen.NewLabel()
					newTemp1 := gen.NewTemp()
					newTemp3 := gen.NewTemp()
					newTemp4 := gen.NewTemp()
					newTemp6 := gen.NewTemp()

					if op1.Val.Name == "" {
						gen.AddExpression(newTemp1, "0", "", "")
						gen.AddSetStack(strconv.Itoa(ast.PosicionStack+1), newTemp1)
						gen.AddGetStack(newTemp3, strconv.Itoa(ast.PosicionStack+1))
					} else {
						gen.AddGetStack(newTemp3, strconv.Itoa(op1.Val.Symbols.Posicion))
					}

					if op2.Val.Name == "" {
						gen.AddExpression(newTemp4, op2.Value, "", "")
						gen.AddSetStack(strconv.Itoa(ast.PosicionStack+1), newTemp4)
						gen.AddGetStack(newTemp6, strconv.Itoa(ast.PosicionStack+1))
					} else {
						gen.AddGetStack(newTemp6, strconv.Itoa(op2.Val.Symbols.Posicion))
					}
					gen.AddIf(newTemp3, newTemp6, "==", tlabel)
					gen.AddGoto(flabel)
					result = environment.NewValue("", false, environment.BOOLEAN, false, false, false, environment.Variable{})
					result.TrueLabel.PushBack(tlabel)
					result.FalseLabel.PushBack(flabel)
					result.Val.FEti = flabel
					result.Val.TEti = tlabel

					return result
				}

			} else {
				Errores := environment.Errores{
					Descripcion: "No es posible comparar los dos valores(==)",
					Fila:        strconv.Itoa(o.Lin),
					Columna:     strconv.Itoa(o.Col),
					Tipo:        "Error Semantico",
					Ambito:      ast.ObtenerAmbito(),
				}
				ast.ErroresHTML(Errores)
				return result
			}
		}
	case "!=":
		{
			op1 = o.Op_izq.Ejecutar(ast, gen)
			op2 = o.Op_der.Ejecutar(ast, gen)
			//validar tipo dominante
			dominante = tabla_dominante[op1.Type][op2.Type]
			//valida el tipo
			if dominante == environment.INTEGER {
				tlabel := gen.NewLabel()
				flabel := gen.NewLabel()

				newTemp1 := gen.NewTemp()
				newTemp2 := gen.NewTemp()

				if op1.IsTemp && ast.IsTempT(op1.Value) && op1.Val.Name == "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else if !op1.IsTemp && op1.Value != "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else {
					gen.AddGetStack(newTemp1, strconv.Itoa(op1.Val.Symbols.Posicion))
				}

				if op2.IsTemp && ast.IsTempT(op2.Value) && op2.Val.Name == "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else if !op2.IsTemp && op2.Value != "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else {
					gen.AddGetStack(newTemp2, strconv.Itoa(op2.Val.Symbols.Posicion))
				}

				gen.AddIf(newTemp1, newTemp2, "!=", tlabel)
				gen.AddGoto(flabel)
				result = environment.NewValue("", false, environment.BOOLEAN, false, false, false, environment.Variable{})
				result.TrueLabel.PushBack(tlabel)
				result.FalseLabel.PushBack(flabel)
				result.Val.FEti = flabel
				result.Val.TEti = tlabel

				return result
			} else if dominante == environment.FLOAT {
				tlabel := gen.NewLabel()
				flabel := gen.NewLabel()

				newTemp1 := gen.NewTemp()
				newTemp2 := gen.NewTemp()

				if op1.IsTemp && ast.IsTempT(op1.Value) && op1.Val.Name == "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else if !op1.IsTemp && op1.Value != "" {
					gen.AddAssign(newTemp1, op1.Value)
				} else {
					gen.AddGetStack(newTemp1, strconv.Itoa(op1.Val.Symbols.Posicion))
				}

				if op2.IsTemp && ast.IsTempT(op2.Value) && op2.Val.Name == "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else if !op2.IsTemp && op2.Value != "" {
					gen.AddAssign(newTemp2, op2.Value)
				} else {
					gen.AddGetStack(newTemp2, strconv.Itoa(op2.Val.Symbols.Posicion))
				}

				gen.AddIf(newTemp1, newTemp2, "!=", tlabel)
				gen.AddGoto(flabel)
				result = environment.NewValue("", false, environment.BOOLEAN, false, false, false, environment.Variable{})
				result.TrueLabel.PushBack(tlabel)
				result.FalseLabel.PushBack(flabel)
				result.Val.FEti = flabel
				result.Val.TEti = tlabel

				return result
			} else if dominante == environment.STRING && (op1.Type == op2.Type) {
				//llamar a generar concatstring
				gen.GenerateCompareDifeString()
				//concat
				gen.AddComment("Comparando strings")
				envSize := strconv.Itoa(ast.PosicionStack)
				tmp1 := gen.NewTemp()
				tmp2 := gen.NewTemp()
				gen.AddExpression(tmp1, "P", envSize, "+")
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op1.Value)
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op2.Value)
				gen.AddExpression("P", "P", envSize, "+")
				gen.AddCall("comparedifeString")
				gen.AddGetStack(tmp2, "(int)P")
				gen.AddExpression("P", "P", envSize, "-")
				gen.AddBr()
				//result = environment.NewValue(tmp2, true, environment.BOOLEAN, false, false, false, environment.Variable{})

				tlabel := gen.NewLabel()
				flabel := gen.NewLabel()

				gen.AddIf(tmp2, "1", "==", tlabel)
				gen.AddGoto(flabel)
				result = environment.NewValue(tmp2, false, environment.BOOLEAN, false, false, false, environment.Variable{})
				result.TrueLabel.PushBack(tlabel)
				result.FalseLabel.PushBack(flabel)
				result.Val.FEti = flabel
				result.Val.TEti = tlabel

				return result
			} else if dominante == environment.CHARACTER && (op1.Type == op2.Type) {
				//llamar a generar concatstring
				gen.GenerateCompareDifeString()
				//concat
				gen.AddComment("Comparando strings")
				envSize := strconv.Itoa(ast.PosicionStack)
				tmp1 := gen.NewTemp()
				tmp2 := gen.NewTemp()
				gen.AddExpression(tmp1, "P", envSize, "+")
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op1.Value)
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op2.Value)
				gen.AddExpression("P", "P", envSize, "+")
				gen.AddCall("comparedifeString")
				gen.AddGetStack(tmp2, "(int)P")
				gen.AddExpression("P", "P", envSize, "-")
				gen.AddBr()
				//result = environment.NewValue(tmp2, true, environment.BOOLEAN, false, false, false, environment.Variable{})

				tlabel := gen.NewLabel()
				flabel := gen.NewLabel()

				gen.AddIf(tmp2, "1", "==", tlabel)
				gen.AddGoto(flabel)
				result = environment.NewValue(tmp2, false, environment.BOOLEAN, false, false, false, environment.Variable{})
				result.TrueLabel.PushBack(tlabel)
				result.FalseLabel.PushBack(flabel)
				result.Val.FEti = flabel
				result.Val.TEti = tlabel

				return result
			} else if dominante == environment.BOOLEAN && (op1.Type == op2.Type) {
				tlabel := gen.NewLabel()
				flabel := gen.NewLabel()
				newTemp1 := gen.NewTemp()
				newTemp3 := gen.NewTemp()
				newTemp4 := gen.NewTemp()
				newTemp6 := gen.NewTemp()

				if op1.Val.Name == "" {
					gen.AddExpression(newTemp1, op1.Value, "", "")
					gen.AddSetStack(strconv.Itoa(ast.PosicionStack+1), newTemp1)
					gen.AddGetStack(newTemp3, strconv.Itoa(ast.PosicionStack+1))
				} else {
					gen.AddGetStack(newTemp3, strconv.Itoa(op1.Val.Symbols.Posicion))
				}

				if op2.Val.Name == "" {
					gen.AddExpression(newTemp4, op2.Value, "", "")
					gen.AddSetStack(strconv.Itoa(ast.PosicionStack+1), newTemp4)
					gen.AddGetStack(newTemp6, strconv.Itoa(ast.PosicionStack+1))
				} else {
					gen.AddGetStack(newTemp6, strconv.Itoa(op2.Val.Symbols.Posicion))
				}
				gen.AddIf(newTemp3, newTemp6, "!=", tlabel)
				gen.AddGoto(flabel)
				result = environment.NewValue("", false, environment.BOOLEAN, false, false, false, environment.Variable{})
				result.TrueLabel.PushBack(tlabel)
				result.FalseLabel.PushBack(flabel)
				result.Val.FEti = flabel
				result.Val.TEti = tlabel

				return result
			} else if dominante == environment.NULL && op1.Type != environment.NULL {
				if op1.Type == environment.INTEGER {
					tlabel := gen.NewLabel()
					flabel := gen.NewLabel()

					gen.AddIf(op1.Value, op2.Value, "!=", tlabel)
					gen.AddGoto(flabel)
					result = environment.NewValue("", false, environment.BOOLEAN, false, false, false, environment.Variable{})
					result.TrueLabel.PushBack(tlabel)
					result.FalseLabel.PushBack(flabel)
					result.Val.FEti = flabel
					result.Val.TEti = tlabel

					return result
				} else if op1.Type == environment.FLOAT {
					tlabel := gen.NewLabel()
					flabel := gen.NewLabel()

					gen.AddIf(op1.Value, op2.Value, "!=", tlabel)
					gen.AddGoto(flabel)
					result = environment.NewValue("", false, environment.BOOLEAN, false, false, false, environment.Variable{})
					result.TrueLabel.PushBack(tlabel)
					result.FalseLabel.PushBack(flabel)
					result.Val.FEti = flabel
					result.Val.TEti = tlabel

					return result
				} else if op1.Type == environment.STRING {
					//llamar a generar concatstring
					gen.GenerateCompareDifeString()
					//concat
					gen.AddComment("Comparando strings")
					envSize := strconv.Itoa(ast.PosicionStack)
					tmp1 := gen.NewTemp()
					tmp2 := gen.NewTemp()
					gen.AddExpression(tmp1, "P", envSize, "+")
					gen.AddExpression(tmp1, tmp1, "1", "+")
					temp3 := gen.NewTemp()
					gen.AddGetStack(temp3, strconv.Itoa(op1.Val.Symbols.Posicion))
					gen.AddSetStack("(int)"+tmp1, temp3)
					gen.AddExpression(tmp1, tmp1, "1", "+")
					temp4 := gen.NewTemp()                //Cremos un nuevo temporal en esta va a ir el Head osea la asignacion de temp = head
					gen.AddAssign(temp4, "H")             //Creamos un nuevo puntero Head para que en este vayan los valores
					gen.AddSetHeap("(int)H", "-1")        // Agregamos el caracter de escape
					gen.AddExpression("H", "H", "1", "+") // agregamos una nueva expresion igual que arriba de addexpression
					gen.AddSetStack("(int)"+tmp1, temp4)
					gen.AddExpression("P", "P", envSize, "+")
					gen.AddCall("comparedifeString")
					gen.AddGetStack(tmp2, "(int)P")
					gen.AddExpression("P", "P", envSize, "-")
					gen.AddBr()
					//result = environment.NewValue(tmp2, true, environment.BOOLEAN, false, false, false, environment.Variable{})

					tlabel := gen.NewLabel()
					flabel := gen.NewLabel()

					gen.AddIf(tmp2, "1", "==", tlabel)
					gen.AddGoto(flabel)
					result = environment.NewValue(tmp2, false, environment.BOOLEAN, false, false, false, environment.Variable{})
					result.TrueLabel.PushBack(tlabel)
					result.FalseLabel.PushBack(flabel)
					result.Val.FEti = flabel
					result.Val.TEti = tlabel

					return result
				} else if op1.Type == environment.CHARACTER {
					//llamar a generar concatstring
					gen.GenerateCompareDifeString()
					//concat
					gen.AddComment("Comparando strings")
					envSize := strconv.Itoa(ast.PosicionStack)
					tmp1 := gen.NewTemp()
					tmp2 := gen.NewTemp()
					gen.AddExpression(tmp1, "P", envSize, "+")
					gen.AddExpression(tmp1, tmp1, "1", "+")
					temp3 := gen.NewTemp()
					gen.AddGetStack(temp3, strconv.Itoa(op1.Val.Symbols.Posicion))
					gen.AddSetStack("(int)"+tmp1, temp3)
					gen.AddExpression(tmp1, tmp1, "1", "+")
					temp4 := gen.NewTemp()                //Cremos un nuevo temporal en esta va a ir el Head osea la asignacion de temp = head
					gen.AddAssign(temp4, "H")             //Creamos un nuevo puntero Head para que en este vayan los valores
					gen.AddSetHeap("(int)H", "-1")        // Agregamos el caracter de escape
					gen.AddExpression("H", "H", "1", "+") // agregamos una nueva expresion igual que arriba de addexpression
					gen.AddSetStack("(int)"+tmp1, temp4)
					gen.AddExpression("P", "P", envSize, "+")
					gen.AddCall("comparedifeString")
					gen.AddGetStack(tmp2, "(int)P")
					gen.AddExpression("P", "P", envSize, "-")
					gen.AddBr()
					//result = environment.NewValue(tmp2, true, environment.BOOLEAN, false, false, false, environment.Variable{})

					tlabel := gen.NewLabel()
					flabel := gen.NewLabel()

					gen.AddIf(tmp2, "1", "==", tlabel)
					gen.AddGoto(flabel)
					result = environment.NewValue(tmp2, false, environment.BOOLEAN, false, false, false, environment.Variable{})
					result.TrueLabel.PushBack(tlabel)
					result.FalseLabel.PushBack(flabel)
					result.Val.FEti = flabel
					result.Val.TEti = tlabel

					return result
				} else if op1.Type == environment.BOOLEAN {
					tlabel := gen.NewLabel()
					flabel := gen.NewLabel()
					newTemp1 := gen.NewTemp()
					newTemp3 := gen.NewTemp()
					newTemp4 := gen.NewTemp()
					newTemp6 := gen.NewTemp()

					if op1.Val.Name == "" {
						gen.AddExpression(newTemp1, op1.Value, "", "")
						gen.AddSetStack(strconv.Itoa(ast.PosicionStack+1), newTemp1)
						gen.AddGetStack(newTemp3, strconv.Itoa(ast.PosicionStack+1))
					} else {
						gen.AddGetStack(newTemp3, strconv.Itoa(op1.Val.Symbols.Posicion))
					}

					if op2.Val.Name == "" {
						gen.AddExpression(newTemp4, "0", "", "")
						gen.AddSetStack(strconv.Itoa(ast.PosicionStack+1), newTemp4)
						gen.AddGetStack(newTemp6, strconv.Itoa(ast.PosicionStack+1))
					} else {
						gen.AddGetStack(newTemp6, strconv.Itoa(op2.Val.Symbols.Posicion))
					}
					gen.AddIf(newTemp3, newTemp6, "!=", tlabel)
					gen.AddGoto(flabel)
					result = environment.NewValue("", false, environment.BOOLEAN, false, false, false, environment.Variable{})
					result.TrueLabel.PushBack(tlabel)
					result.FalseLabel.PushBack(flabel)
					result.Val.FEti = flabel
					result.Val.TEti = tlabel

					return result
				}

			} else if dominante == environment.NULL && op2.Type != environment.NULL {
				if op2.Type == environment.INTEGER {
					tlabel := gen.NewLabel()
					flabel := gen.NewLabel()

					gen.AddIf(op1.Value, op2.Value, "!=", tlabel)
					gen.AddGoto(flabel)
					result = environment.NewValue("", false, environment.BOOLEAN, false, false, false, environment.Variable{})
					result.TrueLabel.PushBack(tlabel)
					result.FalseLabel.PushBack(flabel)
					result.Val.FEti = flabel
					result.Val.TEti = tlabel

					return result
				} else if op2.Type == environment.FLOAT {
					tlabel := gen.NewLabel()
					flabel := gen.NewLabel()

					gen.AddIf(op1.Value, op2.Value, "!=", tlabel)
					gen.AddGoto(flabel)
					result = environment.NewValue("", false, environment.BOOLEAN, false, false, false, environment.Variable{})
					result.TrueLabel.PushBack(tlabel)
					result.FalseLabel.PushBack(flabel)
					result.Val.FEti = flabel
					result.Val.TEti = tlabel

					return result
				} else if op2.Type == environment.STRING {
					//llamar a generar concatstring
					gen.GenerateCompareDifeString()
					//concat
					gen.AddComment("Comparando strings")
					envSize := strconv.Itoa(ast.PosicionStack)
					tmp1 := gen.NewTemp()
					tmp2 := gen.NewTemp()
					gen.AddExpression(tmp1, "P", envSize, "+")
					gen.AddExpression(tmp1, tmp1, "1", "+")
					temp3 := gen.NewTemp()
					gen.AddGetStack(temp3, strconv.Itoa(op2.Val.Symbols.Posicion))
					gen.AddSetStack("(int)"+tmp1, temp3)
					gen.AddExpression(tmp1, tmp1, "1", "+")
					temp4 := gen.NewTemp()                //Cremos un nuevo temporal en esta va a ir el Head osea la asignacion de temp = head
					gen.AddAssign(temp4, "H")             //Creamos un nuevo puntero Head para que en este vayan los valores
					gen.AddSetHeap("(int)H", "-1")        // Agregamos el caracter de escape
					gen.AddExpression("H", "H", "1", "+") // agregamos una nueva expresion igual que arriba de addexpression
					gen.AddSetStack("(int)"+tmp1, temp4)
					gen.AddExpression("P", "P", envSize, "+")
					gen.AddCall("comparedifeString")
					gen.AddGetStack(tmp2, "(int)P")
					gen.AddExpression("P", "P", envSize, "-")
					gen.AddBr()
					//result = environment.NewValue(tmp2, true, environment.BOOLEAN, false, false, false, environment.Variable{})

					tlabel := gen.NewLabel()
					flabel := gen.NewLabel()

					gen.AddIf(tmp2, "1", "!=", tlabel)
					gen.AddGoto(flabel)
					result = environment.NewValue(tmp2, false, environment.BOOLEAN, false, false, false, environment.Variable{})
					result.TrueLabel.PushBack(tlabel)
					result.FalseLabel.PushBack(flabel)
					result.Val.FEti = flabel
					result.Val.TEti = tlabel

					return result
				} else if op2.Type == environment.CHARACTER {
					//llamar a generar concatstring
					gen.GenerateCompareDifeString()
					//concat
					gen.AddComment("Comparando strings")
					envSize := strconv.Itoa(ast.PosicionStack)
					tmp1 := gen.NewTemp()
					tmp2 := gen.NewTemp()
					gen.AddExpression(tmp1, "P", envSize, "+")
					gen.AddExpression(tmp1, tmp1, "1", "+")
					temp3 := gen.NewTemp()
					gen.AddGetStack(temp3, strconv.Itoa(op2.Val.Symbols.Posicion))
					gen.AddSetStack("(int)"+tmp1, temp3)
					gen.AddExpression(tmp1, tmp1, "1", "+")
					temp4 := gen.NewTemp()                //Cremos un nuevo temporal en esta va a ir el Head osea la asignacion de temp = head
					gen.AddAssign(temp4, "H")             //Creamos un nuevo puntero Head para que en este vayan los valores
					gen.AddSetHeap("(int)H", "-1")        // Agregamos el caracter de escape
					gen.AddExpression("H", "H", "1", "+") // agregamos una nueva expresion igual que arriba de addexpression
					gen.AddSetStack("(int)"+tmp1, temp4)
					gen.AddExpression("P", "P", envSize, "+")
					gen.AddCall("comparedifeString")
					gen.AddGetStack(tmp2, "(int)P")
					gen.AddExpression("P", "P", envSize, "-")
					gen.AddBr()
					//result = environment.NewValue(tmp2, true, environment.BOOLEAN, false, false, false, environment.Variable{})

					tlabel := gen.NewLabel()
					flabel := gen.NewLabel()

					gen.AddIf(tmp2, "1", "!=", tlabel)
					gen.AddGoto(flabel)
					result = environment.NewValue(tmp2, false, environment.BOOLEAN, false, false, false, environment.Variable{})
					result.TrueLabel.PushBack(tlabel)
					result.FalseLabel.PushBack(flabel)
					result.Val.FEti = flabel
					result.Val.TEti = tlabel

					return result
				} else if op2.Type == environment.BOOLEAN {
					tlabel := gen.NewLabel()
					flabel := gen.NewLabel()
					newTemp1 := gen.NewTemp()
					newTemp3 := gen.NewTemp()
					newTemp4 := gen.NewTemp()
					newTemp6 := gen.NewTemp()

					if op1.Val.Name == "" {
						gen.AddExpression(newTemp1, "0", "", "")
						gen.AddSetStack(strconv.Itoa(ast.PosicionStack+1), newTemp1)
						gen.AddGetStack(newTemp3, strconv.Itoa(ast.PosicionStack+1))
					} else {
						gen.AddGetStack(newTemp3, strconv.Itoa(op1.Val.Symbols.Posicion))
					}

					if op2.Val.Name == "" {
						gen.AddExpression(newTemp4, op2.Value, "", "")
						gen.AddSetStack(strconv.Itoa(ast.PosicionStack+1), newTemp4)
						gen.AddGetStack(newTemp6, strconv.Itoa(ast.PosicionStack+1))
					} else {
						gen.AddGetStack(newTemp6, strconv.Itoa(op2.Val.Symbols.Posicion))
					}
					gen.AddIf(newTemp3, newTemp6, "!=", tlabel)
					gen.AddGoto(flabel)
					result = environment.NewValue("", false, environment.BOOLEAN, false, false, false, environment.Variable{})
					result.TrueLabel.PushBack(tlabel)
					result.FalseLabel.PushBack(flabel)
					result.Val.FEti = flabel
					result.Val.TEti = tlabel

					return result
				}

			} else {
				Errores := environment.Errores{
					Descripcion: "No es posible comparar los dos valores(!=)",
					Fila:        strconv.Itoa(o.Lin),
					Columna:     strconv.Itoa(o.Col),
					Tipo:        "Error Semantico",
					Ambito:      ast.ObtenerAmbito(),
				}
				ast.ErroresHTML(Errores)
				return result
			}
		}
	case "&&":
		{
			newTemp := gen.NewTemp()
			V1 := gen.NewLabel()
			F1 := gen.NewLabel()
			V2 := gen.NewLabel()
			F2 := gen.NewLabel()
			auxV := gen.NewLabel()
			auxF := gen.NewLabel()

			op1 = o.Op_izq.Ejecutar(ast, gen)
			if op1.Type == environment.BOOLEAN {
				if op1.Value != "" {
					gen.AddIf(op1.Value, "1", "==", V1)
					gen.AddGoto(F1)
					gen.AddLabel(V1)
				} else {
					gen.AddLabel(op1.Val.TEti)
				}
			} else {
				Errores := environment.Errores{
					Descripcion: "No es posible comparar los dos valores(&&)",
					Fila:        strconv.Itoa(o.Lin),
					Columna:     strconv.Itoa(o.Col),
					Tipo:        "Error Semantico",
					Ambito:      ast.ObtenerAmbito(),
				}
				ast.ErroresHTML(Errores)
				return result
			}

			op2 = o.Op_der.Ejecutar(ast, gen)
			if op2.Type == environment.BOOLEAN {
				if op1.Value != "" && op2.Value != "" {
					gen.AddIf(op2.Value, "1", "==", V2)
					gen.AddGoto(F2)

					gen.AddLabel(V2)
					gen.AddExpression(newTemp, "1", "0", "+")
					gen.AddGoto(auxV)

					gen.AddLabel(F2)
					gen.AddExpression(newTemp, "0", "0", "+")
					gen.AddGoto(auxF)

					gen.AddLabel(F1)
					gen.AddExpression(newTemp, "0", "0", "+")
					gen.AddGoto(auxF)

				} else if op1.Value != "" && op2.Value == "" {
					gen.AddLabel(op2.Val.TEti)
					gen.AddExpression(newTemp, "1", "0", "+")
					gen.AddGoto(auxV)

					gen.AddLabel(op2.Val.FEti)
					gen.AddExpression(newTemp, "0", "0", "+")
					gen.AddGoto(auxF)

					gen.AddLabel(F1)
					gen.AddExpression(newTemp, "0", "0", "+")
					gen.AddGoto(auxF)

				} else if op1.Value == "" && op2.Value == "" {
					gen.AddLabel(op2.Val.TEti)
					gen.AddExpression(newTemp, "1", "0", "+")
					gen.AddGoto(auxV)

					gen.AddLabel(op2.Val.FEti)
					gen.AddExpression(newTemp, "0", "0", "+")
					gen.AddGoto(auxF)

					gen.AddLabel(op1.Val.FEti)
					gen.AddExpression(newTemp, "0", "0", "+")
					gen.AddGoto(auxF)

				} else {
					gen.AddIf(op2.Value, "1", "==", V2)
					gen.AddGoto(F2)

					gen.AddLabel(V2)
					gen.AddExpression(newTemp, "1", "0", "+")
					gen.AddGoto(auxV)

					gen.AddLabel(F2)
					gen.AddExpression(newTemp, "0", "0", "+")
					gen.AddGoto(auxF)

					gen.AddLabel(op1.Val.FEti)
					gen.AddExpression(newTemp, "0", "0", "+")
					gen.AddGoto(auxF)
				}
			} else {
				Errores := environment.Errores{
					Descripcion: "No es posible comparar los dos valores(&&)",
					Fila:        strconv.Itoa(o.Lin),
					Columna:     strconv.Itoa(o.Col),
					Tipo:        "Error Semantico",
					Ambito:      ast.ObtenerAmbito(),
				}
				ast.ErroresHTML(Errores)
				return result
			}

			result = environment.NewValue("", true, environment.BOOLEAN, false, false, false, environment.Variable{})
			result.TrueLabel.PushBack(auxV)
			result.FalseLabel.PushBack(auxF)
			result.Val.FEti = auxF
			result.Val.TEti = auxV

			return result
		}
	case "||":
		{
			newTemp := gen.NewTemp()
			V1 := gen.NewLabel()
			F1 := gen.NewLabel()
			V2 := gen.NewLabel()
			F2 := gen.NewLabel()
			auxV := gen.NewLabel()
			auxF := gen.NewLabel()

			op1 = o.Op_izq.Ejecutar(ast, gen)
			if op1.Type == environment.BOOLEAN {
				if op1.Value != "" {
					gen.AddIf(op1.Value, "1", "==", V1)
					gen.AddGoto(F1)
					gen.AddLabel(V1)
					gen.AddExpression(newTemp, "1", "0", "+")
					gen.AddGoto(auxV)
					gen.AddLabel(F1)
				} else {
					gen.AddLabel(op1.Val.TEti)
					gen.AddExpression(newTemp, "1", "0", "+")
					gen.AddGoto(auxV)
					gen.AddLabel(op1.Val.FEti)
				}
			} else {
				Errores := environment.Errores{
					Descripcion: "No es posible comparar los dos valores(&&)",
					Fila:        strconv.Itoa(o.Lin),
					Columna:     strconv.Itoa(o.Col),
					Tipo:        "Error Semantico",
					Ambito:      ast.ObtenerAmbito(),
				}
				ast.ErroresHTML(Errores)
				return result
			}

			op2 = o.Op_der.Ejecutar(ast, gen)
			if op2.Type == environment.BOOLEAN {
				if op1.Value != "" && op2.Value != "" {
					gen.AddIf(op2.Value, "1", "==", V2)
					gen.AddGoto(F2)

					gen.AddLabel(V2)
					gen.AddExpression(newTemp, "1", "0", "+")
					gen.AddGoto(auxV)

					gen.AddLabel(F2)
					gen.AddExpression(newTemp, "0", "0", "+")
					gen.AddGoto(auxF)

				} else if op1.Value != "" && op2.Value == "" {
					gen.AddLabel(op2.Val.TEti)
					gen.AddExpression(newTemp, "1", "0", "+")
					gen.AddGoto(auxV)

					gen.AddLabel(op2.Val.FEti)
					gen.AddExpression(newTemp, "0", "0", "+")
					gen.AddGoto(auxF)

				} else if op1.Value == "" && op2.Value == "" {
					gen.AddLabel(op2.Val.TEti)
					gen.AddExpression(newTemp, "1", "0", "+")
					gen.AddGoto(auxV)

					gen.AddLabel(op2.Val.FEti)
					gen.AddExpression(newTemp, "0", "0", "+")
					gen.AddGoto(auxF)

				} else {
					gen.AddIf(op2.Value, "1", "==", V2)
					gen.AddGoto(F2)

					gen.AddLabel(V2)
					gen.AddExpression(newTemp, "1", "0", "+")
					gen.AddGoto(auxV)

					gen.AddLabel(F2)
					gen.AddExpression(newTemp, "0", "0", "+")
					gen.AddGoto(auxF)
				}
			} else {
				Errores := environment.Errores{
					Descripcion: "No es posible comparar los dos valores(&&)",
					Fila:        strconv.Itoa(o.Lin),
					Columna:     strconv.Itoa(o.Col),
					Tipo:        "Error Semantico",
					Ambito:      ast.ObtenerAmbito(),
				}
				ast.ErroresHTML(Errores)
				return result
			}

			result = environment.NewValue("", true, environment.BOOLEAN, false, false, false, environment.Variable{})
			result.TrueLabel.PushBack(auxV)
			result.FalseLabel.PushBack(auxF)
			result.Val.FEti = auxF
			result.Val.TEti = auxV

			return result
		}
	case "!":
		{
			op1 = o.Op_izq.Ejecutar(ast, gen)
			//validar tipo dominante

			if op1.Type == environment.BOOLEAN {

				if op1.Val.Name == "" {
					if op1.Value == "0" {
						result = environment.NewValue("1", true, environment.BOOLEAN, false, false, false, environment.Variable{})
						result.Val.FEti = op1.Val.FEti
						result.Val.TEti = op2.Val.TEti

					} else {
						result = environment.NewValue("0", false, environment.BOOLEAN, false, false, false, environment.Variable{})
					}
				} else {
					if op1.Value == "0" {
						result = environment.NewValue("1", false, environment.BOOLEAN, false, false, false, environment.Variable{})
					} else {
						result = environment.NewValue("0", false, environment.BOOLEAN, false, false, false, environment.Variable{})
					}
				}

				return result
			} else {
				Errores := environment.Errores{
					Descripcion: "No es posible comparar los dos valores(!)",
					Fila:        strconv.Itoa(o.Lin),
					Columna:     strconv.Itoa(o.Col),
					Tipo:        "Error Semantico",
					Ambito:      ast.ObtenerAmbito(),
				}
				ast.ErroresHTML(Errores)
				return result
			}
		}
	}

	return result
}
