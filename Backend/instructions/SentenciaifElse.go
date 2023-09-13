package instructions

import (
	"Backend/environment"
	"Backend/interfaces"
	"strconv"
)

type SentenciaIfElse struct {
	Lin       int
	Col       int
	Expresion interfaces.Expression
	Ifop      []interface{}
	Elseop    []interface{}
}

func NewSentenciaIfElse(lin int, col int, expresion interfaces.Expression, Ifop []interface{}, Elseop []interface{}) SentenciaIfElse {
	return SentenciaIfElse{lin, col, expresion, Ifop, Elseop}
}

func (v SentenciaIfElse) Ejecutar(ast *environment.AST) interface{} {
	var condicion environment.Symbol
	condicion = v.Expresion.Ejecutar(ast)
	var retornable int = 0
	var reexp environment.Symbol
	ast.AumentarAmbito("If-Else")
	if condicion.Tipo == environment.BOOLEAN {
		var continueflag bool = false
		breakPosition := -1
		if condicion.Valor.(bool) {

			for i, inst := range v.Ifop {
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
					breakPosition = i
					break
				}
				rvari := ast.GetVariable("Return")
				if rvari != nil {
					retornable = 2
					breakPosition = i
					break
				}
				revari := ast.GetVariable("ReturnExp")
				if revari != nil {
					retornable = 3
					reexp = revari.Symbols
					breakPosition = i
					break
				}
				cvari := ast.GetVariable("Continue")
				if cvari != nil {
					continueflag = true
					breakPosition = i
					break
				}
			}

			if continueflag {
				for _, inst := range v.Elseop {
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
						continueflag = true
						break
					}
				}
			}

		} else {
			for i, inst := range v.Elseop {
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
					breakPosition = i
					break
				}
				rvari := ast.GetVariable("Return")
				if rvari != nil {
					retornable = 2
					breakPosition = i
					break
				}
				revari := ast.GetVariable("ReturnExp")
				if revari != nil {
					retornable = 3
					reexp = revari.Symbols
					breakPosition = i
					break
				}
				cvari := ast.GetVariable("Continue")
				if cvari != nil {
					continueflag = true
					breakPosition = i
					break
				}
			}

			if continueflag {
				for i := breakPosition; i < len(v.Elseop); i++ {
					inst := v.Elseop[i]
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
						continueflag = true
						break
					}
				}
			}

		}
	} else {
		Errores := environment.Errores{
			Descripcion: "Se ha querido asignar un valor no correspondiente en la condicion del if-else tiene que ser un tipo boleano.",
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
		if retornable == 1 {
			symbol := environment.Symbol{
				Lin:   v.Lin,
				Col:   v.Col,
				Tipo:  environment.BOOLEAN,
				Valor: true,
				Scope: ast.ObtenerAmbito(),
			}
			Variable := environment.Variable{
				Name:        "Break",
				Symbols:     symbol,
				Mutable:     false,
				TipoSimbolo: "Sentencia de Transferencia",
			}
			ast.GuardarVariable(Variable)
		}
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
