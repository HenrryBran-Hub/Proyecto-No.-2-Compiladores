package sentencias

import (
	"Backend/interfaces"
)

type SentenciaSwitchDefault struct {
	Lin       int
	Col       int
	Expresion interfaces.Expression
	Case      []interface{}
	Default   []interface{}
}

func NewSentenciaSwitchDefault(lin int, col int, expresion interfaces.Expression, cases []interface{}, defaultt []interface{}) SentenciaSwitchDefault {
	return SentenciaSwitchDefault{lin, col, expresion, cases, defaultt}
}

/*
func (v SentenciaSwitchDefault) Ejecutar(ast *environment.AST) interface{} {
	var retornable int = 0
	var reexp environment.Symbol
	for _, inst := range v.Case {
		switchCase, _ := inst.(SentenciaSwitchCase)
		valorcase := switchCase.Exp1.Ejecutar(ast)
		valorswitch := v.Expresion.Ejecutar(ast)
		if valorcase.Valor == valorswitch.Valor && valorcase.Tipo == valorswitch.Tipo {
			ast.AumentarAmbito("Switch Default")
			for _, inst := range switchCase.Case {
				if inst == nil {
					continue
				}
				instruction, ok := inst.(interfaces.Instruction)
				if !ok {
					continue
				}
				instruction.Ejecutar(ast)
				bvari := ast.GetVariable("Break")
				if bvari != nil {
					retornable = 1
					break
				}
				rvari := ast.GetVariable("Return")
				if rvari != nil {
					retornable = 2
					break
				}
				revari := ast.GetVariable("ReturnExp")
				if revari != nil {
					retornable = 3
					reexp = revari.Symbols
					break
				}
				cvari := ast.GetVariable("Continue")
				if cvari != nil {
					continue
				}
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
					Ambito:      "Switch Default",
				}
				ast.ErroresHTML(Errores)
			}
			return nil
		}
	}

	ast.AumentarAmbito("Switch Default")
	for _, inst := range v.Default {
		if inst == nil {
			continue
		}
		instruction, ok := inst.(interfaces.Instruction)
		if !ok {
			continue
		}
		instruction.Ejecutar(ast)
		bvari := ast.GetVariable("Break")
		if bvari != nil {
			retornable = 1
			break
		}
		rvari := ast.GetVariable("Return")
		if rvari != nil {
			retornable = 2
			break
		}
		revari := ast.GetVariable("ReturnExp")
		if revari != nil {
			retornable = 3
			reexp = revari.Symbols
			break
		}
		cvari := ast.GetVariable("Continue")
		if cvari != nil {
			continue
		}
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
	return nil
}
*/
