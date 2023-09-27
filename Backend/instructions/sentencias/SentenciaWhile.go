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
	ambitonuevo := "While" + "-" + ambito
	ast.AumentarAmbito(ambitonuevo)
	if !ast.IsMain(ambitonuevo) {
		gen.MainCodeT()
	}
	var retornable int = 0
	var reexp environment.Symbol
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
				}
				rvari := ast.GetVariable("Return")
				if rvari != nil {
					retornable = 2
				}
				revari := ast.GetVariable("ReturnExp")
				if revari != nil {
					retornable = 3
					reexp = revari.Symbols
				}
			}
			gen.AddGoto(newlabelr)
			gen.AddLabel(fet)
			gen.AddGoto(exitla)
			gen.AddLabel(exitla)
			gen.AddBr()
		} else if condicion.Value == "" && condicion.Val.TEti != "" {
			gen.AddLabel(condicion.Val.TEti)

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
				}
				rvari := ast.GetVariable("Return")
				if rvari != nil {
					retornable = 2
					reexp = rvari.Symbols
				}
				revari := ast.GetVariable("ReturnExp")
				if revari != nil {
					retornable = 3
					reexp = revari.Symbols
				}
			}
			gen.AddGoto(newlabelr)
			gen.AddLabel(condicion.Val.FEti)
			gen.AddGoto(exitla)
			gen.AddLabel(exitla)
			gen.AddBr()
		} else if condicion.Value != "" && condicion.Val.TEti != "" {
			gen.AddLabel(condicion.Val.TEti)

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
				}
				rvari := ast.GetVariable("Return")
				if rvari != nil {
					retornable = 2
					reexp = rvari.Symbols
				}
				revari := ast.GetVariable("ReturnExp")
				if revari != nil {
					retornable = 3
					reexp = revari.Symbols
				}
			}
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
			Ambito:      ambitonuevo,
		}
		ast.ErroresHTML(Errores)
	}

	ast.Lista_Tranferencias.Remove(ast.Lista_Tranferencias.Back())
	return nil
}
