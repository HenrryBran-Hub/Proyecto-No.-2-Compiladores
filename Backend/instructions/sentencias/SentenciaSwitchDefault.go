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
	var reexp environment.Symbol
	ambito := ast.ObtenerAmbito()
	ambitonuevo := "Switch" + "-" + ambito
	ast.AumentarAmbito(ambitonuevo)
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
		switchCase.Ejecutar(ast, gen)
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
	gen.AddGoto(exitla)
	gen.AddLabel(exitla)
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

	ast.Lista_Switch_Case.Remove(ast.Lista_Switch_Case.Back())
	ast.Lista_Switch_Case_Eti.Remove(ast.Lista_Switch_Case_Eti.Back())
	return nil
}
