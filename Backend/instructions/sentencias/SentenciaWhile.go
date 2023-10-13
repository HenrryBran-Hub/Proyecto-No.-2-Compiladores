package sentencias

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"strconv"
)

type SentenciaWhile struct {
	Lin       int
	Col       int
	Expresion interfaces.Expression
	slice     []interface{}
}

func NewSentenciaWhile(lin int, col int, expresion interfaces.Expression, bloque []interface{}) SentenciaWhile {
	return SentenciaWhile{lin, col, expresion, bloque}
}

func (v SentenciaWhile) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	ambito := ast.ObtenerAmbito()
	var errorgeneral int = 0
	ambitonuevo := "While" + "-" + ambito
	ast.AumentarAmbito(ambitonuevo)
	if !ast.IsMain(ambitonuevo) {
		gen.MainCodeT()
	}
	var retornable int = 0
	newlabelr := gen.NewLabel()
	gen.AddLabel(newlabelr)
	condicion := v.Expresion.Ejecutar(ast, gen)
	if !ast.IsMain(ambitonuevo) {
		gen.MainCodeT()
	}
	exitla := gen.NewLabel()

	transferencia := environment.SentenciasdeTransferencia{
		Scope:  ambitonuevo,
		ETrue:  newlabelr,
		EFalse: exitla,
	}
	ast.Lista_Tranferencias.PushBack(transferencia)

	if condicion.Type == environment.BOOLEAN {
		gen.AddComment("Estoy dentro de la sentencia while ")
		if condicion.Value == "1" || condicion.Value == "0" {
			newlabelr := gen.NewLabel()
			gen.AddLabel(newlabelr)
			newtemp1 := gen.NewTemp()
			gen.AddGetStack(newtemp1, strconv.Itoa(condicion.Val.Symbols.Posicion))
			vet := gen.NewLabel()
			fet := gen.NewLabel()

			gen.AddLabel(vet)
			ambito = ast.ObtenerAmbito()
			ambitonuevo = "While" + "-" + ambito
			ast.AumentarAmbito(ambitonuevo)
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
			ast.DisminuirAmbito()
			gen.AddGoto(newlabelr)
			gen.AddLabel(fet)
			gen.AddGoto(exitla)
			gen.AddLabel(exitla)
			gen.AddBr()
		} else if condicion.Value == "" && condicion.Val.TEti != "" {
			gen.AddLabel(condicion.Val.TEti)

			ambito = ast.ObtenerAmbito()
			ambitonuevo = "While" + "-" + ambito
			ast.AumentarAmbito(ambitonuevo)
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
			ast.DisminuirAmbito()
			gen.AddGoto(newlabelr)
			gen.AddLabel(condicion.Val.FEti)
			gen.AddGoto(exitla)
			gen.AddLabel(exitla)
			gen.AddBr()
		} else if condicion.Value != "" && condicion.Val.TEti != "" {
			gen.AddLabel(condicion.Val.TEti)
			ambito = ast.ObtenerAmbito()
			ambitonuevo = "While" + "-" + ambito
			ast.AumentarAmbito(ambitonuevo)
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
			ast.DisminuirAmbito()
			gen.AddGoto(newlabelr)
			gen.AddLabel(condicion.Val.FEti)
			gen.AddGoto(exitla)
			gen.AddLabel(exitla)
			gen.AddBr()
		}

	} else {
		Errores := environment.Errores{
			Descripcion: "Se ha querido asignar un valor no correspondiente en la condicion del while tiene que ser un tipo boleano.",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
	}
	ast.DisminuirAmbito()
	tamanio := ast.Pila_Variables.Len()
	if tamanio == 1 && retornable == 3 {
		Errores := environment.Errores{
			Descripcion: "Estas retornando un valor fuera de una funcion",
			Fila:        strconv.Itoa(v.Lin),
			Columna:     strconv.Itoa(v.Col),
			Tipo:        "Error Semantico",
			Ambito:      ambitonuevo,
		}
		ast.ErroresHTML(Errores)
	}

	ast.Lista_Tranferencias.Remove(ast.Lista_Tranferencias.Back())
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
