package sentencias

import (
	"Backend/environment"
	"Backend/generator"
	"Backend/interfaces"
	"strconv"
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

func (v SentenciaSwitchDefault) Ejecutar(ast *environment.AST, gen *generator.Generator) interface{} {
	var retornable int = 0
	var errorgeneral int = 0
	ambito := ast.ObtenerAmbito()
	ambitonuevo := "Switch" + "-" + ambito
	ast.AumentarAmbito(ambitonuevo)
	if !ast.IsMain(ambitonuevo) {
		gen.MainCodeT()
	}
	if !ast.IsMain(ambitonuevo) {
		gen.MainCodeT()
	}
	valorswitch := v.Expresion.Ejecutar(ast, gen)
	if !ast.IsMain(ambitonuevo) {
		gen.MainCodeT()
	}
	exitla := gen.NewLabel()
	ast.Lista_Switch_Case.PushBack(valorswitch)
	transferencia := environment.SentenciasdeTransferencia{
		Scope:  ambitonuevo,
		ETrue:  exitla,
		EFalse: exitla,
	}
	ast.Lista_Switch_Case_Eti.PushBack(transferencia)
	ast.Lista_Tranferencias.PushBack(transferencia)

	for _, inst := range v.Case {
		switchCase, _ := inst.(SentenciaSwitchCase)
		if !ast.IsMain(ambitonuevo) {
			gen.MainCodeT()
		}
		switchCase.Ejecutar(ast, gen)
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
	}
	ast.DisminuirAmbito()
	ambito = ast.ObtenerAmbito()
	ambitonuevo = "Switch-Default" + "-" + ambito
	ast.AumentarAmbito(ambitonuevo)
	if !ast.IsMain(ambitonuevo) {
		gen.MainCodeT()
	}
	for _, inst := range v.Default {
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
	}
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
			Ambito:      ast.ObtenerAmbito(),
		}
		ast.ErroresHTML(Errores)
	}

	ast.Lista_Switch_Case.Remove(ast.Lista_Switch_Case.Back())
	ast.Lista_Switch_Case_Eti.Remove(ast.Lista_Switch_Case_Eti.Back())
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
