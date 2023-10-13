package sentencias

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"strconv"
)

type SentenciaForId struct {
	Lin   int
	Col   int
	Id    string
	Id2   string
	slice []interface{}
}

func NewSentenciaForId(lin int, col int, id string, id2 string, bloque []interface{}) SentenciaForId {
	return SentenciaForId{lin, col, id, id2, bloque}
}

func (v SentenciaForId) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	variables := ast.GetVariable(v.Id2)
	arreglos := ast.GetArreglo(v.Id2)
	var errorgeneral int = 0
	if variables != nil {
		if variables.Symbols.Tipo == environment.STRING {
			ambito := ast.ObtenerAmbito()
			ambitonuevo := "for" + "-" + ambito
			ast.AumentarAmbito(ambitonuevo)
			if !ast.IsMain(ambitonuevo) {
				gen.MainCodeT()
			}

			newTemp := gen.NewTemp()
			symbol := environment.Symbol{
				Lin:      v.Lin,
				Col:      v.Col,
				Tipo:     environment.CHARACTER,
				Valor:    newTemp,
				Scope:    ast.ObtenerAmbito(),
				Posicion: ast.PosicionStack,
			}
			Variable := environment.Variable{
				Name:        v.Id,
				Symbols:     symbol,
				Mutable:     true,
				TipoSimbolo: "Variable",
			}

			ast.GuardarVariable(Variable)

			gen.AddAssign(newTemp, "H")
			gen.AddSetHeap("(int)H", "-1")
			gen.AddExpression("H", "H", "1", "+")
			gen.AddBr()

			gen.AddSetStack(strconv.Itoa(Variable.Symbols.Posicion), newTemp)
			gen.AddBr()

			var retornable int = 0
			looptl := gen.NewLabel()
			exitla := gen.NewLabel()
			transferencia := environment.SentenciasdeTransferencia{
				Scope:  ambitonuevo,
				ETrue:  looptl,
				EFalse: exitla,
			}
			ast.Lista_Tranferencias.PushBack(transferencia)
			ast.Lista_For_Rango.PushBack(Variable)

			tmp1 := gen.NewTemp()
			tmp2 := gen.NewTemp()
			tmp3 := gen.NewTemp()
			tmp4 := gen.NewTemp()
			lvl2 := gen.NewLabel()
			lvl3 := gen.NewLabel()

			gen.AddAssign(tmp1, "H")
			gen.AddGetStack(tmp2, strconv.Itoa(variables.Symbols.Posicion))
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

			ast.DisminuirAmbito()
			tamanio := ast.Pila_Variables.Len()
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
				Descripcion: "El id ingresado no a una variable tipo string",
				Fila:        strconv.Itoa(v.Lin),
				Columna:     strconv.Itoa(v.Col),
				Tipo:        "Error Semantico",
				Ambito:      ast.ObtenerAmbito(),
			}
			ast.ErroresHTML(Errores)
			return nil
		}
	}

	if arreglos != nil {
		ambito := ast.ObtenerAmbito()
		ambitonuevo := "for" + "-" + ambito
		ast.AumentarAmbito(ambitonuevo)
		if !ast.IsMain(ambitonuevo) {
			gen.MainCodeT()
		}
		symbol := environment.Symbol{
			Lin:      v.Lin,
			Col:      v.Col,
			Tipo:     arreglos.Symbols.Tipo,
			Valor:    nil,
			Scope:    ast.ObtenerAmbito(),
			Posicion: arreglos.Symbols.Posicion,
		}
		Variable := environment.Variable{
			Name:        v.Id,
			Symbols:     symbol,
			Mutable:     true,
			TipoSimbolo: "Variable",
		}

		var retornable int = 0

		looptl := gen.NewLabel()
		exitla := gen.NewLabel()
		transferencia := environment.SentenciasdeTransferencia{
			Scope:  ambitonuevo,
			ETrue:  looptl,
			EFalse: exitla,
		}
		ast.Lista_Tranferencias.PushBack(transferencia)
		ast.Lista_For_Rango.PushBack(Variable)

		tmp := gen.NewTemp()
		tmp2 := gen.NewTemp()

		gen.AddBr()

		gen.AddLabel(looptl)
		ast.GuardarVariable(Variable)
		for i := arreglos.Elements.Front(); i != nil; i = i.Next() {
			if arreglos.Symbols.Tipo == environment.CHARACTER || arreglos.Symbols.Tipo == environment.STRING {
				vari := ast.GetVariable(v.Id)
				valor := i.Value.(environment.Value)
				vari.Symbols.Valor = valor.Value
				gen.AddAssign(tmp, valor.Value)
				gen.AddGetHeap(tmp2, "(int)"+tmp)
				gen.AddSetStack(strconv.Itoa(arreglos.Symbols.Posicion), tmp2)
				ast.ActualizarVariable(vari)
			} else {
				vari := ast.GetVariable(v.Id)
				valor := i.Value.(environment.Value)
				vari.Symbols.Valor = valor.Value
				gen.AddAssign(tmp, valor.Value)
				gen.AddSetStack(strconv.Itoa(arreglos.Symbols.Posicion), tmp)
				ast.ActualizarVariable(vari)
			}

			for _, inst := range v.slice {
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
		}

		gen.AddBr()
		gen.AddGoto(exitla)
		gen.AddLabel(exitla)

		ast.DisminuirAmbito()
		tamanio := ast.Pila_Variables.Len()
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

	}

	if variables == nil && arreglos == nil {
		Errores := environment.Errores{
			Descripcion: "El id ingresado no a una variable tipo string",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
		return nil
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

	ast.Lista_Tranferencias.Remove(ast.Lista_Tranferencias.Back())
	ast.Lista_For_Rango.Remove(ast.Lista_For_Rango.Back())
	gen.MainCodeF()
	return nil
}
