package sentencias

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"fmt"
	"strconv"
)

type SentenciaIfElseIf struct {
	Lin       int
	Col       int
	Expresion interfaces.Expression
	Ifop      []interface{}
	Elseop    interfaces.Instruction
}

func NewSentenciaIfElseIf(lin int, col int, expresion interfaces.Expression, Ifop []interface{}, Elseop interfaces.Instruction) SentenciaIfElseIf {
	return SentenciaIfElseIf{lin, col, expresion, Ifop, Elseop}
}

func (v SentenciaIfElseIf) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	condicion := v.Expresion.Ejecutar(ast, gen)
	ambito := ast.ObtenerAmbito()
	ambitonuevo := "If-If-Else-Else" + "-" + ambito
	ast.AumentarAmbito(ambitonuevo)
	if !ast.IsMain(ambitonuevo) {
		gen.MainCodeT()
	}
	var retornable int = 0
	var reexp environment.Symbol
	var errorgeneral int = 0

	if condicion.Type == environment.BOOLEAN {
		gen.AddComment("Estoy dentro de la sentencia if-if-else-else")
		if condicion.Value == "1" || condicion.Value == "0" {
			vet := gen.NewLabel()
			fet := gen.NewLabel()
			exitla := gen.NewLabel()
			gen.AddIf(condicion.Value, "1", "==", vet)
			gen.AddGoto(fet)
			gen.AddLabel(vet)
			for _, inst := range v.Ifop {
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

			gen.AddGoto(exitla)
			gen.AddLabel(fet)
			v.Elseop.Ejecutar(ast, gen)
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
			gen.AddGoto(exitla)
			gen.AddLabel(exitla)
			gen.AddBr()
		} else if condicion.Value == "" && condicion.Val.TEti != "" {
			fmt.Println("Condicion 2 if-else-if-else")
			exitl := gen.NewLabel()
			gen.AddLabel(condicion.Val.TEti)
			for _, inst := range v.Ifop {
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
					reexp = rvari.Symbols
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

			gen.AddGoto(exitl)
			gen.AddLabel(condicion.Val.FEti)
			v.Elseop.Ejecutar(ast, gen)
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

			gen.AddGoto(exitl)
			gen.AddLabel(exitl)
			gen.AddBr()
		} else if condicion.Value != "" && condicion.Val.TEti != "" {
			exitl := gen.NewLabel()
			gen.AddLabel(condicion.Val.TEti)
			for _, inst := range v.Ifop {
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
			gen.AddGoto(exitl)
			gen.AddLabel(condicion.Val.FEti)
			v.Elseop.Ejecutar(ast, gen)
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

			gen.AddGoto(exitl)
			gen.AddLabel(exitl)
			gen.AddBr()
		}

	} else {
		Errores := environment.Errores{
			Descripcion: "Se ha querido asignar un valor no correspondiente en la condicion del if tiene que ser un tipo boleano.",
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
	return nil
}
