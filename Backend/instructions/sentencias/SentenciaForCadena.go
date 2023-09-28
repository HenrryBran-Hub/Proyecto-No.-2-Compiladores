package sentencias

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"strconv"
)

type SentenciaForCadena struct {
	Lin   int
	Col   int
	Id    string
	Expre interfaces.Expression
	slice []interface{}
}

func NewSentenciaForCadena(lin int, col int, id string, expre interfaces.Expression, bloque []interface{}) SentenciaForCadena {
	return SentenciaForCadena{lin, col, id, expre, bloque}
}

func (v SentenciaForCadena) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	ambito := ast.ObtenerAmbito()
	ambitonuevo := "Switch" + "-" + ambito
	ast.AumentarAmbito(ambitonuevo)
	if !ast.IsMain(ambitonuevo) {
		gen.MainCodeT()
	}
	cadena := v.Expre.Ejecutar(ast, gen)
	if !ast.IsMain(ambitonuevo) {
		gen.MainCodeT()
	}

	var errorgeneral int = 0
	if cadena.Type == environment.VECTOR || cadena.Type == environment.STRING {
		newTemp := gen.NewTemp()
		symbol := environment.Symbol{
			Lin:   v.Lin,
			Col:   v.Col,
			Tipo:  environment.CHARACTER,
			Valor: newTemp,
			Scope: ast.ObtenerAmbito(),
		}
		Variable := environment.Variable{
			Name:        v.Id,
			Symbols:     symbol,
			Mutable:     true,
			TipoSimbolo: "Variable",
		}

		ast.GuardarVariable(Variable)

		gen.AddSetStack(strconv.Itoa(ast.PosicionStack), cadena.Value)
		gen.AddBr()

		gen.AddAssign(newTemp, "H")
		gen.AddSetHeap("(int)H", "-1")
		gen.AddExpression("H", "H", "1", "+")
		gen.AddBr()

		gen.AddSetStack(strconv.Itoa(symbol.Posicion), newTemp)
		gen.AddBr()

		var retornable int = 0
		var reexp environment.Symbol
		looptl := gen.NewLabel()
		exitla := gen.NewLabel()
		transferencia := environment.SentenciasdeTransferencia{
			Scope:  ambitonuevo,
			ETrue:  looptl,
			EFalse: exitla,
		}
		ast.Lista_Tranferencias.PushBack(transferencia)
		ast.Lista_For_Rango.PushBack(Variable)

		if cadena.Type == environment.STRING {
			tmp1 := gen.NewTemp()
			tmp2 := gen.NewTemp()
			tmp3 := gen.NewTemp()
			tmp4 := gen.NewTemp()
			lvl2 := gen.NewLabel()
			lvl3 := gen.NewLabel()

			gen.AddAssign(tmp1, "H")
			gen.AddGetStack(tmp2, strconv.Itoa(ast.PosicionStack))
			gen.AddLabel(looptl)
			gen.AddGetHeap(tmp3, "(int)"+tmp2)
			gen.AddIf(tmp3, "-1", "!=", lvl2)
			gen.AddGoto(lvl3)
			gen.AddLabel(lvl2)
			gen.AddAssign(tmp4, "H")
			gen.AddSetHeap("(int)H", tmp3)
			gen.AddExpression("H", "H", "1", "+")
			gen.AddSetHeap("(int)H", "-1")
			gen.AddExpression("H", "H", "1", "+")
			gen.AddExpression(tmp2, tmp2, "1", "+")
			gen.AddSetStack(strconv.Itoa(Variable.Symbols.Posicion), tmp4)
			Variable.Symbols.Valor = tmp4
			ast.ActualizarVariable(&Variable)
			for _, inst := range v.slice {
				if inst == nil {
					continue
				}
				instruction, ok := inst.(interfaces.Instruction)
				if !ok {
					continue
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

			gen.AddGoto(looptl)
			gen.AddLabel(lvl3)
			gen.AddGoto(exitla)
			gen.AddLabel(exitla)
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
				Ambito:      "For Cadena",
			}
			ast.ErroresHTML(Errores)
		}
	} else {
		Errores := environment.Errores{
			Descripcion: "El for que se esta ejecutando solo funciona con arreglos unidimensionales",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		gen.MainCodeF()
		return nil
	}

	gen.MainCodeF()

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

	ast.Lista_Tranferencias.Remove(ast.Lista_Tranferencias.Back())
	ast.Lista_For_Rango.Remove(ast.Lista_For_Rango.Back())

	return nil
}
