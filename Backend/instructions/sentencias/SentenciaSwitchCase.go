package sentencias

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"strconv"
)

type SentenciaSwitchCase struct {
	Lin  int
	Col  int
	Exp1 interfaces.Expression
	Case []interface{}
}

func NewSentenciaSwitchCase(lin int, col int, expr1 interfaces.Expression, operations []interface{}) SentenciaSwitchCase {
	return SentenciaSwitchCase{lin, col, expr1, operations}
}

func (v SentenciaSwitchCase) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	ambito := ast.ObtenerAmbito()
	ambitonuevo := "Case" + "-" + ambito
	ast.AumentarAmbito(ambitonuevo)
	if !ast.IsMain(ambitonuevo) {
		gen.MainCodeT()
	}
	if !ast.IsMain(ambitonuevo) {
		gen.MainCodeT()
	}
	expresion2 := v.Exp1.Ejecutar(ast, gen)
	if !ast.IsMain(ambitonuevo) {
		gen.MainCodeT()
	}
	var retornable int = 0
	var reexp environment.Symbol
	var errorgeneral int = 0

	expresion1 := ast.Lista_Switch_Case.Back().Value.(environment.Value)
	if expresion1.Type == expresion2.Type {
		{
			op1 := expresion1
			op2 := expresion2

			gen.AddComment("Estoy dentro de la sentencia case ")
			if expresion1.Type == environment.INTEGER {
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
				gen.AddLabel(tlabel)
				for _, inst := range v.Case {
					if inst == nil {
						continue
					}
					instruction, ok := inst.(interfaces.Instruction)
					if !ok {
						continue
					}
					if !ast.IsMain(ambitonuevo) {
						gen.MainCodeT()
					}
					instruction.Ejecutar(ast, gen)
					if !ast.IsMain(ambitonuevo) {
						gen.MainCodeT()
					}
					bvari := ast.GetVariable("Break")
					if bvari != nil {
						retornable = 1
						if ast.Lista_Tranferencias.Len() == 0 {
							errorgeneral = 1
						}
					}
					rvari := ast.GetVariable("Return")
					if rvari != nil {
						retornable = 2
						if ast.Lista_Tranferencias.Len() == 0 {
							errorgeneral = 1
						}
					}
					revari := ast.GetVariable("ReturnExp")
					if revari != nil {
						retornable = 3
						reexp = revari.Symbols
						if ast.Lista_Tranferencias.Len() == 0 {
							errorgeneral = 1
						}
					}
					cvari := ast.GetVariable("Continue")
					if cvari != nil {
						if ast.Lista_Tranferencias.Len() == 0 {
							errorgeneral = 1
						}
					}
				}
				etiqueta := ast.Lista_Switch_Case_Eti.Back().Value.(environment.SentenciasdeTransferencia)
				gen.AddGoto(etiqueta.EFalse)
				gen.AddLabel(flabel)
				gen.AddBr()

			} else if expresion1.Type == environment.FLOAT {
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
				gen.AddLabel(tlabel)
				for _, inst := range v.Case {
					if inst == nil {
						continue
					}
					instruction, ok := inst.(interfaces.Instruction)
					if !ok {
						continue
					}
					if !ast.IsMain(ambitonuevo) {
						gen.MainCodeT()
					}
					instruction.Ejecutar(ast, gen)
					if !ast.IsMain(ambitonuevo) {
						gen.MainCodeT()
					}
					bvari := ast.GetVariable("Break")
					if bvari != nil {
						retornable = 1
						if ast.Lista_Tranferencias.Len() == 0 {
							errorgeneral = 1
						}
					}
					rvari := ast.GetVariable("Return")
					if rvari != nil {
						retornable = 2
						if ast.Lista_Tranferencias.Len() == 0 {
							errorgeneral = 1
						}
					}
					revari := ast.GetVariable("ReturnExp")
					if revari != nil {
						retornable = 3
						reexp = revari.Symbols
						if ast.Lista_Tranferencias.Len() == 0 {
							errorgeneral = 1
						}
					}
					cvari := ast.GetVariable("Continue")
					if cvari != nil {
						if ast.Lista_Tranferencias.Len() == 0 {
							errorgeneral = 1
						}
					}
				}
				etiqueta := ast.Lista_Switch_Case_Eti.Back().Value.(environment.SentenciasdeTransferencia)
				gen.AddGoto(etiqueta.EFalse)
				gen.AddLabel(flabel)
				gen.AddBr()

			} else if expresion1.Type == environment.STRING {
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
				gen.AddLabel(tlabel)
				for _, inst := range v.Case {
					if inst == nil {
						continue
					}
					instruction, ok := inst.(interfaces.Instruction)
					if !ok {
						continue
					}
					if !ast.IsMain(ambitonuevo) {
						gen.MainCodeT()
					}
					instruction.Ejecutar(ast, gen)
					if !ast.IsMain(ambitonuevo) {
						gen.MainCodeT()
					}
					bvari := ast.GetVariable("Break")
					if bvari != nil {
						retornable = 1
						if ast.Lista_Tranferencias.Len() == 0 {
							errorgeneral = 1
						}
					}
					rvari := ast.GetVariable("Return")
					if rvari != nil {
						retornable = 2
						if ast.Lista_Tranferencias.Len() == 0 {
							errorgeneral = 1
						}
					}
					revari := ast.GetVariable("ReturnExp")
					if revari != nil {
						retornable = 3
						reexp = revari.Symbols
						if ast.Lista_Tranferencias.Len() == 0 {
							errorgeneral = 1
						}
					}
					cvari := ast.GetVariable("Continue")
					if cvari != nil {
						if ast.Lista_Tranferencias.Len() == 0 {
							errorgeneral = 1
						}
					}
				}
				etiqueta := ast.Lista_Switch_Case_Eti.Back().Value.(environment.SentenciasdeTransferencia)
				gen.AddGoto(etiqueta.EFalse)
				gen.AddLabel(flabel)
				gen.AddBr()

			} else if expresion1.Type == environment.CHARACTER {
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
				gen.AddLabel(tlabel)
				for _, inst := range v.Case {
					if inst == nil {
						continue
					}
					instruction, ok := inst.(interfaces.Instruction)
					if !ok {
						continue
					}
					if !ast.IsMain(ambitonuevo) {
						gen.MainCodeT()
					}
					instruction.Ejecutar(ast, gen)
					if !ast.IsMain(ambitonuevo) {
						gen.MainCodeT()
					}
					bvari := ast.GetVariable("Break")
					if bvari != nil {
						retornable = 1
						if ast.Lista_Tranferencias.Len() == 0 {
							errorgeneral = 1
						}
					}
					rvari := ast.GetVariable("Return")
					if rvari != nil {
						retornable = 2
						if ast.Lista_Tranferencias.Len() == 0 {
							errorgeneral = 1
						}
					}
					revari := ast.GetVariable("ReturnExp")
					if revari != nil {
						retornable = 3
						reexp = revari.Symbols
						if ast.Lista_Tranferencias.Len() == 0 {
							errorgeneral = 1
						}
					}
					cvari := ast.GetVariable("Continue")
					if cvari != nil {
						if ast.Lista_Tranferencias.Len() == 0 {
							errorgeneral = 1
						}
					}
				}
				etiqueta := ast.Lista_Switch_Case_Eti.Back().Value.(environment.SentenciasdeTransferencia)
				gen.AddGoto(etiqueta.EFalse)
				gen.AddLabel(flabel)
				gen.AddBr()

			} else if expresion1.Type == environment.BOOLEAN {
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
				gen.AddLabel(tlabel)
				for _, inst := range v.Case {
					if inst == nil {
						continue
					}
					instruction, ok := inst.(interfaces.Instruction)
					if !ok {
						continue
					}
					if !ast.IsMain(ambitonuevo) {
						gen.MainCodeT()
					}
					instruction.Ejecutar(ast, gen)
					if !ast.IsMain(ambitonuevo) {
						gen.MainCodeT()
					}
					bvari := ast.GetVariable("Break")
					if bvari != nil {
						retornable = 1
						if ast.Lista_Tranferencias.Len() == 0 {
							errorgeneral = 1
						}
					}
					rvari := ast.GetVariable("Return")
					if rvari != nil {
						retornable = 2
						if ast.Lista_Tranferencias.Len() == 0 {
							errorgeneral = 1
						}
					}
					revari := ast.GetVariable("ReturnExp")
					if revari != nil {
						retornable = 3
						reexp = revari.Symbols
						if ast.Lista_Tranferencias.Len() == 0 {
							errorgeneral = 1
						}
					}
					cvari := ast.GetVariable("Continue")
					if cvari != nil {
						if ast.Lista_Tranferencias.Len() == 0 {
							errorgeneral = 1
						}
					}
				}
				etiqueta := ast.Lista_Switch_Case_Eti.Back().Value.(environment.SentenciasdeTransferencia)
				gen.AddGoto(etiqueta.EFalse)
				gen.AddLabel(flabel)
				gen.AddBr()
			}
		}
	} else {
		Errores := environment.Errores{
			Descripcion: "los valores ingresados para comparar no son del mismo tipo en el switch",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
	}
	ast.DisminuirAmbito()
	tamanio := ast.Pila_Variables.Len()
	if tamanio > 1 {
		if retornable == 2 {
			symbol := environment.Symbol{
				Lin:   v.Lin,
				Col:   v.Col,
				Tipo:  environment.BOOLEAN,
				Valor: true,
				Scope: ast.ObtenerAmbito(),
			}
			Variable := environment.Variable{
				Name:        "Return",
				Symbols:     symbol,
				Mutable:     false,
				TipoSimbolo: "Sentencia de Transferencia",
			}
			ast.GuardarVariable(Variable)
		}
		if retornable == 3 {
			symbol := environment.Symbol{
				Lin:   v.Lin,
				Col:   v.Col,
				Tipo:  reexp.Tipo,
				Valor: reexp.Valor,
				Scope: ast.ObtenerAmbito(),
			}
			Variable := environment.Variable{
				Name:        "ReturnExp",
				Symbols:     symbol,
				Mutable:     false,
				TipoSimbolo: "Sentencia de Transferencia",
			}
			ast.GuardarVariable(Variable)
		}
	}
	if tamanio == 1 && retornable == 3 {
		Errores := environment.Errores{
			Descripcion: "Estas retornando un valor fuera de una funcion",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
	}

	if errorgeneral == 1 {
		Errores := environment.Errores{
			Descripcion: "Se han colocado sentencias de transferencia fuera de ciclos",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
	}
	gen.MainCodeF()
	return nil
}
